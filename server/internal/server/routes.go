package server

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/controller"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/middleware"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func healthCheck(db *sqlx.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var result int
		err := db.Get(&result, "select 1")
		if err != nil {
			return errors.New("database unavailable")
		}
		return response.Ok(ctx, fiber.Map{
			"database": "available",
		})
	}
}

func (s *Server) SetupRoutes(
	uc *controller.AuthController,
	// cc *controller.CategoryController,
	tc *controller.TransactionController,
	dc *controller.DriverController,
) {
	api := s.app.Group("/api")

	api.Get("/", healthCheck(s.db))

	api.Post("/login", uc.Login)
	api.Post("/register", uc.Register)
	api.Get("/me", middleware.Authenticate(s.jwtSecret), uc.Me)

	drivers := api.Group("/Driver")
	drivers.Use(middleware.Authenticate(s.jwtSecret))
	// categories.Get("/", cc.List)
	// categories.Post("/", cc.Create)
	// categories.Get("/:id", cc.Get)
	// categories.Put("/:id", cc.Update)
	// categories.Delete("/:id", cc.Delete)
	drivers.Get("/", dc.Get)
	//categories.Get("/",)

	// transactions := api.Group("/transaction")
	// transactions.Use(middleware.Authenticate(s.jwtSecret))
	// transactions.Get("/", tc.GetAll)
	// transactions.Post("/", tc.Create)
	// transactions.Get("/:id", tc.Get)
	// transactions.Put("/:id", tc.Update)
	// transactions.Delete("/:id", tc.Delete)
}
