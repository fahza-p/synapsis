package response

import (
	"net/http"
	"reflect"

	"github.com/fahza-p/synapsis/lib/store"
	"github.com/gofiber/fiber/v2"
)

type Build struct {
	Status string            `json:"status"`
	Msg    string            `json:"message"`
	Data   interface{}       `json:"data"`
	Err    map[string]string `json:"error"`
}

func (b *Build) BuildResponse(c *fiber.Ctx, code int) error {
	b.setDefaultValue(code)
	b.setResponseStatus(code)

	return c.Status(code).JSON(b)
}

func (b *Build) setDefaultValue(c int) {
	if b.Err == nil {
		b.Err = make(map[string]string)
	}

	if b.Data == nil {
		b.Data = make(map[string]interface{})
	}

	b.setMessage(c)
}

func (b *Build) setResponseStatus(c int) {
	b.Status = http.StatusText(c)
}

func (b *Build) setMessage(c int) {
	for c >= 10 {
		c = c / 10
	}

	if c != 2 {
		if b.Msg != "" && (b.Err == nil || len(b.Err) == 0) {
			b.Err = map[string]string{"msg": b.Msg}
		}
	}
}

func BuildListResponse(q *store.QueryParams, items interface{}, total int64) map[string]interface{} {
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
