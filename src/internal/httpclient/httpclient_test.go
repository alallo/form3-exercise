package httpclient

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetResponseOK(t *testing.T) {
	expectedResponseBody := []byte("body")
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write(expectedResponseBody)
	}))
	defer func() { testServer.Close() }()
	var headers map[string]string
	var queryParams map[string]string
	client, _ := CreateHTTPClient(testServer.URL)
	res, err := client.Get(headers, queryParams)
	if err != nil {
		t.Errorf("request returning a non 200 response: got %v", err)
	}
	comparingResult := bytes.Compare(res, expectedResponseBody)
	if comparingResult != 0 {
		t.Errorf("request returning a different response body: got %v expected %v", res, expectedResponseBody)
	}
}

func TestGetResponseError(t *testing.T) {
	expectedResponseBody := []byte("body")
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(404)
		res.Write(expectedResponseBody)
	}))
	defer func() { testServer.Close() }()
	var headers map[string]string
	var queryParams map[string]string
	client, _ := CreateHTTPClient(testServer.URL)
	res, err := client.Get(headers, queryParams)
	if res != nil {
		t.Errorf("request returning a 200 response status, expected an error")
	}
	expectedStatusMessage := "404 Not Found"
	if err.Error() != expectedStatusMessage {
		t.Errorf("request returning a different error status: got %v expected %v", err.Error(), "404 Not Found")
	}
}

func TestCreateHTTPClientWithValidURL(t *testing.T) {
	validURL := "http://myfakeserver.com:8080"
	client, err := CreateHTTPClient(validURL)
	if err != nil {
		t.Errorf("Failed to create the httpclient: got %v", err.Error())
	}
	if client.baseURL != validURL {
		t.Errorf("Client created with wrong base URL: got %v expected %v", client.baseURL, validURL)
	}
}

func TestCreateHTTPClientWithInvalidURL(t *testing.T) {
	invalidURL := "http//myfakeserver.com:8080"
	client, _ := CreateHTTPClient(invalidURL)
	if client != nil {
		t.Errorf("Client created with invalid base URL")
	}
}

func TestGetHeadersAreAdded(t *testing.T) {
	expectedResponseBody := []byte("body")
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write(expectedResponseBody)
	}))
	defer func() { testServer.Close() }()
	headers := make(map[string]string)
	headers["myheader"] = "cool"
	headers["myheader2"] = "cool2"
	headers["myheader3"] = "cool3"
	var queryParams map[string]string
	client, _ := CreateHTTPClient(testServer.URL)
	_, err := client.Get(headers, queryParams)
	if err != nil {
		t.Errorf("request returning a non 200 response: got %v", err)
	}
}

func TestGetQueryParamsAreAdded(t *testing.T) {
	expectedResponseBody := []byte("body")
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write(expectedResponseBody)
	}))

	testServerBaseURL := testServer.URL

	defer func() { testServer.Close() }()
	var headers map[string]string
	queryParams := make(map[string]string)

	param1 := "param1"
	value1 := "foo"
	param2 := "param2"
	value2 := "bar"

	queryParams[param1] = value1
	queryParams[param2] = value2
	client, _ := CreateHTTPClient(testServer.URL)

	expectedURL := testServerBaseURL + "?" + param1 + "=" + value1 + "&" + param2 + "=" + value2
	_, err := client.Get(headers, queryParams)
	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}
	if client.baseURL != expectedURL {
		t.Errorf("Query parameters not added correctly to the URL: got %v expected %v", client.baseURL, expectedURL)
	}
}

func TestGetFailingToCreateHTTPRequest(t *testing.T) {
	var headers map[string]string
	var queryParams map[string]string

	validURL := "http://myserver.com"

	client, _ := CreateHTTPClient(validURL)

	// force client to have invalid base url
	client.baseURL = "http//myserver"

	resp, _ := client.Get(headers, queryParams)
	if resp != nil {
		t.Errorf("Client created with an unsupported protocol")
	}
}

func TestGetRedirect(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(redirectHandler))
	defer func() { testServer.Close() }()

	var headers map[string]string
	var queryParams map[string]string
	client, _ := CreateHTTPClient(testServer.URL)
	_, err := client.Get(headers, queryParams)
	if err != nil {
		t.Errorf("request returning a non 200 response: got %v", err)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://blablabla.com", http.StatusFound)
}
