package account

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"form3-interview/models"

	"github.com/google/uuid"
)

func TestCreateAccountOK(t *testing.T) {

	expectedBody, expectedResponse := getCreateAccountMockedResponse(t, "testJson/newaccount.json")
	expectedResponseBody := []byte(expectedBody)
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write(expectedResponseBody)
	}))
	defer func() { testServer.Close() }()

	var newAccountAttrs models.AccountAttributes
	newAccountAttrs.Country = "GB"
	newAccountAttrs.BaseCurrency = "GBP"
	newAccountAttrs.BankID = "400300"
	newAccountAttrs.BankIDCode = "GBSDC"
	newAccountAttrs.Bic = "NWBKGB22"
	newAccountAttrs.AccountNumber = "41426819"
	newAccountAttrs.Name = [4]string{"Samantha Holder"}

	var newAccount models.Account
	newAccount.ID = uuid.New()
	newAccount.Type = "accounts"
	newAccount.OrganisationID = uuid.New()
	newAccount.Attributes = &newAccountAttrs

	var newData Data
	newData.Account = &newAccount

	var req CreateRequest
	req.Host = "api.form3.tech"
	req.Data = &newData

	resp, err := CreateAccount(testServer.URL, &req)
	if err != nil {
		t.Errorf("Request is returning an error: got %v", err.Error())
	}
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
	if resp.Attributes.Name[0] != expectedResponse.Account.Attributes.Name[0] {
		t.Errorf("Response contains wrong Name, got %v expected %v", resp.Attributes.Name[0], expectedResponse.Account.Attributes.Name[0])
	}
	if resp.Attributes.Status != "pending" {
		t.Errorf("The status expected is not right")
	}
}

func TestCreateAccountFailed(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(409)
	}))
	defer func() { testServer.Close() }()

	var req CreateRequest
	_, err := CreateAccount(testServer.URL, &req)

	if err.Error() != "409 Conflict" {
		t.Errorf("Response contains wrong error, got %v", err.Error())
	}
}

func TestCreateAccountBadURI(t *testing.T) {
	var req CreateRequest

	_, err := CreateAccount("foo", &req)
	if err == nil {
		t.Errorf("Request is returning an unexpected error: got %v", err.Error())
	}
}

func getCreateAccountMockedResponse(t *testing.T, fileName string) (string, Data) {

	body := readMockedResponseFromFile(t, fileName)

	var response Data
	json.Unmarshal([]byte(body), &response)

	return body, response
}
