package account

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccountList(t *testing.T) {
	expectedBody, expectedResponse := getAccountListMockedResponse(t, "testJson/accountlist.json")

	expectedResponseBody := []byte(expectedBody)
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if len(req.URL.Query()) == 7 {
			res.WriteHeader(200)
			res.Write(expectedResponseBody)
		}
	}))
	defer func() { testServer.Close() }()

	var req ListRequest
	req.PageNumber = 1
	req.PageSize = 2
	req.AccountNumber = []string{"123", "456"}
	req.BankID = []string{"3435345", "3435345"}
	req.Iban = []string{"GB29 NWBK 6016 1331 5678 22", "GB29 NWBK 6016 1331 9268 19"}
	req.CustomerID = []string{"CS75847", "CS34834"}
	req.Country = []string{"GB"}
	req.Host = "myapi.form3.com"

	resp, err := GetAccountList(testServer.URL, &req)
	if err != nil {
		t.Errorf("Request is returning an error: got %v", err.Error())
	} else {
		numberOfAccounts := len(resp)
		if numberOfAccounts == 0 {
			t.Errorf("The list of accounts is empty")
		}
		if numberOfAccounts != len(expectedResponse.Accounts) {
			t.Errorf("Number of accounts returned is wrong: got %v expected %v", numberOfAccounts, len(expectedResponse.Accounts))
		}

		firstAccount := resp[0]
		expectedFirstAccount := expectedResponse.Accounts[0]

		if firstAccount.ID != expectedFirstAccount.ID {
			t.Errorf("Response contains wrong ID, got %v expected %v", firstAccount.ID, expectedFirstAccount.ID)
		}
		if firstAccount.Type != expectedFirstAccount.Type {
			t.Errorf("Response contains wrong Type, got %v expected %v", firstAccount.Type, expectedFirstAccount.Type)
		}
		if firstAccount.OrganisationID != expectedFirstAccount.OrganisationID {
			t.Errorf("Response contains wrong OrganisationID, got %v expected %v", firstAccount.OrganisationID, expectedFirstAccount.OrganisationID)
		}
		if firstAccount.Version != expectedFirstAccount.Version {
			t.Errorf("Response contains wrong Version, got %v expected %v", firstAccount.Version, expectedFirstAccount.Version)
		}
		if firstAccount.Attributes.Country != expectedFirstAccount.Attributes.Country {
			t.Errorf("Response contains wrong Country, got %v expected %v", firstAccount.Attributes.Country, expectedFirstAccount.Attributes.Country)
		}
		if firstAccount.Attributes.BaseCurrency != expectedFirstAccount.Attributes.BaseCurrency {
			t.Errorf("Response contains wrong BaseCurrency, got %v expected %v", firstAccount.Attributes.BaseCurrency, expectedFirstAccount.Attributes.BaseCurrency)
		}
		if firstAccount.Attributes.BankID != expectedFirstAccount.Attributes.BankID {
			t.Errorf("Response contains wrong BankID, got %v expected %v", firstAccount.Attributes.BankID, expectedFirstAccount.Attributes.BankID)
		}
		if firstAccount.Attributes.BankIDCode != expectedFirstAccount.Attributes.BankIDCode {
			t.Errorf("Response contains wrong BankIDCode, got %v expected %v", firstAccount.Attributes.BankIDCode, expectedFirstAccount.Attributes.BankIDCode)
		}
		if firstAccount.Attributes.Bic != expectedFirstAccount.Attributes.Bic {
			t.Errorf("Response contains wrong Bic, got %v expected %v", firstAccount.Attributes.Bic, expectedFirstAccount.Attributes.Bic)
		}
		if firstAccount.Attributes.AccountNumber != expectedFirstAccount.Attributes.AccountNumber {
			t.Errorf("Response contains wrong AccountNumber, got %v expected %v", firstAccount.Attributes.AccountNumber, expectedFirstAccount.Attributes.AccountNumber)
		}
	}
}

func TestGetAccountListInvalidUrl(t *testing.T) {
	var req ListRequest
	_, err := GetAccountList("http//foo", &req)
	if err == nil {
		t.Errorf("Request is returning a response with an invalid URL")
	}
}

func TestGetAccountListNotFoundResponse(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(404)
	}))
	defer func() { testServer.Close() }()

	var req ListRequest
	_, err := GetAccountList(testServer.URL, &req)
	if err.Error() != "404 Not Found" {
		t.Errorf("Request is returning an unexpected error: got %v", err.Error())
	}
}

func getAccountListMockedResponse(t *testing.T, fileName string) (string, AccountList) {

	body := readMockedResponseFromFile(t, fileName)

	var response AccountList
	json.Unmarshal([]byte(body), &response)

	return body, response
}
