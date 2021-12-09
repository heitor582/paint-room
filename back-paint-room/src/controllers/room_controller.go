package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/heitor582/paint_room/src/dtos"
	"gitlab.com/heitor582/paint_room/src/services"
)

//Returns to client how many
func GetListOfCanOfPaint(c *fiber.Ctx) error {
	roomDto := new(dtos.RoomDtoInput)
	if err := c.BodyParser(roomDto); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return nil
	}

	response, err := services.GetListOfCanOfPaint(*roomDto)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return nil
	}
	c.JSON(response)
	return nil
}
