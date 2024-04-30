package model

import "time"

type Geozone struct {
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Positions   []Pos     `json:"positions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Type        string    `json:"type"`
	Color       string    `json:"color"`
	Radius      float32   `json:"radius"`
	Width       float32   `json:"width"`
	Description string    `json:"description"`
	Area        float32   `json:"area"`
	Perimeter   float32   `json:"perimeter"`
}

type Pos struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type CreateGeozoneDTO struct {
	Name        string    `json:"name"`
	Positions   []Pos     `json:"positions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Type        string    `json:"type"`
	Color       string    `json:"color"`
	Radius      float32   `json:"radius"`
	Width       float32   `json:"width"`
	Description string    `json:"description"`
	Area        float32   `json:"area"`
	Perimeter   float32   `json:"perimeter"`
}

type UpdateGeozoneDTO struct {
	Name        string    `json:"name"`
	Positions   []Pos     `json:"positions"`
	UpdatedAt   time.Time `json:"updated_at"`
	Type        string    `json:"type"`
	Color       string    `json:"color"`
	Radius      float32   `json:"radius"`
	Width       float32   `json:"width"`
	Description string    `json:"description"`
	Area        float32   `json:"area"`
	Perimeter   float32   `json:"perimeter"`
}
