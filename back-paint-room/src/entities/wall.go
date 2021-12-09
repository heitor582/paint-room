package entities

import (
	"errors"

	"gitlab.com/heitor582/paint_room/src/constants"
)

type Wall struct {
	DoorsNumber    int
	WindownsNumber int
	Height         float64
	Width          float64
}

// Create new Wall and returns
func NewWall(DoorsNumber int, WindownsNumber int, Height float64, Width float64) (*Wall, error) {
	wall := Wall{
		DoorsNumber:    DoorsNumber,
		WindownsNumber: WindownsNumber,
		Height:         Height,
		Width:          Width,
	}

	err := wall.IsValid()

	if err != nil {
		return nil, err
	}

	return &wall, nil
}

// Return Paintful area of wall
func (wall *Wall) GetPaintfulArea() float64 {
	AreaDoors := float64(wall.DoorsNumber) * constants.DoorMeasure
	AreaWindowns := float64(wall.WindownsNumber) * constants.WindowMeasure
	return (wall.Width * wall.Height) - (AreaDoors + AreaWindowns)
}

// Verify if struct wall is valid
func (wall *Wall) IsValid() error {
	if err := wall.validateWallArea(); err != nil {
		return err
	}

	if err := wall.validateSpaceOfWindownsAndDoors(); err != nil {
		return err
	}

	if err := wall.validateSpaceBetweenDoorAndWall(); err != nil {
		return err
	}

	return nil
}

func (wall *Wall) validateWallArea() error {
	wallArea := wall.Height * wall.Width
	if wallArea < 1.0 || wallArea > 15.0 {
		return errors.New("Invalid amount of square meters for the wall.")
	}
	return nil
}

func (wall *Wall) validateSpaceBetweenDoorAndWall() error {
	if wall.DoorsNumber > 0 && wall.Height-constants.DoorHeight < 0.3 {
		return errors.New("The height of walls with a door must be at least 30 centimeters greater than the height of the door.")
	}
	return nil
}

func (wall *Wall) validateSpaceOfWindownsAndDoors() error {
	AreaDoors := float64(wall.DoorsNumber) * constants.DoorMeasure
	AreaWindowns := float64(wall.WindownsNumber) * constants.WindowMeasure
	AreaWall := wall.Height * wall.Width
	if AreaDoors+AreaWindowns > AreaWall {
		return errors.New("Doors and windows occupy more than 50% of the wall area.")
	}
	return nil
}
