package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallArea(t *testing.T) {
	_, err := NewWall(1, 2, 0.9, 0.9)
	assert.Equal(t, "Invalid amount of square meters for the wall.", err.Error())

	_, err = NewWall(1, 2, 16, 1)
	assert.Equal(t, "Invalid amount of square meters for the wall.", err.Error())

	_, err = NewWall(1, 2, 15, 1)
	assert.Nil(t, err)

	_, err = NewWall(1, 0, 2.2, 1)
	assert.Nil(t, err)
}

func TestAreaOfWindowsAndDoors(t *testing.T) {
	_, err := NewWall(20, 50, 15, 1)
	assert.Equal(t, "Doors and windows occupy more than 50% of the wall area.", err.Error())

	_, err = NewWall(1, 2, 15, 1)
	assert.Nil(t, err)
}

func TestPaintFullArea(t *testing.T) {
	wall, _ := NewWall(1, 1, 15, 1)
	assert.Equal(t, 11.08, wall.GetPaintfulArea())
}

func TestValidateSpaceBetweenDoorAndWall(t *testing.T) {
	_, err := NewWall(1, 0, 2, 1)
	assert.Equal(t, "The height of walls with a door must be at least 30 centimeters greater than the height of the door.", err.Error())
	_, err = NewWall(1, 1, 15, 1)
	assert.Nil(t, err)
}
