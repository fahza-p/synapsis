package product

import (
	"net/http"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/response"
	"github.com/fahza-p/synapsis/lib/validation"
	"github.com/fahza-p/synapsis/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	logger := log.GetLogger(c.Context(), "Product.Handler", "Create")
	logger.Info("Create")

	var (
		res      response.Build
		reqModel model.ProductCreateReq
		model    model.Product
		authData = c.Locals("authData").(map[string]interface{})
	)

	if err := c.BodyParser(&reqModel); err != nil {
		logger.Error("Unable to parsing body data")
		res.Msg = err.Error()
		return res.BuildResponse(c, http.StatusBadRequest)
	}

	// Run Validation
	if err := validation.RunValidate(&reqModel, &res); err != nil {
		logger.Errorf("invalid request data %v", reqModel)
		return res.BuildResponse(c, http.StatusUnprocessableEntity)
	}

	if err := h.service.Create(c.Context(), &reqModel, &model, authData["email"].(string)); err != nil {
		logger.WithError(err).Errorf("unable to create category with request data %v", reqModel)
		res.Msg = err.Error()

		switch res.Msg {
		case "category not found":
			return res.BuildResponse(c, http.StatusNotFound)
		case "data is already exist":
			return res.BuildResponse(c, http.StatusUnprocessableEntity)
		default:
			return res.BuildResponse(c, http.StatusInternalServerError)
		}
	}

	res.Data = map[string]interface{}{"id": model.Id}
	return res.BuildResponse(c, http.StatusCreated)
}
