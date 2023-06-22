package app

import (
	"be_todo_app/configs"
	"be_todo_app/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func BootApp() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if portEnv := os.Getenv("PORT"); portEnv != "" {
		configs.PORT = portEnv
	}

	configs.BootDatabase()
	configs.ConnectDatabase()
	configs.RunMigration()

	app := fiber.New()

	// CORS Config
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     configs.AllowOrigins,
		AllowMethods:     configs.AllowMethods,
		AllowHeaders:     configs.AllowHeaders,
		AllowCredentials: configs.AllowCredentials,
		ExposeHeaders:    configs.ExposeHeaders,
		MaxAge:           configs.MaxAge,
	}))

	// Init route
	routes.InitRoute(app)

	app.Listen(configs.PORT)
}
