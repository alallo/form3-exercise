package httpclient

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockClient struct {
	MockedDo func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.MockedDo != nil {
		return m.MockedDo(req)
	}

	return &http.Response{}, nil
}

const testServerUrl = "http://api.test.com"

func getMockedClientResponse(body string, status int, statusMessage string) (Client, []byte) {
	expectedResponseBody := []byte(body)

	goClient := &MockClient{
		MockedDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: status,
				Status:     statusMessage,
				Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
			}, nil
		},
	}

	client := Client{
		HTTPClient: goClient,
		baseURL:    testServerUrl,
	}

	return client, expectedResponseBody
}

func getMockedClientResponseError() Client {

	goClient := &MockClient{
		MockedDo: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("Do method returned an error")
		},
	}

	client := Client{
		HTTPClient: goClient,
		baseURL:    testServerUrl,
	}

	return client
}

func TestGetResponseOK(t *testing.T) {

	var headers map[string]string
	var queryParams map[string]string

	client, expectedResponseBody := getMockedClientResponse("body", http.StatusOK, "")
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
	expectedStatusMessage := "404 Not Found"
	client, _ := getMockedClientResponse("", http.StatusNotFound, expectedStatusMessage)

	var headers map[string]string
	var queryParams map[string]string

	res, err := client.Get(headers, queryParams)
	if res != nil {
		t.Errorf("request returning a 200 response status, expected an error")
	}

	if err.Error() != expectedStatusMessage {
		t.Errorf("request returning a different error status: got %v expected %v", err.Error(), expectedStatusMessage)
	}
}

func TestGetError(t *testing.T) {

	expectedStatusMessage := "Do method returned an error"
	var headers map[string]string
	var queryParams map[string]string

	client := getMockedClientResponseError()
	_, err := client.Get(headers, queryParams)
	if err.Error() != expectedStatusMessage {
		t.Errorf("request returning a different error status: got %v expected %v", err.Error(), expectedStatusMessage)
	}
}

func TestGetRetryWhenError(t *testing.T) {
	callCount := 0
	const retryCountExpected = 11

	goClient := &MockClient{
		MockedDo: func(req *http.Request) (*http.Response, error) {
			callCount = callCount + 1

			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Status:     "Internal Server Error",
				Body:       ioutil.NopCloser(bytes.NewBufferString("")),
			}, nil
		},
	}

	client := Client{
		HTTPClient: goClient,
		baseURL:    testServerUrl,
	}

	var headers map[string]string
	var queryParams map[string]string

	res, _ := client.Get(headers, queryParams)
	if res != nil {
		t.Errorf("request returning a 200 response status, expected an error")
	}

	if callCount != retryCountExpected {
		t.Errorf("Retrying policy not working as expected. Number of retry %v expected %v", callCount, retryCountExpected)
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
	client, _ := getMockedClientResponse("body", http.StatusOK, "")

	headers := make(map[string]string)
	headers["myheader"] = "cool"
	headers["myheader2"] = "cool2"
	headers["myheader3"] = "cool3"
	var queryParams map[string]string

	_, err := client.Get(headers, queryParams)
	if err != nil {
		t.Errorf("request returning a non 200 response: got %v", err)
	}
}

