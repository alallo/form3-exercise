package account

import (
	"encoding/json"
	"time"

	"form3.com/httpclient"
)

const accountEndpoint = "/v1/organisation/accounts/"

type AccountFetchRequest struct {
	AccountId string
	Host      string
}

func GetAccount(url string, request *AccountFetchRequest) (Account, error) {

	var headers = map[string]string{
		"Host":   request.Host,
		"Date":   time.Now().String(),
		"Accept": "application/vnd.api+json",
	}

	var account Account

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
