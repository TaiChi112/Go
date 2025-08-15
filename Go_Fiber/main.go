package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:200;not null"`
	Description string    `json:"description" gorm:"type:text"`
	PriceCents  int64     `json:"price_cents" gorm:"not null;default:0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
func buildDSN() string {
	host := getenv("DB_HOST", "localhost")
	port := getenv("DB_PORT", "5432")
	user := getenv("DB_USER", "postgres")
	password := getenv("DB_PASSWORD", "postgres")
	db := getenv("DB_NAME", "postgres")
	ssl := getenv("DB_SSL", "disable")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, db, ssl)
}
func connectWithRetry(dsn string, maxWait time.Duration) (*gorm.DB, error) {
	start := time.Now()
	var db *gorm.DB
	var err error
	for {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, _ := db.DB()
			if pingErr := sqlDB.Ping(); pingErr == nil {
				return db, nil
			}
			err = fmt.Errorf("ping failed: %w", err)
		}
		if time.Since(start) > maxWait {
			return nil, fmt.Errorf("failed to connect timeout: %w", err)
		}
		log.Printf("DB not ready, retrying... (%v)", err)
		time.Sleep(2 * time.Second)
	}
}
func main() {

	appPort := getenv("APP_PORT", "8080")
	dsn := buildDSN()

	db, err := connectWithRetry(dsn, 60*time.Second)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Product{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	app := fiber.New()

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	api := app.Group("/api/v1")

	api.Post("/products", func(c *fiber.Ctx) error {
		var in Product
		if err := c.BodyParser(&in); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
		}
		if in.Name == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Name is required")
		}
		if in.PriceCents <= 0 {
			return fiber.NewError(fiber.StatusBadRequest, "Price must be greater than zero")
		}
		if err := db.Create(&in).Error; err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.Status(fiber.StatusCreated).JSON(in)
	})

	api.Get("/products", func(c *fiber.Ctx) error {
		var items []Product
		if err := db.Order("id desc").Find(&items).Error; err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(items)
	})

	api.Get("/products/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var p Product
		if err := db.First(&p, "id = ?", id).Error; err != nil {
			return fiber.NewError(fiber.StatusNotFound, "not found")
		}
		return c.JSON(p)
	})

	api.Patch("/products/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var p Product
		if err := db.First(&p, "id = ?", id).Error; err != nil {
			return fiber.NewError(fiber.StatusNotFound, "not found")
		}
		var in map[string]string
		if err := c.BodyParser(&in); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid body")
		}
		if name, ok := in["name"]; ok && name != "" {
			p.Name = name
		}
		if desc, ok := in["description"]; ok {
			p.Description = desc
		}
		if ps, ok := in["price_cents"]; ok {
			if v, err := strconv.ParseInt(ps, 10, 64); err == nil && v >= 0 {
				p.PriceCents = v
			} else {
				return fiber.NewError(fiber.StatusBadRequest, "price_cents invalid")
			}
		}
		if err := db.Save(&p).Error; err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(p)
	})

	api.Delete("/products/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := db.Delete(&Product{}, "id = ?", id).Error; err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(fiber.Map{
			"message": "Product deleted successfully"})
	})

	log.Printf("Listening on port: %s", appPort)
	if err := app.Listen(":" + appPort); err != nil {
		log.Fatal(err)
	}
}
