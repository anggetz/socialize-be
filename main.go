package main

import (
	"context"
	"smart/cores"
	"smart/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New())

	route.RunRoute(app)

	dbContext := context.TODO()

	cores.PGDB = cores.ConnectToDb(dbContext)

	defer cores.PGDB.Close()

	app.Listen(":3000")
}
