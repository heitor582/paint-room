package entities

import (
	"errors"
)

type Room struct {
	Walls [4]Wall
}

// Create new Room and returns
func NewRoom() *Room {
	return &Room{}
}

// Return quantity of liters for paint the room
func (room *Room) CalculateTotalOfLitersArea() (float64, error) {
	var sumOfWallsArea float64 = 0.0

	if err := room.IsValid(); err != nil {
		return 0.0, err
	}

	for _, wall := range room.Walls {
		sumOfWallsArea += wall.GetPaintfulArea()
	}

	liters := float64(int(sumOfWallsArea/5*100)) / 100
	return liters, nil
}

// Verify if struct room is valid
func (room *Room) IsValid() error {
	if err := room.quantityOfWalls(); err != nil {
		return err
	}

	return nil
}

func (room *Room) quantityOfWalls() error {
	var wallsEmpyt [4]Wall
	if room.Walls == wallsEmpyt {
		return errors.New("Invalid quantity of walls minumiun is 4.")
	}
	return nil
}
