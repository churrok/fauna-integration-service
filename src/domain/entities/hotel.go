package entities

type Hotel struct {
	Id      *string `json:"id" fauna:"id"`
	Name    string  `json:"name" fauna:"nombre"`
	Country string  `json:"country" fauna:"pais"`
}
