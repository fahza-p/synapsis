package auth

import (
	"context"
	"errors"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/model"
)

func (s *Service) Signin(ctx context.Context, req *model.AuthSigninReq) (string, error) {
	logger := log.GetLogger(ctx, "Auth.Service", "Signin")
	logger.Info("Signin")

	// Find User
	userData, err := s.auth.FindOne(ctx, "email", req.Email)
	if err != nil {
		if err.Error() == "document not found" {
			return "", errors.New("user not found")
		}

		return "", err
	}

	// Check User Password
	if !userData.IsPasswordValid(req.Password) {
		return "", errors.New("invalid password")
	}

	// Generate Token
	sub := map[string]interface{}{
		"id":    userData.Id,
		"role":  userData.Role,
		"email": userData.Email,
	}

	return s.auth.GenerateJwt(ctx, sub)
}
