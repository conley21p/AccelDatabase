package controller

import (
	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type DriverController struct {
	s *service.DriverService
}

func NewDriverController(s *service.DriverService) *DriverController {
	return &DriverController{
		s: s,
	}
}

func (c *DriverController) Get(ctx *fiber.Ctx) error {
	//Authenticate request
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)

	driver, err := c.s.GetByUserId(userId)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	if driver.UserId != userId {
		return response.ErrorUnauthorized(nil, "unauthorized")
	}
	return response.Ok(ctx, fiber.Map{
		"driver": driver,
	})
}

func (c *DriverController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	println()
	if userId == "" {
		return response.ErrorUnauthorized(nil, "Unauthorized")
	}
	input := model.DriverInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	driver, err := c.s.Create(userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Created(ctx, fiber.Map{
		"driver": driver,
	})
}

// func (c *DriverController) Update(ctx *fiber.Ctx) error {
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

// func (c *DriverController) Delete(ctx *fiber.Ctx) error {
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
