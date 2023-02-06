package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
	"github.com/fahza-p/synapsis/model"
)

type CategoryRepository interface {
	FindOne(ctx context.Context, key string, val interface{}) (*model.Category, error)
	Get(ctx context.Context, queryParams *store.QueryParams) ([]*model.Category, int64, error)
	IsExists(ctx context.Context, key string, val interface{}) (bool, error)
	Create(ctx context.Context, req *model.Category) error
	Delete(ctx context.Context, key string, val interface{}) error
	UpdatePatch(ctx context.Context, email string, fields map[string]interface{}) error
}

type CategoryStore struct {
	db    *store.MysqlStore
	table string
}

func NewCategoryRepository(store *store.MysqlStore) (CategoryRepository, error) {
	logger := log.GetLogger(context.Background(), "Category.Repository", "NewCategoryRepository")
	logger.Info("Add New Category Repository")

	return &CategoryStore{db: store, table: "category"}, nil
}

func (s *CategoryStore) FindOne(ctx context.Context, key string, val interface{}) (*model.Category, error) {
	logger := log.GetLogger(ctx, "Category.Repository", "FindOne")
	logger.Info("Repository FindOne Category")

	var model model.Category

	statment := fmt.Sprintf("SELECT * FROM %s WHERE %s= ?", s.table, key)
	if err := s.db.QueryRow(ctx, &model, statment, val); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return &model, errors.New("document not found")
		}

		return &model, err
	}

	return &model, nil
}

func (s *CategoryStore) IsExists(ctx context.Context, key string, val interface{}) (bool, error) {
	logger := log.GetLogger(ctx, "Category.Repository", "IsExists")
	logger.Info("Repository IsExists Category")

	statment := fmt.Sprintf("SELECT COUNT(id) AS total FROM %s WHERE %s=?", s.table, key)

	total, err := s.db.Count(ctx, statment, val)
	if err != nil {
		return false, err
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

func (s *CategoryStore) Create(ctx context.Context, req *model.Category) error {
	logger := log.GetLogger(ctx, "Category.Repository", "Create")
	logger.Info("Repository Create Category")

	id, err := s.db.Insert(ctx, s.table, req)
	if err != nil {
		return err
	}

	req.Id = id

	return nil
}

func (s *CategoryStore) Delete(ctx context.Context, key string, val interface{}) error {
	logger := log.GetLogger(ctx, "Category.Repository", "Delete")
	logger.Info("Repository Delete Category")

	statment := fmt.Sprintf("DELETE FROM %s WHERE %s= ?", s.table, key)

	return s.db.Delete(ctx, statment, val)
}

func (s *CategoryStore) UpdatePatch(ctx context.Context, id string, fields map[string]interface{}) error {
	logger := log.GetLogger(ctx, "Category.Repository", "UpdatePatch")
	logger.Info("Repository UpdatePatch Category")

	query := "id=?"
	return s.db.Update(ctx, s.table, query, fields, id)
}

func (s *CategoryStore) Get(ctx context.Context, queryParams *store.QueryParams) ([]*model.Category, int64, error) {
	logger := log.GetLogger(ctx, "Category.Repository", "Get")
	logger.Info("Repository Get Category")

	var models []*model.Category
	limit, offset, sort, filter, keywords := queryParams.BuildPagination(model.CategoryFilter)

	statment := fmt.Sprintf(`
	SELECT *
	FROM category
	WHERE %s AND (%s)
	%s
	%s %s
	`, filter, keywords, sort, limit, offset)

	if err := s.db.Query(ctx, &models, statment, true); err != nil {
		return nil, 0, err
	}

	countStetment := fmt.Sprintf(`
	SELECT 
		COUNT(id) AS total
	FROM category
	WHERE %s AND (%s)
	`, filter, keywords)

	totalData, err := s.db.Count(ctx, countStetment)
	if err != nil {
		return nil, 0, err
	}

	return models, totalData, nil
}
