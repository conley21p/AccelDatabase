package controller

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AutoController struct {
	s *service.AutoService
}

func NewAutoController(s *service.AutoService) *AutoController {
	return &AutoController{
		s: s,
	}
}

func (c *AutoController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	auto, err := c.s.GetById(id)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	return response.Ok(ctx, fiber.Map{
		"auto": auto,
	})
}

func (c *AutoController) Create(ctx *fiber.Ctx) error {
	input := model.AutoInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	// Validate required fields
	if input.Make == "" || input.Model == "" || input.Year == 0 {
		return response.ErrorBadRequest(errors.New("make, model, and year are required"))
	}

	auto, err := c.s.Create(input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Created(ctx, fiber.Map{
		"auto": auto,
	})
}

func (c *AutoController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	input := model.AutoInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	auto, err := c.s.Update(id, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"auto": auto,
	})
}

func (c *AutoController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	auto, err := c.s.Delete(id)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"auto": auto,
	})
}
