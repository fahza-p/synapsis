package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
	"github.com/fahza-p/synapsis/model"
	"github.com/golang-jwt/jwt/v4"
)

type AuthRepository interface {
	IsExists(ctx context.Context, key string, val interface{}) (bool, error)
	Create(ctx context.Context, req *model.AuthUserData) error
	FindOne(ctx context.Context, key string, val interface{}) (*model.AuthUserData, error)
	GenerateJwt(ctx context.Context, sub map[string]interface{}) (string, error)
}

type AuthStore struct {
	db    *store.MysqlStore
	table string
}

func NewAuthRepository(store *store.MysqlStore) (AuthRepository, error) {
	logger := log.GetLogger(context.Background(), "Auth.Repository", "NewAuthRepository")
	logger.Info("Add New Auth Repository")

	return &AuthStore{db: store, table: "user"}, nil
}

func (s *AuthStore) IsExists(ctx context.Context, key string, val interface{}) (bool, error) {
	logger := log.GetLogger(ctx, "Auth.Repository", "IsExists")
	logger.Info("Repository IsExists Auth")

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

func (s *AuthStore) Create(ctx context.Context, req *model.AuthUserData) error {
	logger := log.GetLogger(ctx, "Auth.Repository", "Create")
	logger.Info("Repository Create Auth")

	id, err := s.db.Insert(ctx, s.table, req)
	if err != nil {
		return err
	}

	req.Id = id

	return nil
}

func (s *AuthStore) FindOne(ctx context.Context, key string, val interface{}) (*model.AuthUserData, error) {
	logger := log.GetLogger(ctx, "Auth.Repository", "FindOne")
	logger.Info("Repository FindOne Auth")

	var model model.AuthUserData

	statment := fmt.Sprintf("SELECT * FROM %s WHERE %s= ?", s.table, key)
	if err := s.db.QueryRow(ctx, &model, statment, val); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return &model, errors.New("document not found")
		}

		return &model, err
	}

	return &model, nil
}

func (s *AuthStore) GenerateJwt(ctx context.Context, sub map[string]interface{}) (string, error) {
	logger := log.GetLogger(ctx, "Auth.Repository", "GenerateJwt")
	logger.Info("Repository GenerateJwt Auth")

	claims := jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(time.Minute * 120).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_KEY")))
}
