package integration

import (
	"form3-interview/account"
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestDeleteAccountReturnsSuccess(t *testing.T) {

	serverURL := os.Getenv("SERVER_URL")

	var req account.DeleteRequest
	req.AccountID, _ = uuid.Parse(deleteAccountID)
	req.Version = 0
	req.Host = serverURL

	err := account.DeleteAccount(serverURL, &req)

	if err != nil {
		t.Errorf("Received unexpected error: %v", err.Error())
	}
}

func TestDeleteAccountReturnsAccountNotFound(t *testing.T) {

	serverURL := os.Getenv("SERVER_URL")

	var req account.DeleteRequest
	req.AccountID = uuid.New()
	req.Version = 0
	req.Host = serverURL

	err := account.DeleteAccount(serverURL, &req)

	if err != nil && err.Error() != "404 Not Found" {
		t.Errorf("Request is returning an unexpected error: got %v expected %v", err.Error(), "404 Not Found")
	}
}

func TestDeleteAccountReturnsConflict(t *testing.T) {

	serverURL := os.Getenv("SERVER_URL")

	var req account.DeleteRequest
	req.AccountID, _ = uuid.Parse("346acd74-2c18-422c-b32a-e5b1f657658b")
	req.Version = 1
	req.Host = serverURL

	err := account.DeleteAccount(serverURL, &req)

	if err != nil && err.Error() != "409 Conflict" {
		t.Errorf("Request is returning an unexpected error: got %v expected %v", err.Error(), "409 Conflict")
	}
}
