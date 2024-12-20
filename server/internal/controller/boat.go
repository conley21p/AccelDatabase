package controller

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/gofiber/fiber/v2"
)

type BoatController struct {
	s *service.BoatService
}

func NewBoatController(s *service.BoatService) *BoatController {
	return &BoatController{
		s: s,
	}
}

func (c *BoatController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	boat, err := c.s.GetById(id)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	return response.Ok(ctx, fiber.Map{
		"boat": boat,
	})
}

func (c *BoatController) Create(ctx *fiber.Ctx) error {
	input := model.BoatInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	// Validate required fields
	if input.Make == "" || input.Model == "" || input.Year == 0 {
		return response.ErrorBadRequest(errors.New("make, model, and year are required"))
	}

	boat, err := c.s.Create(input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Created(ctx, fiber.Map{
		"boat": boat,
	})
}

func (c *BoatController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	input := model.BoatInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	boat, err := c.s.Update(id, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"boat": boat,
	})
}

func (c *BoatController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	boat, err := c.s.Delete(id)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"boat": boat,
	})
}
