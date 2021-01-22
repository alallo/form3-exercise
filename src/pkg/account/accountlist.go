package accountlist

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"form3.com/httpclient"
	"form3.com/models"
)

type Account = models.Account
type AccountList struct {
	Accounts []Account `json:"data"`
}

const defaultPageNumber = 0
const defaultPageSize = 100

type AccountListRequest struct {
	PageNumber    int
	PageSize      int
	BankID        []string
	AccountNumber []string
	Iban          []string
	CustomerID    []string
	Country       []string
	Host          string
}

func GetAccountList(url string, request *AccountListRequest) (AccountList, error) {

	var headers = map[string]string{
		"Host":   request.Host,
		"Date":   time.Now().String(),
		"Accept": "application/vnd.api+json",
	}

	queryParams := populateQueryParams(request)
	client := httpclient.CreateHTTPClient(url)
	resp, err := client.Get(headers, queryParams)
	var accountlist AccountList
	if err != nil {
		return accountlist, err
	}

	json.Unmarshal(resp, &accountlist)
	return accountlist, err
}

func populateQueryParams(request *AccountListRequest) map[string]string {
	queryParams := make(map[string]string)

	if request.PageNumber > 0 {
		queryParams["page[number]"] = strconv.Itoa(request.PageNumber)
	}

	if request.PageSize != 100 {
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
