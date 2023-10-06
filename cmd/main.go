package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/useresd/golang-fiber-starter/internal/database"
	"github.com/useresd/golang-fiber-starter/internal/handlers"
	"github.com/useresd/golang-fiber-starter/internal/repositories"
	"github.com/useresd/golang-fiber-starter/internal/services"
)

var db *sql.DB

func main() {

	app := fiber.New()

	app.Use(logger.New())

	app.Use(func(c *fiber.Ctx) error {
		c.SetUserContext(database.SetContextDB(c.Context(), db))
		return c.Next()
	})

	userRepository := repositories.NewDefaultUserRepository()
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	app.Post("/users", userHandler.HandlePostUser)
	app.Get("/users", userHandler.HandleGetUsers)

	if err := app.Listen(":3003"); err != nil {
		log.Fatal(err)
	}

}

func init() {

	var err error

	if db != nil {
		return
	}

	db, err = sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			"root",
			"password",
			"localhost",
			"3306",
			"mygotest",
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

}
