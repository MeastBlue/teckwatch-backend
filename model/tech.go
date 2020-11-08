package model

type Tech struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type NewTech struct {
	Label       string `json:"label"`
	Description string `json:"description"`
}
