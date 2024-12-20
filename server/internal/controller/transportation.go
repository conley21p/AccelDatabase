package controller

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/gofiber/fiber/v2"
)

type TransportationController struct {
	s *service.TransportationService
}

func NewTransportationController(s *service.TransportationService) *TransportationController {
	return &TransportationController{
		s: s,
	}
}

func (c *TransportationController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	transportation, err := c.s.GetById(id)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	return response.Ok(ctx, fiber.Map{
		"transportation": transportation,
	})
}

func (c *TransportationController) Create(ctx *fiber.Ctx) error {
	input := model.TransportationInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	// Validate required fields
	if input.Description == "" || input.TransportDate == "" ||
		input.PickupAddress == "" || input.DeliveryAddress == "" {
		return response.ErrorBadRequest(errors.New("description, dates, and addresses are required"))
	}

	transportation, err := c.s.Create(input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Created(ctx, fiber.Map{
		"transportation": transportation,
	})
}

func (c *TransportationController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	input := model.TransportationInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	transportation, err := c.s.Update(id, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"transportation": transportation,
	})
}

func (c *TransportationController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	transportation, err := c.s.Delete(id)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"transportation": transportation,
	})
}
