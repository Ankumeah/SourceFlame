package db

import (
	"uuid"
)

type User struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type Repo struct {
	UUID  uuid.UUID `json:"uuid"`
	Name  string    `json:"name"`
	Owner string    `json:"owner"`
}
