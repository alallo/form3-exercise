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
	newAccountAttrs.Iban = "GB11NWBK40030041426819"
	newAccountAttrs.CustomerID = "Ref123"
	newAccountAttrs.FirstName = "Alessandro"
	newAccountAttrs.BankAccountName = "Alessandro Lallo"
	newAccountAttrs.AlternativeBankAccountNames = []string{"Alessandro", "Paolo", "Maria"}
	newAccountAttrs.AccountClassification = "Personal"
	newAccountAttrs.JointAccount = false
	newAccountAttrs.Switched = true
	newAccountAttrs.AccountMatchingOptOut = false
	newAccountAttrs.SecondaryIdentification = "A1B2C3D4"

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
	} else {
		checkAccountResponse(t, resp, expectedResponse.Account)
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
