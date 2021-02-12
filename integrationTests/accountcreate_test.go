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
		checkAccountResponse(t, resp, &newAccount)
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

// Helper function
func checkAccountResponse(t *testing.T, resp *models.Account, expectedAccount *models.Account) {
	if resp.ID != expectedAccount.ID {
		t.Errorf("Response contains wrong ID, got %v expected %v", resp.ID, validAccountID)
	}
	if resp.Type != expectedAccount.Type {
		t.Errorf("Response contains wrong Type, got %v expected %v", resp.Type, expectedAccount.Type)
	}
	if resp.OrganisationID != expectedAccount.OrganisationID {
		t.Errorf("Response contains wrong OrganisationID, got %v expected %v", resp.OrganisationID, expectedAccount.OrganisationID)
	}
	if resp.Version != expectedAccount.Version {
		t.Errorf("Response contains wrong Version, got %v expected %v", resp.Version, expectedAccount.Version)
	}
	if resp.Attributes.Country != expectedAccount.Attributes.Country {
		t.Errorf("Response contains wrong Country, got %v expected %v", resp.Attributes.Country, expectedAccount.Attributes.Country)
	}
	if resp.Attributes.BaseCurrency != expectedAccount.Attributes.BaseCurrency {
		t.Errorf("Response contains wrong BaseCurrency, got %v expected %v", resp.Attributes.BaseCurrency, expectedAccount.Attributes.BaseCurrency)
	}
	if resp.Attributes.BankID != expectedAccount.Attributes.BankID {
		t.Errorf("Response contains wrong BankID, got %v expected %v", resp.Attributes.BankID, expectedAccount.Attributes.BankID)
	}
	if resp.Attributes.BankIDCode != expectedAccount.Attributes.BankIDCode {
		t.Errorf("Response contains wrong BankIDCode, got %v expected %v", resp.Attributes.BankIDCode, expectedAccount.Attributes.BankIDCode)
	}
	if resp.Attributes.Bic != expectedAccount.Attributes.Bic {
		t.Errorf("Response contains wrong Bic, got %v expected %v", resp.Attributes.Bic, expectedAccount.Attributes.Bic)
	}
	if resp.Attributes.AccountNumber != expectedAccount.Attributes.AccountNumber {
		t.Errorf("Response contains wrong AccountNumber, got %v expected %v", resp.Attributes.AccountNumber, expectedAccount.Attributes.AccountNumber)
	}
	if resp.Attributes.CustomerID != expectedAccount.Attributes.CustomerID {
		t.Errorf("Response contains wrong CustomerID, got %v expected %v", resp.Attributes.CustomerID, expectedAccount.Attributes.CustomerID)
	}
	if resp.Attributes.FirstName != expectedAccount.Attributes.FirstName {
		t.Errorf("Response contains wrong FirstName, got %v expected %v", resp.Attributes.FirstName, expectedAccount.Attributes.FirstName)
	}
	if resp.Attributes.BankAccountName != expectedAccount.Attributes.BankAccountName {
		t.Errorf("Response contains wrong BankAccountName, got %v expected %v", resp.Attributes.BankAccountName, expectedAccount.Attributes.BankAccountName)
	}
	if resp.Attributes.AlternativeBankAccountNames[0] != expectedAccount.Attributes.AlternativeBankAccountNames[0] {
		t.Errorf("Response contains wrong AlternativeBankAccountNames, got %v expected %v", resp.Attributes.AlternativeBankAccountNames[0], expectedAccount.Attributes.AlternativeBankAccountNames[0])
	}
	if resp.Attributes.AccountClassification != expectedAccount.Attributes.AccountClassification {
		t.Errorf("Response contains wrong AccountClassification, got %v expected %v", resp.Attributes.AccountClassification, expectedAccount.Attributes.AccountClassification)
	}
	if resp.Attributes.JointAccount != expectedAccount.Attributes.JointAccount {
		t.Errorf("Response contains wrong JointAccount, got %v expected %v", resp.Attributes.JointAccount, expectedAccount.Attributes.JointAccount)
	}
	if resp.Attributes.Switched != expectedAccount.Attributes.Switched {
		t.Errorf("Response contains wrong Switched, got %v expected %v", resp.Attributes.Switched, expectedAccount.Attributes.Switched)
	}
	if resp.Attributes.AccountMatchingOptOut != expectedAccount.Attributes.AccountMatchingOptOut {
		t.Errorf("Response contains wrong AccountMatchingOptOut, got %v expected %v", resp.Attributes.AccountMatchingOptOut, expectedAccount.Attributes.AccountMatchingOptOut)
	}
	if resp.Attributes.Status != expectedAccount.Attributes.Status {
		t.Errorf("Response contains wrong Status, got %v expected %v", resp.Attributes.Status, expectedAccount.Attributes.Status)
	}
	if resp.Attributes.SecondaryIdentification != expectedAccount.Attributes.SecondaryIdentification {
		t.Errorf("Response contains wrong SecondaryIdentification, got %v expected %v", resp.Attributes.SecondaryIdentification, expectedAccount.Attributes.SecondaryIdentification)
	}
}
