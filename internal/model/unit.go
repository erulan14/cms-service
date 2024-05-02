package model

type Unit struct {
	UUID   string `json:"uuid,omitempty"`
	Name   string `json:"name"`
	UID    string `json:"uid"`
	Phone  string `json:"phone"`
	Phone2 string `json:"phone2"`
	Tariff string `json:"tariff"`

	Sensors Sensor `json:"sensors"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CreateUnitDTO struct{}
type UpdateUnitDTO struct{}
