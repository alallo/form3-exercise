package account

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func TestGetAccount(t *testing.T) {
	expectedBody, expectedResponse := getAccountMockedResponse(t, "testJson/account.json")

	expectedResponseBody := []byte(expectedBody)
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write(expectedResponseBody)
	}))
	defer func() { testServer.Close() }()

	var req FetchRequest
	req.AccountID, _ = uuid.Parse("ea6239c1-99e9-42b3-bca1-92f5c068da6b")
	req.Host = "myapi.form3.com"

	resp, err := GetAccount(testServer.URL, &req)
	if err != nil {
		t.Errorf("Request is returning an error: got %v", err.Error())
	} else {
		if resp.ID != expectedResponse.Account.ID {
			t.Errorf("Response contains wrong ID, got %v expected %v", resp.ID, expectedResponse.Account.ID)
		}
		if resp.Type != expectedResponse.Account.Type {
			t.Errorf("Response contains wrong Type, got %v expected %v", resp.Type, expectedResponse.Account.Type)
		}
		if resp.OrganisationID != expectedResponse.Account.OrganisationID {
			t.Errorf("Response contains wrong OrganisationID, got %v expected %v", resp.OrganisationID, expectedResponse.Account.OrganisationID)
		}
		if resp.Version != expectedResponse.Account.Version {
			t.Errorf("Response contains wrong Version, got %v expected %v", resp.Version, expectedResponse.Account.Version)
		}
		if resp.Attributes.Country != expectedResponse.Account.Attributes.Country {
			t.Errorf("Response contains wrong Country, got %v expected %v", resp.Attributes.Country, expectedResponse.Account.Attributes.Country)
		}
		if resp.Attributes.BaseCurrency != expectedResponse.Account.Attributes.BaseCurrency {
			t.Errorf("Response contains wrong BaseCurrency, got %v expected %v", resp.Attributes.BaseCurrency, expectedResponse.Account.Attributes.BaseCurrency)
		}
		if resp.Attributes.BankID != expectedResponse.Account.Attributes.BankID {
			t.Errorf("Response contains wrong BankID, got %v expected %v", resp.Attributes.BankID, expectedResponse.Account.Attributes.BankID)
		}
		if resp.Attributes.BankIDCode != expectedResponse.Account.Attributes.BankIDCode {
			t.Errorf("Response contains wrong BankIDCode, got %v expected %v", resp.Attributes.BankIDCode, expectedResponse.Account.Attributes.BankIDCode)
		}
		if resp.Attributes.Bic != expectedResponse.Account.Attributes.Bic {
			t.Errorf("Response contains wrong Bic, got %v expected %v", resp.Attributes.Bic, expectedResponse.Account.Attributes.Bic)
		}
		if resp.Attributes.AccountNumber != expectedResponse.Account.Attributes.AccountNumber {
			t.Errorf("Response contains wrong AccountNumber, got %v expected %v", resp.Attributes.AccountNumber, expectedResponse.Account.Attributes.AccountNumber)
		}
		if resp.Attributes.Iban != expectedResponse.Account.Attributes.Iban {
			t.Errorf("Response contains wrong Iban, got %v expected %v", resp.Attributes.Iban, expectedResponse.Account.Attributes.Iban)
		}
		if resp.Attributes.CustomerID != expectedResponse.Account.Attributes.CustomerID {
			t.Errorf("Response contains wrong CustomerID, got %v expected %v", resp.Attributes.CustomerID, expectedResponse.Account.Attributes.CustomerID)
		}
		if resp.Attributes.FirstName != expectedResponse.Account.Attributes.FirstName {
			t.Errorf("Response contains wrong FirstName, got %v expected %v", resp.Attributes.FirstName, expectedResponse.Account.Attributes.FirstName)
		}
		if resp.Attributes.BankAccountName != expectedResponse.Account.Attributes.BankAccountName {
			t.Errorf("Response contains wrong BankAccountName, got %v expected %v", resp.Attributes.BankAccountName, expectedResponse.Account.Attributes.BankAccountName)
		}
		if resp.Attributes.AlternativeBankAccountNames[0] != expectedResponse.Account.Attributes.AlternativeBankAccountNames[0] {
			t.Errorf("Response contains wrong AlternativeBankAccountNames, got %v expected %v", resp.Attributes.AlternativeBankAccountNames[0], expectedResponse.Account.Attributes.AlternativeBankAccountNames[0])
		}
		if resp.Attributes.AccountClassification != expectedResponse.Account.Attributes.AccountClassification {
			t.Errorf("Response contains wrong AccountClassification, got %v expected %v", resp.Attributes.AccountClassification, expectedResponse.Account.Attributes.AccountClassification)
		}
		if resp.Attributes.JointAccount != expectedResponse.Account.Attributes.JointAccount {
			t.Errorf("Response contains wrong JointAccount, got %v expected %v", resp.Attributes.JointAccount, expectedResponse.Account.Attributes.JointAccount)
		}
		if resp.Attributes.Switched != expectedResponse.Account.Attributes.Switched {
			t.Errorf("Response contains wrong Switched, got %v expected %v", resp.Attributes.Switched, expectedResponse.Account.Attributes.Switched)
		}
		if resp.Attributes.AccountMatchingOptOut != expectedResponse.Account.Attributes.AccountMatchingOptOut {
			t.Errorf("Response contains wrong AccountMatchingOptOut, got %v expected %v", resp.Attributes.AccountMatchingOptOut, expectedResponse.Account.Attributes.AccountMatchingOptOut)
		}
		if resp.Attributes.Status != expectedResponse.Account.Attributes.Status {
			t.Errorf("Response contains wrong Status, got %v expected %v", resp.Attributes.Status, expectedResponse.Account.Attributes.Status)
		}
		if resp.Attributes.SecondaryIdentification != expectedResponse.Account.Attributes.SecondaryIdentification {
			t.Errorf("Response contains wrong SecondaryIdentification, got %v expected %v", resp.Attributes.SecondaryIdentification, expectedResponse.Account.Attributes.SecondaryIdentification)
		}
	}
}

func TestGetAccountNotFound(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(404)
	}))
	defer func() { testServer.Close() }()

	var req FetchRequest
	req.AccountID = uuid.New()
	req.Host = "myapi.form3.com"

	_, err := GetAccount(testServer.URL, &req)
	if err.Error() != "404 Not Found" {
		t.Errorf("Request is returning an unexpected error: got %v", err.Error())
	}
}

func TestGetAccountBadURI(t *testing.T) {
	var req FetchRequest
	req.AccountID = uuid.New()
	req.Host = "myapi.form3.com"

	_, err := GetAccount("foo", &req)
	if err == nil {
		t.Errorf("Request is returning an unexpected error: got %v", err.Error())
	}
}

func getAccountMockedResponse(t *testing.T, fileName string) (string, *AccountResponse) {

	body := readMockedResponseFromFile(t, fileName)

	var response AccountResponse
	json.Unmarshal([]byte(body), &response)

	return body, &response
}
