package utils

type Position struct {
	Start PosXY `json:"start"`
	End   PosXY `json:"end"`
}

type PosXY struct {
	X int `json:"x"`
	Y int `json:"y"`
}
