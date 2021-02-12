package integration

import (
	"os"
	"testing"

	"form3-interview/account"
	"form3-interview/models"

	"github.com/google/uuid"
)

func TestCreateAccountReturnsNewAccountSuccessful(t *testing.T) {

	serverURL := os.Getenv("SERVER_URL")

	var newAccountAttrs models.AccountAttributes
	newAccountAttrs.Country = "GB"
	newAccountAttrs.BaseCurrency = "GBP"
	newAccountAttrs.BankID = "401301"
	newAccountAttrs.BankIDCode = "GBSDC"
	newAccountAttrs.Bic = "NWBKGB22"
	newAccountAttrs.AccountNumber = "41426719"
	newAccountAttrs.Iban = "GB11NWBK40030041426820"
	newAccountAttrs.CustomerID = "Ref1234"
	newAccountAttrs.FirstName = "Mario"
	newAccountAttrs.BankAccountName = "Mario Bross"
	newAccountAttrs.AlternativeBankAccountNames = []string{"Alessandro", "Paolo", "Maria"}
	newAccountAttrs.AccountClassification = "Personal"
	newAccountAttrs.JointAccount = true
	newAccountAttrs.Switched = true
	newAccountAttrs.AccountMatchingOptOut = true
	newAccountAttrs.SecondaryIdentification = "A1B2C3D4"

	var newAccount models.Account
	newAccount.ID = uuid.New()

	newAccount.Type = "accounts"
	newAccount.OrganisationID = uuid.New()
	newAccount.Attributes = &newAccountAttrs

	var newData account.Data
	newData.Account = &newAccount

	var req account.CreateRequest
	req.Host = serverURL
	req.Data = &newData

	resp, err := account.CreateAccount(serverURL, &req)
	if err != nil {
		t.Errorf("Request is returning an error: got %v", err.Error())
	} else {
		account.CheckAccountResponse(t, resp, &newAccount)
	}
}

func TestCreateAccountReturnsBadRequest(t *testing.T) {

	serverURL := os.Getenv("SERVER_URL")

	var newAccountAttrs models.AccountAttributes
	newAccountAttrs.Country = "GB"
	newAccountAttrs.BaseCurrency = "GBP"
	newAccountAttrs.BankID = "400300"
	newAccountAttrs.BankIDCode = "GBSDC"

	// Set an invalid Bic value
	newAccountAttrs.Bic = "invalidBic"

	var newAccount models.Account
	newAccount.ID = uuid.New()
	newAccount.Type = "accounts"
	newAccount.OrganisationID = uuid.New()
	newAccount.Attributes = &newAccountAttrs

	var newData account.Data
	newData.Account = &newAccount

	var req account.CreateRequest
	req.Host = serverURL
	req.Data = &newData

	resp, err := account.CreateAccount(serverURL, &req)
	if resp != nil {
		t.Errorf("Received unexpected response")
	}

	if err != nil && err.Error() != "400 Bad Request" {
		t.Errorf("Received unexpected error: %v expected %v", err.Error(), "400 Bad Request")
	}
}

func TestCreateAccountInvalidCoreAttributesReturnsError(t *testing.T) {

	serverURL := os.Getenv("SERVER_URL")

	var newAccountAttrs models.AccountAttributes
	newAccountAttrs.BaseCurrency = "GBP"
	newAccountAttrs.BankID = "400300"
	newAccountAttrs.BankIDCode = "GBSDC"

	var newAccount models.Account
	newAccount.Type = "accounts"
	newAccount.Attributes = &newAccountAttrs

	var newData account.Data

	var req account.CreateRequest
	req.Host = serverURL
	req.Data = &newData

	//test empty ID
	resp, err := account.CreateAccount(serverURL, &req)
	if resp != nil {
		t.Errorf("Invalid ID should return error")
	}

	if err != nil && err.Error() != "500 Internal Server Error" {
		t.Errorf("Received unexpected error: %v expected %v", err.Error(), "500 Internal Server Error")
	}

	//test empty Organisation ID
	newAccount.ID = uuid.New()
	resp, err = account.CreateAccount(serverURL, &req)
	if resp != nil {
		t.Errorf("Invalid Organisation ID should return error")
	}

	if err != nil && err.Error() != "500 Internal Server Error" {
		t.Errorf("Received unexpected error: %v expected %v", err.Error(), "500 Internal Server Error")
	}
}

