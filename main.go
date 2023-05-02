package main

import (
	"github.com/andhikasamudra/test-skyshi-digital/config"
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
)

func main() {

	db := config.GetConnection()
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("OK"))
	})
	api := app.Group("")
	todo.InitRoute(api, db)
	log.Fatal(app.Listen(":3030"))
}
