package account

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteAccount(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if len(req.URL.Query()) == 1 {
			res.WriteHeader(204)
		}
	}))
	defer func() { testServer.Close() }()

	var req DeleteRequest
	req.AccountID = "c93d6404-8990-4c6b-81f8-7ce67533733d"
	req.Version = 0
	req.Host = "myapi.form3.com"

	err := DeleteAccount(testServer.URL, &req)
	if err != nil {
		t.Errorf("Request is returning an error: got %v", err.Error())
	}
}

func TestDeleteAccountFailed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if len(req.URL.Query()) == 1 {
			res.WriteHeader(404)
		}
	}))
	defer func() { testServer.Close() }()

	var req DeleteRequest
	req.AccountID = "c93d6404-8990-4c6b-81f8-7ce67533733d"
	req.Version = 0
	req.Host = "myapi.form3.com"

	err := DeleteAccount(testServer.URL, &req)
	if err.Error() != "404 Not Found" {
		t.Errorf("Request is returning an unexpected error: got %v", err.Error())
	}
}

func TestDeleteAccountInvalidUrl(t *testing.T) {
	var req DeleteRequest
	err := DeleteAccount("http//foo", &req)
	if err == nil {
		t.Errorf("Request is returning a response with an invalid URL")
	}
}
