package controller

import (
	"log"
	"workspace_booking/model"

	"github.com/gofiber/fiber/v2"
)

// Index
func AllBuildings(c *fiber.Ctx) error {
	buildings := model.GetAllBuildings()
	if len(buildings) != 0 {
		if err := c.JSON(&fiber.Map{
			"success":   true,
			"buildings": buildings,
			"message":   "All Buildings returned successfully",
		}); err != nil {
			log.Println(3, err)
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}
	} else {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "No Records found for Building",
		})
	}
	return nil
}

// CreateBuilding handler
func CreateBuilding(c *fiber.Ctx) error {
	building := new(model.Building)
	if err := c.BodyParser(building); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	err := building.CreateBuilding()
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	// Return Building in JSON format
	if err := c.JSON(&fiber.Map{
		"success":  true,
		"building": building,
		"message":  "Building successfully created",
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating Building",
		})
	}
	return nil
}
