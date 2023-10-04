package main

import (
	"backend/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func setUpRoutes(app *fiber.App){
	app.Post("/api/organization", operations.AddOrganization)
	app.Post("/api/user", operations.AddUser)
	app.Post("/api/event", operations.AddEvent)
	app.Delete("/api/event", operations.DeleteEvent)
	app.Put("/api/user", operations.UpdateUser)
	app.Put("/api/organization", operations.UpdateOrganization)
	app.Put("/api/event", operations.UpdateEvent)
	app.Get("/api/event", operations.GetEvent)
	app.Get("/api/organization", operations.GetOrganization)
	app.Get("/api/user", operations.GetUser)
	app.Get("/api/events", operations.GetEvents)
	app.Get("/api/organizations", operations.GetOrganizations)
}

func main(){
	app := fiber.New()

	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Fatal(app.Listen(":3000"))
}