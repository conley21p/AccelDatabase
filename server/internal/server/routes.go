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
	licenseController *controller.LicenseController,
	trailerController *controller.TrailerController,
	haulerController *controller.HaulerController,
	vehicleController *controller.VehicleController,
	autoController *controller.AutoController,
	boatController *controller.BoatController,
	transactionController *controller.TransactionController,
	transportationController *controller.TransportationController,
	ownerController *controller.OwnerController,
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

	// Driver CRUD operations
	drivers.Get("/", drvController.Get)                                // Get current driver
	drivers.Get("/:id", drvController.GetById)                         // Get specific driver
	drivers.Post("/register", drvController.Create)                    // Create new driver
	drivers.Put("/:id", drvController.Update)                          // Update driver
	drivers.Delete("/:id", drvController.Delete)                       // Delete driver
	drivers.Post("/register/details", drvController.CreateWithDetails) // Add this new route
	drivers.Get("/details", drvController.GetWithDetails)              // Get current driver with all details

	// Contact Info operations
	drivers.Get("/contact/:id", contactController.Get)
	drivers.Post("/contact", contactController.Create)
	drivers.Put("/contact/:id", contactController.Update)
	drivers.Delete("/contact/:id", contactController.Delete)

	// License operations
	drivers.Get("/license/:id", licenseController.Get)
	drivers.Post("/license", licenseController.Create)
	drivers.Put("/license/:id", licenseController.Update)
	drivers.Delete("/license/:id", licenseController.Delete)

	// Hauler operations
	drivers.Get("/hauler/:id", haulerController.Get)
	drivers.Post("/hauler", haulerController.Create)
	drivers.Put("/hauler/:id", haulerController.Update)
	drivers.Delete("/hauler/:id", haulerController.Delete)

	// Trailer operations
	drivers.Get("/trailer/:id", trailerController.Get)
	drivers.Post("/trailer", trailerController.Create)
	drivers.Put("/trailer/:id", trailerController.Update)
	drivers.Delete("/trailer/:id", trailerController.Delete)

	// Vehicle operations
	api.Get("/vehicle/:id", vehicleController.Get)
	api.Post("/vehicle", vehicleController.Create)
	api.Put("/vehicle/:id", vehicleController.Update)
	api.Delete("/vehicle/:id", vehicleController.Delete)

	// Auto operations
	api.Get("/auto/:id", autoController.Get)
	api.Post("/auto", autoController.Create)
	api.Put("/auto/:id", autoController.Update)
	api.Delete("/auto/:id", autoController.Delete)

	// Boat operations
	api.Get("/boat/:id", boatController.Get)
	api.Post("/boat", boatController.Create)
	api.Put("/boat/:id", boatController.Update)
	api.Delete("/boat/:id", boatController.Delete)

	// Transaction operations
	api.Get("/transaction/:id", transactionController.Get)
	api.Post("/transaction", transactionController.Create)
	api.Put("/transaction/:id", transactionController.Update)
	api.Delete("/transaction/:id", transactionController.Delete)

	// Transportation operations
	api.Get("/transportation/:id", transportationController.Get)
	api.Post("/transportation", transportationController.Create)
	api.Put("/transportation/:id", transportationController.Update)
	api.Delete("/transportation/:id", transportationController.Delete)

	// ********************************************
	// 		Owner Request
	// ********************************************
	owners := api.Group("/owner")
	owners.Use(middleware.Authenticate(s.jwtSecret))

	// Transportation Request
	owners.Get("/", ownerController.Get)
	owners.Post("/register", ownerController.Create)
	owners.Put("/:id", ownerController.Update)
	owners.Delete("/:id", ownerController.Delete)
}
