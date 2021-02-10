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

	var newData account.Data
	newData.Account = &newAccount

	var req account.CreateRequest
	req.Host = serverURL
	req.Data = &newData

	resp, err := account.CreateAccount(serverURL, &req)
	if err != nil {
		t.Errorf("Request is returning an error: got %v", err.Error())
	}
	if resp.ID != newData.Account.ID {
		t.Errorf("Response contains wrong ID, got %v expected %v", resp.ID, newData.Account.ID)
	}
	if resp.Type != newData.Account.Type {
		t.Errorf("Response contains wrong Type, got %v expected %v", resp.Type, newData.Account.Type)
	}
	if resp.OrganisationID != newData.Account.OrganisationID {
		t.Errorf("Response contains wrong OrganisationID, got %v expected %v", resp.OrganisationID, newData.Account.OrganisationID)
	}
	if resp.Version != newData.Account.Version {
		t.Errorf("Response contains wrong Version, got %v expected %v", resp.Version, newData.Account.Version)
	}
	if resp.Attributes.Country != newData.Account.Attributes.Country {
		t.Errorf("Response contains wrong Country, got %v expected %v", resp.Attributes.Country, newData.Account.Attributes.Country)
	}
	if resp.Attributes.BaseCurrency != newData.Account.Attributes.BaseCurrency {
		t.Errorf("Response contains wrong BaseCurrency, got %v expected %v", resp.Attributes.BaseCurrency, newData.Account.Attributes.BaseCurrency)
	}
	if resp.Attributes.BankID != newData.Account.Attributes.BankID {
		t.Errorf("Response contains wrong BankID, got %v expected %v", resp.Attributes.BankID, newData.Account.Attributes.BankID)
	}
	if resp.Attributes.BankIDCode != newData.Account.Attributes.BankIDCode {
		t.Errorf("Response contains wrong BankIDCode, got %v expected %v", resp.Attributes.BankIDCode, newData.Account.Attributes.BankIDCode)
	}
	if resp.Attributes.Bic != newData.Account.Attributes.Bic {
		t.Errorf("Response contains wrong Bic, got %v expected %v", resp.Attributes.Bic, newData.Account.Attributes.Bic)
	}
	if resp.Attributes.AccountNumber != newData.Account.Attributes.AccountNumber {
		t.Errorf("Response contains wrong AccountNumber, got %v expected %v", resp.Attributes.AccountNumber, newData.Account.Attributes.AccountNumber)
	}
	if resp.Attributes.Iban != newData.Account.Attributes.Iban {
		t.Errorf("Response contains wrong Iban, got %v expected %v", resp.Attributes.Iban, newData.Account.Attributes.Iban)
	}
	if resp.Attributes.CustomerID != newData.Account.Attributes.CustomerID {
		t.Errorf("Response contains wrong CustomerID, got %v expected %v", resp.Attributes.CustomerID, newData.Account.Attributes.CustomerID)
	}
	if resp.Attributes.FirstName != newData.Account.Attributes.FirstName {
		t.Errorf("Response contains wrong FirstName, got %v expected %v", resp.Attributes.FirstName, newData.Account.Attributes.FirstName)
	}
	if resp.Attributes.BankAccountName != newData.Account.Attributes.BankAccountName {
		t.Errorf("Response contains wrong BankAccountName, got %v expected %v", resp.Attributes.BankAccountName, newData.Account.Attributes.BankAccountName)
	}
	if resp.Attributes.AlternativeBankAccountNames[0] != newData.Account.Attributes.AlternativeBankAccountNames[0] {
		t.Errorf("Response contains wrong AlternativeBankAccountNames, got %v expected %v", resp.Attributes.AlternativeBankAccountNames[0], newData.Account.Attributes.AlternativeBankAccountNames[0])
	}
	if resp.Attributes.AccountClassification != newData.Account.Attributes.AccountClassification {
		t.Errorf("Response contains wrong AccountClassification, got %v expected %v", resp.Attributes.AccountClassification, newData.Account.Attributes.AccountClassification)
	}
	if resp.Attributes.JointAccount != newData.Account.Attributes.JointAccount {
		t.Errorf("Response contains wrong JointAccount, got %v expected %v", resp.Attributes.JointAccount, newData.Account.Attributes.JointAccount)
	}
	if resp.Attributes.Switched != newData.Account.Attributes.Switched {
		t.Errorf("Response contains wrong Switched, got %v expected %v", resp.Attributes.Switched, newData.Account.Attributes.Switched)
	}
	if resp.Attributes.AccountMatchingOptOut != newData.Account.Attributes.AccountMatchingOptOut {
		t.Errorf("Response contains wrong AccountMatchingOptOut, got %v expected %v", resp.Attributes.AccountMatchingOptOut, newData.Account.Attributes.AccountMatchingOptOut)
	}
	if newData.Account.Attributes.Status != "confirmed" {
		t.Errorf("Response contains wrong Status, got %v expected %v", newData.Account.Attributes.Status, "confirmed")
	}
	if resp.Attributes.SecondaryIdentification != newData.Account.Attributes.SecondaryIdentification {
		t.Errorf("Response contains wrong SecondaryIdentification, got %v expected %v", resp.Attributes.SecondaryIdentification, newData.Account.Attributes.SecondaryIdentification)
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

	if err.Error() != "400 Bad Request" {
		t.Errorf("Received unexpected error: %v", err.Error())
	}
}
