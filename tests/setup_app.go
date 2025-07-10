package tests

import (
	"be-evowish/routes"
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	rootDir, err := filepath.Abs("..") // asumsi tests/ ada di bawah root
	if err != nil {
		log.Fatal("Gagal menemukan root path:", err)
	}

	envPath := filepath.Join(rootDir, ".env")
	print(envPath)
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Gagal load .env dari %s: %v", envPath, err)
	}
}

func SetupApp() *fiber.App {
	// Load ENV khusus test jika perlu
	// LoadEnv()

	// config.ConnectDatabase()
	// config.DB.Exec("DROP TABLE IF EXISTS users")
	// config.DB.AutoMigrate(&models.User{})

	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "Hello World",
		})
	})

	routes.AuthRoutes(app)
	return app
}
