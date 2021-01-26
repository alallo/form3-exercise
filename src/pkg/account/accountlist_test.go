package account

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetAccountList(t *testing.T) {
	expectedBody, expectedResponse := getAccountListMockedResponse(t, "accountlist.json")

	expectedResponseBody := []byte(expectedBody)
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if len(req.URL.Query()) == 7 {
			res.WriteHeader(200)
			res.Write(expectedResponseBody)
		}
	}))
	defer func() { testServer.Close() }()

	var req AccountListRequest
	req.PageNumber = 1
	req.PageSize = 2
	req.AccountNumber = []string{"123", "456"}
	req.BankID = []string{"3435345", "3435345"}
	req.Iban = []string{"GB29 NWBK 6016 1331 5678 22", "GB29 NWBK 6016 1331 9268 19"}
	req.CustomerID = []string{"CS75847", "CS34834"}
	req.Country = []string{"GB"}
	req.Host = "myapi.form3.com"

	resp, err := GetAccountList(testServer.URL, &req)
	if err != nil {
		t.Errorf("Request is returning an error: got %v", err.Error())
	}
	if len(resp.Accounts) != len(expectedResponse.Accounts) {
		t.Errorf("Number of accounts returned is wrong: got %v expected %v", len(resp.Accounts), len(expectedResponse.Accounts))
	}
	isResponseCorrect := reflect.DeepEqual(resp.Accounts, expectedResponse.Accounts)

	if !isResponseCorrect {
		t.Errorf("The response received is not the one expected")
	}
}

func TestGetAccountListInvalidUrl(t *testing.T) {
	var req AccountListRequest
	_, err := GetAccountList("http//foo", &req)
	if err == nil {
		t.Errorf("Request is returning a response with an invalid URL")
	}
}

func TestGetAccountListNotFoundResponse(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(404)
	}))
	defer func() { testServer.Close() }()

	var req AccountListRequest
	_, err := GetAccountList(testServer.URL, &req)
	if err.Error() != "404 Not Found" {
		t.Errorf("Request is returning an unexpected error: got %v", err.Error())
	}
}

func getAccountListMockedResponse(t *testing.T, fileName string) (string, AccountList) {

	body := readMockedResponseFromFile(t, fileName)

	var response AccountList
	json.Unmarshal([]byte(body), &response)

	return body, response
}
