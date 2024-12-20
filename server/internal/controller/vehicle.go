package controller

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/gofiber/fiber/v2"
)

type VehicleController struct {
	s *service.VehicleService
}

func NewVehicleController(s *service.VehicleService) *VehicleController {
	return &VehicleController{
		s: s,
	}
}

func (c *VehicleController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	vehicle, err := c.s.GetById(id)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	return response.Ok(ctx, fiber.Map{
		"vehicle": vehicle,
	})
}

func (c *VehicleController) Create(ctx *fiber.Ctx) error {
	input := model.VehicleInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	// Validate required fields
	if input.Length == 0 || input.Width == 0 || input.Height == 0 {
		return response.ErrorBadRequest(errors.New("dimensions are required"))
	}

	vehicle, err := c.s.Create(input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Created(ctx, fiber.Map{
		"vehicle": vehicle,
	})
}

func (c *VehicleController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	input := model.VehicleInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	vehicle, err := c.s.Update(id, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"vehicle": vehicle,
	})
}

func (c *VehicleController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	vehicle, err := c.s.Delete(id)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"vehicle": vehicle,
	})
}
