package controller

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LicenseController struct {
	s *service.LicenseService
}

func NewLicenseController(s *service.LicenseService) *LicenseController {
	return &LicenseController{
		s: s,
	}
}

func (c *LicenseController) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	driverId := ctx.Params("id")

	license, err := c.s.GetByDriverId(driverId, userId)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	if license == nil {
		return response.ErrorNotFound(errors.New("license not found"))
	}
	return response.Ok(ctx, fiber.Map{
		"license": license,
	})
}

func (c *LicenseController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)

	if userId == "" {
		return response.ErrorUnauthorized(nil, "Unauthorized")
	}

	input := model.LicenseInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	// Validate required fields
	if input.LicenseNumber == "" {
		return response.ErrorBadRequest(errors.New("license number is required"))
	}

	license, err := c.s.Create(userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Created(ctx, fiber.Map{
		"license": license,
	})
}

func (c *LicenseController) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	licenseId := ctx.Params("id")

	input := model.LicenseInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	license, err := c.s.Update(licenseId, userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"license": license,
	})
}

func (c *LicenseController) Delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	licenseId := ctx.Params("id")

	license, err := c.s.Delete(licenseId, userId)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"license": license,
	})
}
