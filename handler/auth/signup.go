package auth

import (
	"net/http"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/response"
	"github.com/fahza-p/synapsis/lib/validation"
	"github.com/fahza-p/synapsis/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Signup(c *fiber.Ctx) error {
	logger := log.GetLogger(c.Context(), "Auth.Handler", "SignUp")
	logger.Info("SignUp")

	var (
		res      response.Build
		reqModel model.AuthSignupReq
		model    model.AuthUserData
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

	model.Email = reqModel.Email
	model.Password = reqModel.Password
	model.Name = reqModel.Name

	if err := h.service.Signup(c.Context(), &model); err != nil {
		logger.WithError(err).Errorf("unable to signup with request data %v", reqModel)
		res.Msg = err.Error()

		switch res.Msg {
		case "data is already exist":
			return res.BuildResponse(c, http.StatusUnprocessableEntity)
		default:
			return res.BuildResponse(c, http.StatusInternalServerError)
		}
	}

	res.Data = map[string]interface{}{"id": model.Id}
	return res.BuildResponse(c, http.StatusCreated)
}
