package auth

import (
	"context"
	"errors"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/model"
)

func (s *Service) Signup(ctx context.Context, model *model.AuthUserData) error {
	logger := log.GetLogger(ctx, "Auth.Service", "Signup")
	logger.Info("Signup")

	// Check User Is Exist
	isExists, err := s.auth.IsExists(ctx, "email", model.Email)
	if err != nil {
		logger.Errorf("can't get user data with email: '%s'", model.Email)
		return err
	}

	if isExists {
		return errors.New("data is already exist")
	}

	// Create Customer
	if err := model.SetUserSignupData(); err != nil {
		logger.Error(err.Error())
		return err
	}

	return s.auth.Create(ctx, model)
}
