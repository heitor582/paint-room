package dtos

type RoomDtoInput struct {
	Walls [4]WallDtoInput `json:"walls"`
}

type RoomDtoOutput struct {
	Cans                 [4]CanDtoOutput `json:"cans"`
	TotalPaitfulRoomArea float64         `json:"totalPaitfulRoomArea"`
	TotalLitersOfCans    float64         `json:"totalLitersOfCans"`
}
