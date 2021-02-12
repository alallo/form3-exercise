package integration

import (
	"form3-interview/account"
	"form3-interview/models"
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestFetchAccountReturnsAccountSuccessful(t *testing.T) {

	serverURL := os.Getenv("SERVER_URL")

	// initial set up to create a new account
	var expectedAccountAttrs models.AccountAttributes
	expectedAccountAttrs.Country = "GB"
	expectedAccountAttrs.BaseCurrency = "GBP"
	expectedAccountAttrs.BankID = "400300"
	expectedAccountAttrs.BankIDCode = "GBSDC"
	expectedAccountAttrs.Bic = "NWBKGB22"
	expectedAccountAttrs.AccountNumber = "41426819"
	expectedAccountAttrs.Iban = "GB11NWBK40030041426819"
	expectedAccountAttrs.CustomerID = "Ref123"
	expectedAccountAttrs.FirstName = "Alessandro"
	expectedAccountAttrs.BankAccountName = "Alessandro Lallo"
	expectedAccountAttrs.AlternativeBankAccountNames = []string{"Alessandro", "Paolo", "Maria"}
	expectedAccountAttrs.AccountClassification = "Personal"
	expectedAccountAttrs.JointAccount = true
	expectedAccountAttrs.Switched = true
	expectedAccountAttrs.AccountMatchingOptOut = true
	expectedAccountAttrs.SecondaryIdentification = "A1B2C3D4"

	var expectedAccount models.Account
	expectedAccount.Type = "accounts"
	expectedAccount.ID, _ = uuid.Parse(validAccountID)
	expectedAccount.OrganisationID, _ = uuid.Parse("26069a13-8380-4a06-844d-684ca26f5c2e")
	expectedAccount.Attributes = &expectedAccountAttrs

	var req account.FetchRequest
	req.AccountID, _ = uuid.Parse(validAccountID)
	req.Host = serverURL

	resp, err := account.GetAccount(serverURL, &req)
	if err != nil {
		t.Errorf("Request is returning an error: got %v", err.Error())
	} else {
		account.CheckAccountResponse(t, resp, &expectedAccount)
	}
}

func TestFetchAccountReturnsAccountNotFound(t *testing.T) {

	serverURL := os.Getenv("SERVER_URL")
	var req account.FetchRequest
	req.AccountID = uuid.New()
	req.Host = serverURL

	resp, err := account.GetAccount(serverURL, &req)
	if resp != nil {
		t.Errorf("Request is returning an unexpected response")
	}

	if err != nil && err.Error() != "404 Not Found" {
		t.Errorf("Received unexpected error: %v expected %v", err.Error(), "404 Not Found")
	}
}
