package auth

import (
	"net/http"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/response"
	"github.com/fahza-p/synapsis/lib/validation"
	"github.com/fahza-p/synapsis/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Signin(c *fiber.Ctx) error {
	logger := log.GetLogger(c.Context(), "Auth.Handler", "Signin")
	logger.Info("Signin")

	var (
		res      response.Build
		reqModel model.AuthSigninReq
	)

	if err := c.BodyParser(&reqModel); err != nil {
		logger.Error("Unable to parsing body data")
		res.Msg = err.Error()
		return res.BuildResponse(c, http.StatusBadRequest)
	}

	if err := validation.RunValidate(&reqModel, &res); err != nil {
		logger.Errorf("invalid request data %v", reqModel)
		return res.BuildResponse(c, http.StatusUnprocessableEntity)
	}

	token, err := h.service.Signin(c.Context(), &reqModel)
	if err != nil {
		logger.WithError(err).Errorf("unable to signup with request data %v", reqModel)
		res.Msg = err.Error()

		switch res.Msg {
		case "user not found":
			return res.BuildResponse(c, http.StatusNotFound)
		case "invalid password":
			return res.BuildResponse(c, http.StatusUnauthorized)
		default:
			return res.BuildResponse(c, http.StatusInternalServerError)
		}
	}

	res.Data = map[string]interface{}{"token": token}
	return res.BuildResponse(c, http.StatusOK)
}
