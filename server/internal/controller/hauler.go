package controller

import (
	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type HaulerController struct {
	s *service.HaulerService
}

func NewHaulerController(s *service.HaulerService) *HaulerController {
	return &HaulerController{
		s: s,
	}
}

// Get retrieves a hauler by driver ID
func (c *HaulerController) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	driverId := ctx.Params("id")

	hauler, err := c.s.GetByDriverId(driverId, userId)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	return response.Ok(ctx, fiber.Map{
		"hauler": hauler,
	})
}

// Create creates a new hauler
func (c *HaulerController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)

	if userId == "" {
		return response.ErrorUnauthorized(nil, "Unauthorized")
	}

	input := model.HaulerInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	hauler, err := c.s.Create(userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Created(ctx, fiber.Map{
		"hauler": hauler,
	})
}

// Update updates an existing hauler
func (c *HaulerController) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	haulerId := ctx.Params("id")

	input := model.HaulerInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	hauler, err := c.s.Update(haulerId, userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Ok(ctx, fiber.Map{
		"hauler": hauler,
	})
}

// Delete deletes a hauler
func (c *HaulerController) Delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	haulerId := ctx.Params("id")

	hauler, err := c.s.Delete(haulerId, userId)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Ok(ctx, fiber.Map{
		"hauler": hauler,
	})
}

// GetAll retrieves all haulers for a driver
func (c *HaulerController) GetAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)

	haulers, err := c.s.GetAll(userId)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	return response.Ok(ctx, fiber.Map{
		"haulers": haulers,
	})
}
