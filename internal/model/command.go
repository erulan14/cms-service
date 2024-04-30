package model

type Command struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Command string `json:"command"`
	Unit    string `json:"unit"`
	Creator string `json:"creator"`
}

type CreateCommandDTO struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Unit    string `json:"unit"`
	Creator string `json:"creator"`
}

type UpdateCommandDTO struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Unit    string `json:"unit"`
	Creator string `json:"creator"`
}
