package cart

import (
	"net/http"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/response"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetCartMe(c *fiber.Ctx) error {
	logger := log.GetLogger(c.Context(), "Cart.Handler", "GetCartMe")
	logger.Info("GetCartMe")

	var (
		res      response.Build
		authData = c.Locals("authData").(map[string]interface{})
	)

	cartData, err := h.service.GetCartMe(c.Context(), authData)
	if err != nil {
		logger.WithError(err).Error("can't cart")
		res.Msg = err.Error()
		if res.Msg == "document not found" {
			return res.BuildResponse(c, http.StatusNotFound)
		}

		return res.BuildResponse(c, http.StatusInternalServerError)
	}

	res.Data = cartData
	return res.BuildResponse(c, http.StatusOK)
}

func (h *Handler) GetCartItemMe(c *fiber.Ctx) error {
	logger := log.GetLogger(c.Context(), "Cart.Handler", "GetCartItemMe")
	logger.Info("GetCartItemMe")

	var (
		res      response.Build
		authData = c.Locals("authData").(map[string]interface{})
	)

	cartItem, err := h.service.GetCartItemMe(c.Context(), authData)
	if err != nil {
		logger.WithError(err).Error("can't cart")
		res.Msg = err.Error()
		if res.Msg == "document not found" {
			return res.BuildResponse(c, http.StatusNotFound)
		}

		return res.BuildResponse(c, http.StatusInternalServerError)
	}

	res.Data = cartItem
	return res.BuildResponse(c, http.StatusOK)
}
