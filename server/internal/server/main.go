package server

import (
	"github.com/conley21p/AccelDatabase/Server/internal/config"
	"github.com/conley21p/AccelDatabase/Server/internal/controller"
	"github.com/conley21p/AccelDatabase/Server/internal/database"
	"github.com/conley21p/AccelDatabase/Server/internal/server/router/response"
	"github.com/conley21p/AccelDatabase/Server/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	app       *fiber.App
	port      string
	jwtSecret string
	db        *sqlx.DB
}

func NewServer(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: response.DefaultErrorHandler,
	})

	app.Use(cors.New())

	port := ":" + cfg.Port
	db := database.Connect(cfg.DatabaseUrl)

	return &Server{
		app:       app,
		port:      port,
		jwtSecret: cfg.JwtSecret,
		db:        db,
	}
}

func (s *Server) Start() error {
	us := service.NewUserService(s.db)
	ds := service.NewDriverService(s.db)
	contactService := service.NewContactInfoService(s.db)
	licenseService := service.NewLicenseService(s.db)
	trailerService := service.NewTrailerService(s.db)
	haulerService := service.NewHaulerService(s.db)
	vehicleService := service.NewVehicleService(s.db)
	autoService := service.NewAutoService(s.db)
	boatService := service.NewBoatService(s.db)
	transactionService := service.NewTransactionService(s.db)
	transportationService := service.NewTransportationService(s.db)
	ownerService := service.NewOwnerService(s.db)

	uc := controller.NewAuthController(us, s.jwtSecret)
	dc := controller.NewDriverController(ds)
	contactController := controller.NewContactInfoController(contactService)
	licenseController := controller.NewLicenseController(licenseService)
	trailerController := controller.NewTrailerController(trailerService)
	haulerController := controller.NewHaulerController(haulerService)
	vehicleController := controller.NewVehicleController(vehicleService)
	autoController := controller.NewAutoController(autoService)
	boatController := controller.NewBoatController(boatService)
	transactionController := controller.NewTransactionController(transactionService)
	transportationController := controller.NewTransportationController(transportationService)
	ownerController := controller.NewOwnerController(ownerService)

	s.SetupRoutes(uc,
		dc,
		contactController,
		licenseController,
		trailerController,
		haulerController,
		vehicleController,
		autoController,
		boatController,
		transactionController,
		transportationController,
		ownerController)

	return s.app.Listen(s.port)
}

func (s *Server) Stop() error {
	s.db.Close()
	return s.app.Shutdown()
}
