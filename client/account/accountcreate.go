// Package account provides methods for creating, retrieving or deleteing accounts.
package account

import (
	"encoding/json"
	"strconv"
	"time"

	"form3-interview/httpclient"
	"form3-interview/models"
)

const accountCreateEndpoint = "/v1/organisation/accounts"

// CreateRequest contains the data of the new account and the host
type CreateRequest struct {
	Data *Data `json:"version"`
	Host string
}

// Data wraps the account model in a Data object. Used for json conversion
type Data struct {
	Account *models.Account `json:"data"`
}

// CreateAccount call the endpoint to create a new account.
// It needs a CreateRequest containing the data on the account to create and the base url of the API.
// It returns and Account populated with some extra info after creation.
// https://api-docs.form3.tech/api.html#organisation-accounts-create
func CreateAccount(url string, request *CreateRequest) (*models.Account, error) {

	var data Data

	body, err := json.Marshal(request.Data)
	if err != nil {
		return nil, err
	}

	var headers = map[string]string{
		"Host":           request.Host,
		"Date":           time.Now().String(),
		"Accept":         "application/vnd.api+json",
		"Content-Type":   "application/vnd.api+json",
		"Content-Length": strconv.Itoa(len(body)),
	}

	client, err := httpclient.CreateHTTPClient(url + accountCreateEndpoint)
	if err != nil {
		return nil, err
	}

	resp, err := client.Post(headers, body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(resp, &data)

	account := data.Account

	return account, err
}
