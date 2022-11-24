package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	_ "github.com/lib/pq"
	"learn-go/controller"
	"log"
	"os"
)

func main() {
	conStr := "postgresql://root:secret@localhost/app?sslmode=disable"
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Fatal(err)
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/accounts", func(c *fiber.Ctx) error {
		return controller.GetHandler(c, db)
	})
	app.Post("/accounts", func(c *fiber.Ctx) error {
		return controller.PostHandler(c, db)
	})
	app.Put("/accounts", func(c *fiber.Ctx) error {
		return controller.PutHandler(c, db)
	})
	app.Delete("/accounts", func(c *fiber.Ctx) error {
		return controller.DeleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Static("/", "./views/public")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
