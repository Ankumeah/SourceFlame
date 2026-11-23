package db

import "uuid"

type Repo struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type PR struct {
	Repo   Repo      `json:"repo"`
	UUID   uuid.UUID `json:"uuid"`
	Status string    `json:"status"`
	Owner  string    `json:"owner"`
	Title  string    `json:"title"`
	Body   string    `json:"body"`
	URL    string    `json:"url"`
}
