package models

import (
	"github.com/google/uuid"
)

type Account struct {
	// The type of resource
	Type string `json:"type"`

	// The unique ID of the resource in UUID 4 format
	ID uuid.UUID `json:"id"`

	// The organisation ID of the organisation by which this resource has been created
	OrganisationID uuid.UUID `json:"organisation_id"`

	// A counter indicating how many times this resource has been modified
	Version int `json:"version"`

	// The specific attributes for each type of resource
	Attributes *AccountAttributes `json:"attributes"`
}
