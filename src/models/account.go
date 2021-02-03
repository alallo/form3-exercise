package models

type Account struct {
	Type           string             `json:"type"`
	ID             string             `json:"id"`
	OrganisationID string             `json:"organisation_id"`
	Version        int                `json:"version"`
	Attributes     *AccountAttributes `json:"attributes"`
}
