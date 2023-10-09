package operations

import (
	"backend/database"
	"backend/models"
	"github.com/gofiber/fiber/v2"
)

func AddOrganization(c * fiber.Ctx) error {
	var organizations []models.Organization
	organ := new(models.Organization)
	if err := c.BodyParser(&organ); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Find(&organizations)
	organ.ID = len(organizations) + 1
	database.DB.Db.Create(&organ)
	return c.Status(200).JSON(organ)
}

func AddUser(c * fiber.Ctx) error {
	var users []models.User
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Find(&users)
	user.ID = len(users) + 1
	database.DB.Db.Create(&user)
	return c.Status(200).JSON(user)
}

func AddEvent(c * fiber.Ctx) error {
	var events []models.Event
	event := new(models.Event)
	
	if err := c.BodyParser(&event); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	
	database.DB.Db.Find(&events)
	event.ID = len(events) + 1
	database.DB.Db.Create(&event)
	return c.Status(200).JSON(event)
}

func DeleteEvent(c * fiber.Ctx) error {
	event := []models.Event{}
	organization := new(models.Event)
	if err := c.BodyParser(organization); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Where("organization = ?", organization.Organization).Unscoped().Delete(&event)

	return c.Status(200).JSON("deleted")
}

// func UpdateUser(c * fiber.Ctx) error {
// 	users := []models.User{}
// 	event := new(models.EventType)
// 	if err := c.BodyParser(event); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}
	

// 	return c.Status(200).JSON("updated")
// }

// func UpdateOrganization(c * fiber.Ctx) error {
// 	book := []models.Book{}
// 	title := new(models.Book)
// 	if err := c.BodyParser(title); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	database.DB.Db.Model(&book).Where("title = ?", title.Title).Update("author", title.Author)

// 	return c.Status(400).JSON("updated")
// }

// func UpdateEvent(c * fiber.Ctx) error {
// 	book := []models.Book{}
// 	title := new(models.Book)
// 	if err := c.BodyParser(title); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}

// 	database.DB.Db.Model(&book).Where("title = ?", title.Title).Update("author", title.Author)

// 	return c.Status(400).JSON("updated")
// }

// func GetEvent(c * fiber.Ctx) error {
// 	events := models.Event{}
// 	 := new(models.Event)
// 	if err := c.BodyParser(); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}
// 	database.DB.Db.Where("title = ?", title.Title).Find(&)
// 	return c.Status(200).JSON(event)
// }

func GetOrganization(c * fiber.Ctx) error {
	organizations := models.Organization{}
	organization := new(models.Organization)
	if err := c.BodyParser(&organization); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Where("name = ? AND id = ?", organization.Name, organization.ID).Find(&organizations)
	return c.Status(200).JSON(organization)
}

func GetUser(c * fiber.Ctx) error {
	users := []models.User{}
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Where("name = ? AND id = ?", user.Name, user.ID).Find(&users)
	return c.Status(200).JSON(user)
}

func GetEvents(c * fiber.Ctx) error {
	event := []models.Event{}
	database.DB.Db.Find(&event)
	return c.Status(200).JSON(event)
}

func GetOrganizations(c * fiber.Ctx) error {
	organization := []models.Organization{}
	database.DB.Db.Find(&organization)

	return c.Status(200).JSON(organization)
}