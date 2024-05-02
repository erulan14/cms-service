package model

type Tariff struct {
	UUID string  `json:"uuid"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
	Curr string  `json:"curr"`
}

type CreateTariffDTO struct {
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
	Curr string  `json:"curr"`
}
type UpdateTariffDTO struct {
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
	Curr string  `json:"curr"`
}
