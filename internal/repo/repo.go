package repo

import (
	"ecom/internal/client"
)

type Repo struct {
}

func NewRepo(client *client.Client) *Repo {
	return &Repo{}
}
