package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/heitor582/paint_room/src/dtos"
)

func TestRoomService(t *testing.T) {
	var room dtos.RoomDtoInput = dtos.RoomDtoInput{
		Walls: [4]dtos.WallDtoInput{
			{
				DoorsNumber:    0,
				WindownsNumber: 0,
				Height:         15,
				Width:          1,
			},
			{
				DoorsNumber:    0,
				WindownsNumber: 0,
				Height:         15,
				Width:          1,
			},
			{
				DoorsNumber:    0,
				WindownsNumber: 0,
				Height:         15,
				Width:          1,
			},
			{
				DoorsNumber:    0,
				WindownsNumber: 0,
				Height:         15,
				Width:          1,
			},
		},
	}
	var outputMock dtos.RoomDtoOutput = dtos.RoomDtoOutput{
		Cans: [4]dtos.CanDtoOutput{
			{
				Label:    "18 L",
				Value:    18.0,
				Quantity: 0,
			},
			{
				Label:    "3,6 L",
				Value:    3.6,
				Quantity: 3,
			},
			{
				Label:    "2,5 L",
				Value:    2.5,
				Quantity: 0,
			},
			{
				Label:    "0,5 L",
				Value:    0.5,
				Quantity: 3,
			},
		},
		TotalPaitfulRoomArea: 12.0,
		TotalLitersOfCans:    12.3,
	}
	output, err := GetListOfCanOfPaint(room)
	assert.Nil(t, err)
	assert.Equal(t, outputMock, output)
}
