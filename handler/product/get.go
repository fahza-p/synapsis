package product

import (
	"net/http"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/response"
	"github.com/fahza-p/synapsis/lib/store"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetList(c *fiber.Ctx) error {
	logger := log.GetLogger(c.Context(), "Product.Handler", "GetList")
	logger.Info("GetList")

	var (
		res         response.Build
		queryParams store.QueryParams
	)

	c.QueryParser(&queryParams)

	items, totalData, err := h.service.GetList(c.Context(), &queryParams)
	if err != nil {
		logger.WithError(err).Error("can't cart")
		res.Msg = err.Error()
		if res.Msg == "document not found" {
			return res.BuildResponse(c, http.StatusNotFound)
		}

		return res.BuildResponse(c, http.StatusInternalServerError)
	}

	res.Data = response.BuildListResponse(&queryParams, items, totalData)
	return res.BuildResponse(c, http.StatusOK)
}
