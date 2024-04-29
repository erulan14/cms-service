package model

type Device struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Port    int    `json:"port"`
}

type CreateUserDTO struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Port    int    `json:"port"`
}

type UpdateDeviceDTO struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Port    int    `json:"port"`
}
