package repository

import (
	"context"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
	"github.com/fahza-p/synapsis/model"
)

type UserRepository interface {
	Get(ctx context.Context) ([]*model.User, error)
}

type UserStore struct {
	db    *store.MysqlStore
	table string
}

func NewUserRepository(store *store.MysqlStore) (UserRepository, error) {
	logger := log.GetLogger(context.Background(), "User.Repository", "NewUserRepository")
	logger.Info("Add New User Repository")

	return &UserStore{db: store, table: "user"}, nil
}

func (s *UserStore) Get(ctx context.Context) ([]*model.User, error) {
	logger := log.GetLogger(ctx, "User.Repository", "Get")
	logger.Info("Repository Get User")

	var models []*model.User
	statment := "SELECT * FROM user"
	if err := s.db.Query(ctx, &models, statment); err != nil {
		return nil, err
	}

	return models, nil
}
