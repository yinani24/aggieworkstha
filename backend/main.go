package main

import (
	"backend/database"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"backend/operations"
)

func setUpRoutes(app *fiber.App){
	app.Post("/api/organization", operations.AddOrganization)
	app.Post("/api/user", operations.AddUser)
	app.Post("/api/event", operations.AddEvent)
	app.Delete("/api/event", operations.DeleteEvent)
	// app.Put("/api/user", operations.UpdateUser)
	// app.Put("/api/organization", operations.UpdateOrganization)
	// app.Put("/api/event", operations.UpdateEvent)
	// app.Get("/api/event", operations.GetEvent)
	app.Get("/api/organization", operations.GetOrganization)
	app.Get("/api/user", operations.GetUser)
	// app.Get("/api/events", operations.GetEvents)
	app.Get("/api/organizations", operations.GetOrganizations)
}

func main(){
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://localhost:3000",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))
	
	setUpRoutes(app)	

	database.ConnectDB()
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Db.Close()
	log.Fatal(app.Listen(":4000"))
}