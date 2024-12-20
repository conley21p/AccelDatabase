package controller

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TrailerController struct {
	s *service.TrailerService
}

func NewTrailerController(s *service.TrailerService) *TrailerController {
	return &TrailerController{
		s: s,
	}
}

func (c *TrailerController) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	trailerId := ctx.Params("id")

	trailer, err := c.s.GetById(trailerId, userId)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	if trailer == nil {
		return response.ErrorNotFound(errors.New("trailer not found"))
	}
	return response.Ok(ctx, fiber.Map{
		"trailer": trailer,
	})
}

func (c *TrailerController) Create(ctx *fiber.Ctx) error {
	input := model.TrailerInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	// Validate required fields
	if input.Type == "" {
		return response.ErrorBadRequest(errors.New("trailer type is required"))
	}

	trailer, err := c.s.Create(input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Created(ctx, fiber.Map{
		"trailer": trailer,
	})
}

func (c *TrailerController) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	trailerId := ctx.Params("id")

	input := model.TrailerInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	trailer, err := c.s.Update(trailerId, userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"trailer": trailer,
	})
}

func (c *TrailerController) Delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	trailerId := ctx.Params("id")

	trailer, err := c.s.Delete(trailerId, userId)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"trailer": trailer,
	})
}
