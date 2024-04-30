package model

type Sensor struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Unit  string `json:"unit"`
	Param string `json:"param"`
	Table Table  `json:"table"`
	Cal   string `json:"cal"`
}

type Table struct {
	XY    [][]int `json:"xy"`
	Bound Bound   `json:"bound"`
}

type Bound struct {
	Lower int `json:"lower"`
	Upper int `json:"upper"`
}
