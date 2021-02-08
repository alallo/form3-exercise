package account

import (
	"encoding/json"
	"time"

	"form3.com/httpclient"
	"form3.com/models"
)

const accountEndpoint = "/v1/organisation/accounts/"

type AccountResponse struct {
	Account models.Account `json:"data"`
}

type AccountFetchRequest struct {
	AccountId string
	Host      string
}

func GetAccount(url string, request *AccountFetchRequest) (AccountResponse, error) {

	var headers = map[string]string{
		"Host":   request.Host,
		"Date":   time.Now().String(),
		"Accept": "application/vnd.api+json",
	}

	var account AccountResponse

	client, err := httpclient.CreateHTTPClient(url + accountEndpoint + request.AccountId)
	if err != nil {
		return account, err
	}

	resp, err := client.Get(headers, nil)
	if err != nil {
		return account, err
	}

	json.Unmarshal(resp, &account)

	return account, err
}
