package account

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetAccount(t *testing.T) {
	expectedBody, expectedResponse := GetAccountMockedResponse(t, "account.json")

	expectedResponseBody := []byte(expectedBody)
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write(expectedResponseBody)
	}))
	defer func() { testServer.Close() }()

	var req AccountFetchRequest
	req.AccountId = "ea6239c1-99e9-42b3-bca1-92f5c068da6b"
	req.Host = "myapi.form3.com"

	resp, err := GetAccount(testServer.URL, &req)
	if err != nil {
		t.Errorf("Request is returning an error: got %v", err.Error())
	}
	isResponseCorrect := reflect.DeepEqual(resp, expectedResponse)

	if !isResponseCorrect {
		t.Errorf("The response received is not the one expected")
	}
}

func GetAccountMockedResponse(t *testing.T, fileName string) (string, Account) {

	body := ReadMockedResponseFromFile(t, fileName)

	var response Account
	json.Unmarshal([]byte(body), &response)

	return body, response
}
