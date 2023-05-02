package main

import (
	"database/sql"
	"fmt"
	"github.com/andhikasamudra/test-skyshi-digital/config"
	"github.com/andhikasamudra/test-skyshi-digital/pkg/todo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	//migrate
	migrateDSN := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?multiStatements=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DBNAME"))
	dbMigrate, _ := sql.Open("mysql", migrateDSN)
	driver, _ := migrateMysql.WithInstance(dbMigrate, &migrateMysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	m.Up()

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
