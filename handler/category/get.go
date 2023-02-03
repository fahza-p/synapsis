package category

import (
	"net/http"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/response"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) FindById(c *fiber.Ctx) error {
	logger := log.GetLogger(c.Context(), "Category.Handler", "FindById")
	logger.Info("FindById")

	var (
		res response.Build
		id  = c.Params("id")
	)

	categoryData, err := h.service.FindById(c.Context(), id)
	if err != nil {
		logger.WithError(err).Errorf("can't category with Id: %s", id)
		res.Msg = err.Error()
		if res.Msg == "document not found" {
			return res.BuildResponse(c, http.StatusNotFound)
		}

		return res.BuildResponse(c, http.StatusInternalServerError)
	}

	res.Data = categoryData
	return res.BuildResponse(c, http.StatusOK)
}