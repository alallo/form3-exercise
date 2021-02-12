package integration

import (
	"form3-interview/account"
	"form3-interview/models"
	"os"

	"github.com/google/uuid"
)

const validAccountID = "f7822df5-0549-41b4-9cd9-fad62f3845b2"
const deleteAccountID = "f8a7291f-cb71-4c22-a1fb-203e194a3798"

var listAccountIDs = []string{
	"346acd74-2c18-422c-b32a-e5b1f657658b",
	"45afb6d5-6c85-4320-b982-69cf1e340624",
	"bbd935ab-a4dc-4db5-8cea-1a86e1c25dbb",
	"709d0e6b-0843-4d34-8d11-6d907f959e78",
	"2892067c-cc30-446c-9739-a8c43e417fc5",
	"778ab291-be2d-4645-ae5b-cc78e73f2b42",
	"e6391c11-7b8a-448e-aaba-6f9ffe906146",
	"2b383bfc-c00c-4e94-9747-57adfc69a5d2",
	"1f058487-c79c-4db3-862f-66fed753e463",
	"43c447de-155b-4137-bfd5-0ee684e9189c",
}

func init() {
	serverURL := os.Getenv("SERVER_URL")

	// create account used in the fetch account test and create account conflict test
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
	var newData account.Data
	newData.Account = &expectedAccount

	var req account.CreateRequest
	req.Host = serverURL
	req.Data = &newData

	_, err := account.CreateAccount(serverURL, &req)
	if err != nil && err.Error() != "409 Conflict" {
		panic(err)
	}

	// create account used in the delete account test
	expectedAccount.ID, _ = uuid.Parse(deleteAccountID)

	_, err = account.CreateAccount(serverURL, &req)
	if err != nil && err.Error() != "409 Conflict" {
		panic(err)
	}

	// create 10 accounts to test Get Account List
	i := 0
	for i < 10 {
		expectedAccount.ID, _ = uuid.Parse(listAccountIDs[i])
		_, err = account.CreateAccount(serverURL, &req)
		if err != nil && err.Error() != "409 Conflict" {
			panic(err)
		}
		i = i + 1
	}
}
