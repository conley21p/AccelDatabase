package controller

import (
	"errors"

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
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)

	driver, err := c.s.GetByUserId(userId)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	if driver == nil || driver.UserId != userId {
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

func (c *DriverController) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	driverId := ctx.Params("id")

	input := model.DriverInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	driver, err := c.s.Update(driverId, userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Ok(ctx, fiber.Map{
		"driver": driver,
	})
}

func (c *DriverController) Delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	driverId := ctx.Params("id")

	driver, err := c.s.Delete(driverId, userId)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Ok(ctx, fiber.Map{
		"driver": driver,
	})
}

func (c *DriverController) GetById(ctx *fiber.Ctx) error {
	driverId := ctx.Params("id")

	driver, err := c.s.GetById(driverId)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	return response.Ok(ctx, fiber.Map{
		"driver": driver,
	})
}

func (c *DriverController) CreateWithDetails(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)

	if userId == "" {
		return response.ErrorUnauthorized(nil, "Unauthorized")
	}

	input := model.DriverRegistrationInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	// Validate required fields
	if input.Driver.FirstName == "" || input.Driver.LastName == "" {
		return response.ErrorBadRequest(errors.New("first name and last name are required"))
	}

	driver, err := c.s.CreateWithDetails(userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Created(ctx, fiber.Map{
		"driver": driver,
	})
}

func (c *DriverController) GetWithDetails(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)

	driver, err := c.s.GetByUserId(userId)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	if driver == nil || driver.UserId != userId {
		return response.ErrorUnauthorized(nil, "unauthorized")
	}

	// Fetch all details
	driver, err = c.s.GetWithDetails(driver.Id)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"driver": driver,
	})
}
