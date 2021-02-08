package account

import (
	"encoding/json"
	"strconv"
	"time"

	"form3-interview/httpclient"
	"form3-interview/models"
)

const accountCreateEndpoint = "/v1/organisation/accounts"

type AccountCreateRequest struct {
	Data *Data `json:"version"`
	Host string
}

type Data struct {
	Account *models.Account `json:"data"`
}

func CreateAccount(url string, request *AccountCreateRequest) (models.Account, error) {

	var data Data
	var account models.Account

	body, err := json.Marshal(request.Data)
	if err != nil {
		return account, err
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
		return account, err
	}

	resp, err := client.Post(headers, body)
	if err != nil {
		return account, err
	}

	json.Unmarshal(resp, &data)

	account = *data.Account

	return account, err
}
