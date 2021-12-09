package dtos

type WallDtoInput struct {
	DoorsNumber    int     `json:"doorsNumber"`
	WindownsNumber int     `json:"windownsNumber"`
	Height         float64 `json:"height"`
	Width          float64 `json:"width"`
}
