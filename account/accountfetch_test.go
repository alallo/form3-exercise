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
		CheckAccountResponse(t, resp, &expectedResponse.Account)
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
