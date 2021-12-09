package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuantityOfWalls(t *testing.T) {
	room := NewRoom()
	for i := range room.Walls {
		room.Walls[i] = Wall{
			DoorsNumber:    0,
			WindownsNumber: 0,
			Height:         15.0,
			Width:          1.0,
		}
	}

	err := room.IsValid()
	assert.Nil(t, err)

	room.Walls = [4]Wall{}
	err = room.IsValid()
	assert.Equal(t, "Invalid quantity of walls minumiun is 4.", err.Error())
}

func TestCalculateTotalOfLitersArea(t *testing.T) {
	room := NewRoom()
	for i := range room.Walls {
		room.Walls[i] = Wall{
			DoorsNumber:    0,
			WindownsNumber: 0,
			Height:         15.0,
			Width:          1.0,
		}
	}
	valor, err := room.CalculateTotalOfLitersArea()
	assert.Nil(t, err)
	assert.Equal(t, 12.0, valor)
}
