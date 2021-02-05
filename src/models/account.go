package models

import (
	"github.com/google/uuid"
)

type Account struct {
	Type           string             `json:"type"`
	ID             uuid.UUID          `json:"id"`
	OrganisationID uuid.UUID          `json:"organisation_id"`
	Version        int                `json:"version"`
	Attributes     *AccountAttributes `json:"attributes"`
}
