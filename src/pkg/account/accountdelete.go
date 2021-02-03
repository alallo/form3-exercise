package account

import (
	"strconv"
	"time"

	"form3.com/httpclient"
)

type AccountDeleteRequest struct {
	AccountID string
	Version   int
	Host      string
}

func DeleteAccount(url string, request *AccountDeleteRequest) error {

	var headers = map[string]string{
		"Host":   request.Host,
		"Date":   time.Now().String(),
		"Accept": "application/vnd.api+json",
	}

	client, err := httpclient.CreateHTTPClient(url + accountEndpoint + request.AccountID)
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
