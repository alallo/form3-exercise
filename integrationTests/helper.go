package integration

import (
	"form3-interview/models"
	"testing"
)

// Helper function to compare accounts
func checkAccountResponse(t *testing.T, resp *models.Account, expectedAccount *models.Account) {
	if resp.ID != expectedAccount.ID {
		t.Errorf("Response contains wrong ID, got %v expected %v", resp.ID, expectedAccount.ID)
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
