// Package account provides methods for creating, retrieving or deleteing accounts.
package account

import (
	"strconv"
	"time"

	"form3-interview/httpclient"

	"github.com/google/uuid"
)

// DeleteRequest contains the account ID and version of the account to be deleted
type DeleteRequest struct {
	AccountID uuid.UUID
	Version   int
	Host      string
}

// DeleteAccount call the endpoint to delete an existing account.
// It needs and DeleteRequest containing the Account ID and the version of the account
// It returns an error if the operation fails
// https://api-docs.form3.tech/api.html#organisation-accounts-delete
func DeleteAccount(url string, request *DeleteRequest) error {

	var headers = map[string]string{
		"Host":   request.Host,
		"Date":   time.Now().String(),
		"Accept": "application/vnd.api+json",
	}

	client, err := httpclient.CreateHTTPClient(url + accountEndpoint + request.AccountID.String())
	if err != nil {
		return err
	}

	queryParams := make(map[string]string)
	queryParams["version"] = strconv.Itoa(request.Version)

	err = client.Delete(headers, queryParams)
	if err != nil {
		return err
	}

	return nil
}
