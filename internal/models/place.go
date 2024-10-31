package models

type PlacePayload struct {
	Items []Place `json:"items"`
}

type Place struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Vote struct {
	Ids []string `json:"ids"`
}
