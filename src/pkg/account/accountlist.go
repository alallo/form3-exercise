package accountlist

import (
	"encoding/json"
	"strconv"
	"time"

	"form3.com/httpclient"
)

type Filters struct {
	PageNumber    int
	PageSize      int
	BankID        string
	AccountNumber string
}

//AccountList list of accounts
type AccountList struct {
	Accounts []struct {
		Type           string `json:"type"`
		ID             string `json:"id"`
		OrganisationID string `json:"organisation_id"`
		Version        int    `json:"version"`
		Attributes     struct {
			Country               string `json:"country"`
			BaseCurrency          string `json:"base_currency"`
			AccountNumber         string `json:"account_number"`
			BankID                string `json:"bank_id"`
			BankIDCode            string `json:"bank_id_code"`
			Bic                   string `json:"bic"`
			Iban                  string `json:"iban"`
			AccountClassification string `json:"account_classification"`
			JointAccount          bool   `json:"joint_account"`
			Switched              bool   `json:"switched"`
			AccountMatchingOptOut bool   `json:"account_matching_opt_out"`
			Status                string `json:"status"`
		} `json:"attributes"`
	} `json:"data"`
}

func GetAccountList(url string, filters *Filters) (AccountList, error) {

	var headers = map[string]string{
		"Host":   "api.form3.tech",
		"Date":   time.Now().String(),
		"Accept": "application/vnd.api+json",
	}

	queryParams := populateQueryParams(filters)

	resp, err := httpclient.Get(url, headers, queryParams)
	var accountlist AccountList
	if err != nil {
		return accountlist, err
	}

	json.Unmarshal(resp, &accountlist)
	return accountlist, err
}

func populateQueryParams(filters *Filters) map[string]string {
	var queryParams = map[string]string{
		"page[number]":           strconv.Itoa(filters.PageNumber),
		"page[size]=":            strconv.Itoa(filters.PageSize),
		"filter[bank_id]":        filters.BankID,
		"filter[account_number]": filters.AccountNumber,
	}
	return queryParams
}
