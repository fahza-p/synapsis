package product

import (
	"net/http"
	"reflect"

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

	cartItem, totalData, err := h.service.GetList(c.Context(), &queryParams)
	if err != nil {
		logger.WithError(err).Error("can't cart")
		res.Msg = err.Error()
		if res.Msg == "document not found" {
			return res.BuildResponse(c, http.StatusNotFound)
		}

		return res.BuildResponse(c, http.StatusInternalServerError)
	}

	res.Data = buildListResponse(&queryParams, cartItem, totalData)
	return res.BuildResponse(c, http.StatusOK)
}

/* Local Functions */
func buildListResponse(q *store.QueryParams, items interface{}, total int64) map[string]interface{} {
	result := map[string]interface{}{
		"data":  make([]string, 0),
		"query": q.BuildQueryResponse(total),
	}

	if reflect.TypeOf(items).Kind() == reflect.Slice {
		if reflect.ValueOf(items).Len() > 0 {
			result["data"] = items
		}
	}

	return result
}
