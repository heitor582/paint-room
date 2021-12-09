package services

import (
	"math"

	"gitlab.com/heitor582/paint_room/src/dtos"
	"gitlab.com/heitor582/paint_room/src/entities"
)

func GetListOfCanOfPaint(roomDto dtos.RoomDtoInput) (dtos.RoomDtoOutput, error) {
	room := entities.NewRoom()
	for i, roomIndex := range roomDto.Walls {
		wall, err := entities.NewWall(roomIndex.DoorsNumber, roomIndex.WindownsNumber, roomIndex.Height, roomIndex.Width)
		if err != nil {
			return dtos.RoomDtoOutput{}, err
		}
		room.Walls[i] = *wall
	}

	total, err := room.CalculateTotalOfLitersArea()
	if err != nil {
		return dtos.RoomDtoOutput{}, err
	}

	var output dtos.RoomDtoOutput = dtos.RoomDtoOutput{
		Cans: [4]dtos.CanDtoOutput{
			{
				Label:    "18 L",
				Value:    18.0,
				Quantity: 0,
			},
			{
				Label:    "3,6 L",
				Value:    3.6,
				Quantity: 0,
			},
			{
				Label:    "2,5 L",
				Value:    2.5,
				Quantity: 0,
			},
			{
				Label:    "0,5 L",
				Value:    0.5,
				Quantity: 0,
			},
		},
		TotalPaitfulRoomArea: total,
		TotalLitersOfCans:    0.0,
	}

	for i, v := range output.Cans {
		if total > v.Value {
			intValue, _ := math.Modf(total / v.Value)
			output.Cans[i].Quantity = int(intValue)
			total -= (v.Value * intValue)
			output.TotalLitersOfCans += (v.Value * intValue)
		}
		if total < 0.5 {
			output.Cans[i].Quantity += 1
			output.TotalLitersOfCans += 0.5
		}
	}

	return output, nil
}
