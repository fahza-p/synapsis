package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/fahza-p/synapsis/lib/response"
	fiber "github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/mitchellh/mapstructure"
)

type TokenClaims struct {
	Raw       string                 // The raw token.  Populated when you Parse a token
	Header    map[string]interface{} // The first segment of the token
	Claims    map[string]interface{} // The second segment of the token
	Signature string                 // The third segment of the token.  Populated when you Parse a token
	Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
}

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(os.Getenv("JWT_KEY")),
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
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

func jwtSuccess(c *fiber.Ctx) error {
	var (
		tokenClaims TokenClaims
		res         response.Build
	)

	tokenClaims.jwtDataClaims(c)

	if !isAccessible(c) {
		res.Msg = "this account have no access for this resource"
		return res.BuildResponse(c, http.StatusForbidden)
	}

	return c.Next()
}

func (t *TokenClaims) jwtDataClaims(c *fiber.Ctx) {
	mapstructure.Decode(c.Locals("user"), &t)
	c.Locals("authData", t.Claims["sub"].(map[string]interface{}))
}

func isAccessible(c *fiber.Ctx) bool {
	destination := strings.Split(c.OriginalURL(), "/")
	authData := c.Locals("authData").(map[string]interface{})

	if isDestinationContainsAdmin(destination) {
		return int(authData["role"].(float64)) == 1
	}

	return true
}

func isDestinationContainsAdmin(des []string) bool {
	for _, v := range des {
		if v == "admin" {
			return true
		}
	}

	return false
}