func TestCreateAccountInvalidAttributesReturnsError(t *testing.T) {

	serverURL := os.Getenv("SERVER_URL")

	var newAccountAttrs models.AccountAttributes
	newAccountAttrs.BaseCurrency = "GBP"
	newAccountAttrs.BankID = "400300"
	newAccountAttrs.BankIDCode = "GBSDC"

	var newAccount models.Account
	newAccount.ID = uuid.New()
	newAccount.Type = "accounts"
	newAccount.OrganisationID = uuid.New()
	newAccount.Attributes = &newAccountAttrs

	var newData account.Data
	newData.Account = &newAccount

	var req account.CreateRequest
	req.Host = serverURL
	req.Data = &newData

	//test empty country
	newAccountAttrs.Country = ""

	resp, err := account.CreateAccount(serverURL, &req)
	if resp != nil {
		t.Errorf("Invalid Country should return error")
	}

	if err != nil && err.Error() != "400 Bad Request" {
		t.Errorf("Received unexpected error: %v expected %v", err.Error(), "400 Bad Request")
	}

	//test invalid Base Currency
	newAccountAttrs.Country = "GB"
	newAccountAttrs.BaseCurrency = "111"

	resp, err = account.CreateAccount(serverURL, &req)
	if resp != nil {
		t.Errorf("Invalid Base Currency should return error")
	}

	if err != nil && err.Error() != "400 Bad Request" {
		t.Errorf("Received unexpected error: %v expected %v", err.Error(), "400 Bad Request")
	}

	//test invalid Bic
	newAccountAttrs.BaseCurrency = "GBP"
	newAccountAttrs.Bic = "invalid"

	resp, err = account.CreateAccount(serverURL, &req)
	if resp != nil {
		t.Errorf("Invalid Bic should return error")
	}

	if err != nil && err.Error() != "400 Bad Request" {
		t.Errorf("Received unexpected error: %v expected %v", err.Error(), "400 Bad Request")
	}

	//test invalid Status
	newAccountAttrs.Bic = ""
	newAccountAttrs.AccountClassification = "myClassification"

	resp, err = account.CreateAccount(serverURL, &req)
	if resp != nil {
		t.Errorf("Invalid Account Classification should return error")
	}

	if err != nil && err.Error() != "400 Bad Request" {
		t.Errorf("Received unexpected error: %v expected %v", err.Error(), "400 Bad Request")
	}
}

func TestCreateAccountAlreadyExistReturnsConflictError(t *testing.T) {

	serverURL := os.Getenv("SERVER_URL")

	var newAccountAttrs models.AccountAttributes
	newAccountAttrs.Country = "GB"
	newAccountAttrs.BaseCurrency = "GBP"
	newAccountAttrs.BankID = "401301"
	newAccountAttrs.BankIDCode = "GBSDC"
	newAccountAttrs.Bic = "NWBKGB22"
	newAccountAttrs.AccountNumber = "41426719"
	newAccountAttrs.Iban = "GB11NWBK40030041426820"
	newAccountAttrs.CustomerID = "Ref1234"
	newAccountAttrs.FirstName = "Mario"
	newAccountAttrs.BankAccountName = "Mario Bross"
	newAccountAttrs.AlternativeBankAccountNames = []string{"Alessandro", "Paolo", "Maria"}
	newAccountAttrs.AccountClassification = "Personal"
	newAccountAttrs.JointAccount = true
	newAccountAttrs.Switched = true
	newAccountAttrs.AccountMatchingOptOut = true
	newAccountAttrs.SecondaryIdentification = "A1B2C3D4"

	var newAccount models.Account
	newAccount.ID, _ = uuid.Parse(validAccountID)
	newAccount.Type = "accounts"
	newAccount.OrganisationID = uuid.New()
	newAccount.Attributes = &newAccountAttrs

	var newData account.Data
	newData.Account = &newAccount

	var req account.CreateRequest
	req.Host = serverURL
	req.Data = &newData

	resp, err := account.CreateAccount(serverURL, &req)
	if resp != nil {
		t.Errorf("Existing Account ID should return error")
	}

	if err != nil && err.Error() != "409 Conflict" {
		t.Errorf("Received unexpected error: %v expected %v", err.Error(), "409 Conflict")
	}
}
