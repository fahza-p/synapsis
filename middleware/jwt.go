package middleware

import (
	"net/http"
	"os"

	"github.com/fahza-p/synapsis/lib/response"
	fiber "github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_KEY")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	var res response.Build

	switch err.Error() {
	case "Missing or malformed JWT":
		res.Msg = "Missing or malformed JWT"
		return res.BuildResponse(c, http.StatusBadRequest)
	case "Token is expired":
		res.Msg = "Token is expired"
		return res.BuildResponse(c, http.StatusUnauthorized)
	default:
		res.Msg = "Invalid token"
		return res.BuildResponse(c, http.StatusUnauthorized)
	}
}
