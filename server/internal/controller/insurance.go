package controller

import (
	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type InsuranceController struct {
	s *service.InsuranceService
}

func NewInsuranceController(s *service.InsuranceService) *InsuranceController {
	return &InsuranceController{
		s: s,
	}
}

func (c *InsuranceController) Get(ctx *fiber.Ctx, driverId string) error {
	//Authenticate request
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)

	insurance, err := c.s.GetByDriverId(driverId, userId)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	return response.Ok(ctx, fiber.Map{
		"insurance": insurance,
	})
}

func (c *InsuranceController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)

	if userId == "" {
		return response.ErrorUnauthorized(nil, "Unauthorized")
	}
	input := model.InsuranceInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	insurance, err := c.s.Create(userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Created(ctx, fiber.Map{
		"insurance": insurance,
	})
}

// func (c *InsuranceController) Update(ctx *fiber.Ctx) error {
// 	user := ctx.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	userId := claims["sub"].(string)
// 	id := ctx.Params("id")
// 	input := model.CategoryInput{}
// 	if err := ctx.BodyParser(&input); err != nil {
// 		return response.ErrorBadRequest(err)
// 	}
// 	category, err := c.s.Update(id, userId, input.Title)
// 	if err != nil {
// 		return response.ErrorBadRequest(err)
// 	}
// 	return response.Ok(ctx, fiber.Map{
// 		"category": category,
// 	})
// }

// func (c *InsuranceController) Delete(ctx *fiber.Ctx) error {
// 	user := ctx.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	userId := claims["sub"].(string)
// 	id := ctx.Params("id")
// 	category, err := c.s.Delete(id, userId)
// 	if err != nil {
// 		return response.ErrorBadRequest(err)
// 	}
// 	return response.Ok(ctx, fiber.Map{
// 		"category": category,
// 	})
// }
