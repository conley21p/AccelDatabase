package controller

import (
	"time"

	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/conley21p/AccelDatabase/Server/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	s      *service.UserService
	secret string
}

func NewAuthController(s *service.UserService, secret string) *AuthController {
	return &AuthController{
		s:      s,
		secret: secret,
	}
}

func (c *AuthController) createToken(id string) (string, error) {
	claims := jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.secret))
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	input := model.AuthRegInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}

	// Validate required fields
	if input.Email == "" {
		return response.ErrorBadRequest(errors.New("email is required"))
	}

	// Check if username already exists
	exists, err := c.s.UsernameExists(input.Username)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	if exists {
		return response.ErrorBadRequest(errors.New("username already exists"))
	}

	password, err := util.HashPassword(input.Password)
	if err != nil {
		return response.ErrorBadRequest(err)
	}

	user, err := c.s.Create(input.Username, input.Email, password)
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	token, err := c.createToken(user.Id)
	if err != nil {
		return response.ErrorUnauthorized(err, "Registration error")
	}

	return response.Created(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}
	user, err := c.s.GetIDByUsername(input.Username)

	if err != nil {
		return response.ErrorUnauthorized(err, "Login error 1")
	}

	if !util.CheckPassword(input.Password, user.Password) {
		return response.ErrorUnauthorized(err, "Login error 2")
	}
	println("Create token ID: " + user.Id)
	token, err := c.createToken(user.Id)
	if err != nil {
		return response.ErrorUnauthorized(err, "Login error 3")
	}
	return response.Ok(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c *AuthController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)
	println("id" + id)
	currentUser, err := c.s.GetById(id)
	if err != nil {
		return response.ErrorUnauthorized(err, "Invalid credentials")
	}
	token, err := c.createToken(currentUser.Username)
	if err != nil {
		return response.ErrorUnauthorized(err, "Invalid credentials")
	}
	return response.Ok(ctx, fiber.Map{
		"user":  currentUser,
		"token": token,
	})
}
