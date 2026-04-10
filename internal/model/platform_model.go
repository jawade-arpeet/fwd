package model

import (
	"github.com/google/uuid"
)

type PlatformModel struct {
	ID      uuid.UUID `db:"id" json:"id"`
	Name    string    `db:"name" json:"name"`
	LogoURL *string   `db:"logo_url" json:"logo_url"`
}
