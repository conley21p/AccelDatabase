package middleware

import (
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == jwtware.ErrJWTMissingOrMalformed.Error() {
		return response.ErrorBadRequest(err)
	}
	return response.ErrorUnauthorized(err, "Invalid or expired token")
}

func Authenticate(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: errorHandler,
	})
}
