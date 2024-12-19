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
	usrController *controller.AuthController,
	drvController *controller.DriverController,
	contactController *controller.ContactInfoController,
) {
	api := s.app.Group("/api")

	// ********************************************
	// 		Login Request
	// ********************************************
	api.Get("/", healthCheck(s.db))
	api.Post("/login", usrController.Login)
	api.Post("/register", usrController.Register)

	// ********************************************
	// 		Driver Request
	// ********************************************
	drivers := api.Group("/driver")
	drivers.Use(middleware.Authenticate(s.jwtSecret))
	drivers.Get("/", drvController.Get)
	drivers.Post("/register", drvController.Create)
	drivers.Post("/ContactInfo", contactController.Create)
	// drivers.Post("/Hauler", dc.CreateHauler)
	// drivers.Post("/Trailer", dc.CreateTrailer)
	// drivers.Post("/Insurance", dc.CreateInsurance)
	// drivers.Post("/License", dc.CreateLicense)

	// Transportation Request

	// ********************************************
	// 		Buyer Request
	// ********************************************
	buyers := api.Group("/buyer")
	buyers.Use(middleware.Authenticate(s.jwtSecret))

	// Transportation Request

}
