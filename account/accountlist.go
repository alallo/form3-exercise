// Package account provides methods for creating, retrieving or deleteing accounts.
package account

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"form3-interview/httpclient"
	"form3-interview/models"
)

// AccountList wraps an array of accounts in a Data object. Used for json conversion
type AccountList struct {
	Accounts []models.Account `json:"data"`
}

const defaultPageNumber = 0
const defaultPageSize = 100

const accountListEndpoint = "/v1/organisation/accounts"

// ListRequest contains the filter used to get a list of accounts
type ListRequest struct {
	// Page number being requested
	PageNumber int
	// Size of the page being requested
	PageSize      int
	BankID        []string
	AccountNumber []string
	Iban          []string
	CustomerID    []string
	Country       []string
	Host          string
}

// GetAccountList call the endpoint to fetch a list of accounts.
// It needs a ListRequest containing the filters to apply to the list
// The list can be filtered by BankIDs, Account Numbers, Ibans,CustomerIDs and/or Countries
// It returns a list of Accounts matching the filter specified in the request
// https://api-docs.form3.tech/api.html#organisation-accounts-list
func GetAccountList(url string, request *ListRequest) ([]models.Account, error) {

	var headers = map[string]string{
		"Host":   request.Host,
		"Date":   time.Now().String(),
		"Accept": "application/vnd.api+json",
	}

	queryParams := populateQueryParams(request)

	client, err := httpclient.CreateHTTPClient(url + accountListEndpoint)
	if err != nil {
		return nil, err
	}

	resp, err := client.Get(headers, queryParams)
	if err != nil {
		return nil, err
	}

	var accountlist AccountList
	json.Unmarshal(resp, &accountlist)

	return accountlist.Accounts, err
}

func populateQueryParams(request *ListRequest) map[string]string {
	queryParams := make(map[string]string)

	if request.PageNumber != defaultPageNumber {
		queryParams["page[number]"] = strconv.Itoa(request.PageNumber)
	}

	if request.PageSize != defaultPageSize && request.PageSize != 0 {
		queryParams["page[size]="] = strconv.Itoa(request.PageSize)
	}

	if request.BankID != nil {
		queryParams["filter[bank_id]"] = (strings.Join(request.BankID[:], ","))
	}

	if request.AccountNumber != nil {
		queryParams["filter[account_number]"] = (strings.Join(request.AccountNumber[:], ","))
	}

	if request.Iban != nil {
		queryParams["filter[iban]"] = (strings.Join(request.Iban[:], ","))
	}

	if request.CustomerID != nil {
		queryParams["filter[customer_id]"] = (strings.Join(request.CustomerID[:], ","))
	}

	if request.Country != nil {
		queryParams["filter[country]"] = (strings.Join(request.Country[:], ","))
	}

	return queryParams
}
