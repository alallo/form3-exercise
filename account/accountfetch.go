// Package account provides methods for creating, retrieving or deleteing accounts.
package account

import (
	"encoding/json"
	"errors"
	"time"

	"form3-interview/httpclient"
	"form3-interview/models"

	"github.com/google/uuid"
)

const accountEndpoint = "/v1/organisation/accounts/"

// AccountResponse wraps the account model in a Data object.  Used for json conversion
type AccountResponse struct {
	Account models.Account `json:"data"`
}

// FetchRequest contains the account ID to find and the host
type FetchRequest struct {
	AccountID uuid.UUID
	Host      string
}

// GetAccount call the endpoint to fetch a single account.
// It needs a FetchRequest containing the account ID that needs to be fetched.
// Account ID is mandatory to avoid
// It returns an Account if the account ID matches a record in the database.
// https://api-docs.form3.tech/api.html#organisation-accounts-fetch
func GetAccount(url string, request *FetchRequest) (*models.Account, error) {

	// this check is needed to avoid making this a call to get a list of accounts
	if request.AccountID.String() == "" {
		return nil, errors.New("AccountID is mandatory to fetch an account")
	}

	var headers = map[string]string{
		"Host":   request.Host,
		"Date":   time.Now().String(),
		"Accept": "application/vnd.api+json",
	}

	var accountResponse AccountResponse

	client, err := httpclient.CreateHTTPClient(url + accountEndpoint + request.AccountID.String())
	if err != nil {
		return nil, err
	}

	resp, err := client.Get(headers, nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp, &accountResponse)

	return &accountResponse.Account, err
}
