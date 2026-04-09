package repo

import "fwd/internal/client"

type Repo struct {
	Account *AccountRepo
}

func NewRepo(client *client.Client) *Repo {
	return &Repo{
		Account: NewAccountRepo(client.Postgres),
	}
}
