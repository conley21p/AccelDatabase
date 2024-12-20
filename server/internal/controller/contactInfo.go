package controller

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ContactInfoController struct {
	s *service.ContactInfoService
}

func NewContactInfoController(s *service.ContactInfoService) *ContactInfoController {
	return &ContactInfoController{
		s: s,
	}
}

func (c *ContactInfoController) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	driverId := ctx.Params("id")

	contactInfo, err := c.s.GetByDriverId(driverId, userId)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	if contactInfo == nil {
		return response.ErrorNotFound(errors.New("contact info not found"))
	}
	return response.Ok(ctx, fiber.Map{
		"contactInfo": contactInfo,
	})
}

func (c *ContactInfoController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)

	if userId == "" {
		return response.ErrorUnauthorized(nil, "Unauthorized")
	}

	input := model.ContactInfoInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	// Validate required fields
	if input.PhoneNumber == "" || input.StreetAddress == "" {
		return response.ErrorBadRequest(errors.New("phone number and street address are required"))
	}

	contactInfo, err := c.s.Create(userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Created(ctx, fiber.Map{
		"contactInfo": contactInfo,
	})
}

func (c *ContactInfoController) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	contactId := ctx.Params("id")

	input := model.ContactInfoInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	contactInfo, err := c.s.Update(contactId, userId, input)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"contactInfo": contactInfo,
	})
}

func (c *ContactInfoController) Delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	contactId := ctx.Params("id")

	contactInfo, err := c.s.Delete(contactId, userId)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	return response.Ok(ctx, fiber.Map{
		"contactInfo": contactInfo,
	})
}
