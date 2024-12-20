package controller

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	s *service.TransactionService
}

func NewTransactionController(s *service.TransactionService) *TransactionController {
	return &TransactionController{
		s: s,
	}
}

func (c *TransactionController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	transaction, err := c.s.GetById(id)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	return response.Ok(ctx, fiber.Map{
		"transaction": transaction,
	})
}

func (c *TransactionController) Create(ctx *fiber.Ctx) error {
	input := model.TransactionInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	// Validate required fields
	if input.TransportationId == "" || input.PaymentMethod == "" || input.Amount == 0 {
		return response.ErrorBadRequest(errors.New("transportation ID, payment method, and amount are required"))
	}

	transaction, err := c.s.Create(input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Created(ctx, fiber.Map{
		"transaction": transaction,
	})
}

func (c *TransactionController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	input := model.TransactionInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	transaction, err := c.s.Update(id, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"transaction": transaction,
	})
}

func (c *TransactionController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	transaction, err := c.s.Delete(id)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"transaction": transaction,
	})
}
