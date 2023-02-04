package cart

import (
	"net/http"
	"strconv"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/response"
	"github.com/fahza-p/synapsis/lib/validation"
	"github.com/fahza-p/synapsis/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) AddProduct(c *fiber.Ctx) error {
	logger := log.GetLogger(c.Context(), "Cart.Handler", "AddProduct")
	logger.Info("AddProduct")

	var (
		res      response.Build
		reqModel model.CartAddItemReq
		model    model.CartItemData
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

	model.ProductId = reqModel.ProductId
	model.Qty = reqModel.Qty

	if err := h.service.AddProduct(c.Context(), &model, authData); err != nil {
		logger.WithError(err).Errorf("unable to add item to cart with product id %v", reqModel)
		res.Msg = err.Error()

		switch res.Msg {
		case "product not found", "document not found":
			return res.BuildResponse(c, http.StatusNotFound)
		default:
			return res.BuildResponse(c, http.StatusInternalServerError)
		}
	}

	res.Data = map[string]interface{}{"id": model.Id}
	return res.BuildResponse(c, http.StatusCreated)
}

func (h *Handler) RemoveProduct(c *fiber.Ctx) error {
	logger := log.GetLogger(c.Context(), "Cart.Handler", "RemoveProduct")
	logger.Info("RemoveProduct")

	var (
		res      response.Build
		id       = c.Params("id")
		authData = c.Locals("authData").(map[string]interface{})
	)

	productId, err := strconv.Atoi(id)
	if err != nil {
		return res.BuildResponse(c, http.StatusBadRequest)
	}

	if err := h.service.RemoveProduct(c.Context(), int64(productId), authData); err != nil {
		logger.WithError(err).Errorf("unable to remove cart item with product id %s", id)
		res.Msg = err.Error()

		switch res.Msg {
		case "document not found":
			return res.BuildResponse(c, http.StatusNotFound)
		case "can't delete this category as long as product still uses it":
			return res.BuildResponse(c, http.StatusUnprocessableEntity)
		default:
			return res.BuildResponse(c, http.StatusInternalServerError)
		}
	}

	return res.BuildResponse(c, http.StatusOK)
}