func TestGetQueryParamsAreAdded(t *testing.T) {
	client, _ := getMockedClientResponse("body", http.StatusOK, "")

	var headers map[string]string
	queryParams := make(map[string]string)

	param1 := "param1"
	value1 := "foo"
	param2 := "param2"
	value2 := "bar"

	queryParams[param1] = value1
	queryParams[param2] = value2

	expectedURL := testServerUrl + "?" + param1 + "=" + value1 + "&" + param2 + "=" + value2
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
	client.baseURL = ""

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

func TestPostOK(t *testing.T) {
	client, expectedResponseBody := getMockedClientResponse("body", http.StatusOK, "")
	var headers map[string]string
	body := []byte("hello world")
	res, err := client.Post(headers, body)
	if err != nil {
		t.Errorf("request returning a non 200 response: got %v", err)
	}
	comparingResult := bytes.Compare(res, expectedResponseBody)
	if comparingResult != 0 {
		t.Errorf("request returning a different response body: got %v expected %v", res, expectedResponseBody)
	}
}

func TestPostResponseError(t *testing.T) {
	expectedStatusMessage := "500 Internal Server Error"
	client, _ := getMockedClientResponse("", http.StatusInternalServerError, expectedStatusMessage)

	var headers map[string]string
	body := []byte("hello world")
	res, err := client.Post(headers, body)
	if res != nil {
		t.Errorf("request returning a 200 response status, expected an error")
	}

	if err.Error() != expectedStatusMessage {
		t.Errorf("request returning a different error status: got %v expected %v", err.Error(), expectedStatusMessage)
	}
}

func TestPostError(t *testing.T) {

	expectedStatusMessage := "Do method returned an error"
	var headers map[string]string
	body := []byte("hello world")

	client := getMockedClientResponseError()
	_, err := client.Post(headers, body)
	if err.Error() != expectedStatusMessage {
		t.Errorf("request returning a different error status: got %v expected %v", err.Error(), expectedStatusMessage)
	}
}

func TestPostFailingToCreateHTTPRequest(t *testing.T) {
	var headers map[string]string

	validURL := "http://myserver.com"

	client, _ := CreateHTTPClient(validURL)

	// force client to have invalid base url
	client.baseURL = ""

	body := []byte("hello world")
	resp, _ := client.Post(headers, body)
	if resp != nil {
		t.Errorf("Client created with an unsupported protocol")
	}
}

func TestPostHeadersAreAdded(t *testing.T) {
	client, _ := getMockedClientResponse("body", http.StatusOK, "")
	headers := make(map[string]string)
	headers["myheader"] = "cool"
	headers["myheader2"] = "cool2"
	headers["myheader3"] = "cool3"

	body := []byte("hello world")
	_, err := client.Post(headers, body)
	if err != nil {
		t.Errorf("request returning a non 200 response: got %v", err)
	}
}

func TestPostRetryWhenError(t *testing.T) {
	callCount := 0
	const retryCountExpected = 11

	goClient := &MockClient{
		MockedDo: func(req *http.Request) (*http.Response, error) {
			callCount = callCount + 1

			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Status:     "Internal Server Error",
				Body:       ioutil.NopCloser(bytes.NewBufferString("")),
			}, nil
		},
	}

	client := Client{
		HTTPClient: goClient,
		baseURL:    testServerUrl,
	}

	var headers map[string]string
	body := []byte("hello world")

	res, _ := client.Post(headers, body)
	if res != nil {
		t.Errorf("request returning a 200 response status, expected an error")
	}

	if callCount != retryCountExpected {
		t.Errorf("Retrying policy not working as expected. Number of retry %v expected %v", callCount, retryCountExpected)
	}
}

func TestDeleteResponseOK(t *testing.T) {

	queryParams := make(map[string]string)
	param1 := "param1"
	value1 := "foo"
	param2 := "param2"
	value2 := "bar"
	queryParams[param1] = value1
	queryParams[param2] = value2

	headers := make(map[string]string)
	headers["myheader"] = "cool"
	headers["myheader2"] = "cool2"
	headers["myheader3"] = "cool3"

	client, _ := getMockedClientResponse("body", http.StatusNoContent, "")
	err := client.Delete(headers, queryParams)
	if err != nil {
		t.Errorf("request not returning a 204 response: got %v", err)
	}
}

func TestDeleteResponseError(t *testing.T) {

	queryParams := make(map[string]string)
	param1 := "param1"
	value1 := "foo"
	param2 := "param2"
	value2 := "bar"
	queryParams[param1] = value1
	queryParams[param2] = value2

	headers := make(map[string]string)
	headers["myheader"] = "cool"
	headers["myheader2"] = "cool2"
	headers["myheader3"] = "cool3"

	expectedStatusMessage := "Specified version incorrect"

	client, _ := getMockedClientResponse("body", http.StatusConflict, expectedStatusMessage)
	err := client.Delete(headers, queryParams)
	if err.Error() != expectedStatusMessage {
		t.Errorf("request returning a different error status: got %v expected %v", err.Error(), expectedStatusMessage)
	}
}

func TestDeleteError(t *testing.T) {

	expectedStatusMessage := "Do method returned an error"
	var headers map[string]string
	var queryParams map[string]string

	client := getMockedClientResponseError()
	err := client.Delete(headers, queryParams)
	if err.Error() != expectedStatusMessage {
		t.Errorf("request returning a different error status: got %v expected %v", err.Error(), expectedStatusMessage)
	}
}

func TestErrorRetryWhenError(t *testing.T) {
	callCount := 0
	const retryCountExpected = 11

	goClient := &MockClient{
		MockedDo: func(req *http.Request) (*http.Response, error) {
			callCount = callCount + 1

			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Status:     "Internal Server Error",
				Body:       ioutil.NopCloser(bytes.NewBufferString("")),
			}, nil
		},
	}

	client := Client{
		HTTPClient: goClient,
		baseURL:    testServerUrl,
	}

	var headers map[string]string
	var queryParams map[string]string

	client.Delete(headers, queryParams)

	if callCount != retryCountExpected {
		t.Errorf("Retrying policy not working as expected. Number of retry %v expected %v", callCount, retryCountExpected)
	}
}
