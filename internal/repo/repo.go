package repo

import "fwd/internal/client"

type Repo struct {
	Account  *AccountRepo
	Platform *PlatformRepo
}

func NewRepo(client *client.Client) *Repo {
	return &Repo{
		Account:  NewAccountRepo(client.Postgres),
		Platform: NewPlatformRepo(client.Postgres),
	}
}
