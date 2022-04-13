package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/unit2022-bosch/teapot/backend/internal/app/orm"
	"github.com/unit2022-bosch/teapot/backend/internal/services/journeys"
	"log"
	"os"
	"strconv"
	"time"
)

type App struct {
	port int
	api  *fiber.App
}

func NewApp(config *webConfig, router *Router, migrator *orm.Migrator, worker journeys.IJourneyWorker) *App {
	err := migrator.Seed()
	if err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}

	log.Println("Start worker")
	err = worker.Start()
	if err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}

	log.Println("Starting webserver...")

	api := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	})

	api.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	api.Use(logger.New(logger.Config{
		Format:     "${time} [${method}] ${path} ${status} - ${latency}\n",
		Output:     log.Writer(),
		TimeFormat: "2006/01/02 15:04:05",
	}))
	api.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	router.Setup(api)

	return &App{
		api:  api,
		port: config.httpPort,
	}
}

func (app *App) Start() error {
	log.Println("Listening on :" + strconv.Itoa(app.port))
	log.Println("PORT: " + os.Getenv("PORT"))
	return app.api.Listen(":" + strconv.Itoa(app.port))
}
