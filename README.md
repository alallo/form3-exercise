## Description 
Take home exercise done as part of the form3 hiring process. 

The project is formed by 3 main bits:
* the client, to access some of the account functionalities available through the light version of the API available
* httpclient, a wrapper for the http functionalities
* cmd, a command line useful to manually try the client and interact with the API
* integration tests, they are part of the dockerfile and run at startup 

Few of the integration tests are failing to highlight some of the problem found with the API. Please refer to [Issues](#issues)

For the purpose of this exercise the set up of data needed for the integration tests are done in code through the init function of the integration package. In an ideal world the data would be part of a sql script that would run with the table creation (or a sandbox provided by the third-party API)
## Project structure

Folder Name | Description
------------ | -------------
account | a Go client to inteface with form3 APIs. This implements some of the "Account" functionalities
cmd | Command line app. Useful to play with the client
httpclient | a wrapper to help handling an http client
models | this contains the account and acccountattributes models that are shared and used in different files
integrationTests | contains the integration tests that will be run through the docker-compose file
scripts | the origninal sql script provided by form3 to create the DB

## Prerequisite
To run the command line tool and to test the client you will need

* Go https://golang.org/doc/install
* Docker https://www.docker.com/get-started

## Try it

First of all use the docker compose file from the scripts folder to spin up the form3 API

```
docker-compose up
```

This will run the integration tests as part of the process.

Now you can run the command line tool and interact with the client running

```
SERVER_URL={your_server_url} HOST={your_host} go run main.go
```

because we are using the docker-compose file and not the real API endpoints, we need to point the client to our local instance of the API

```
cd cmd
SERVER_URL=http://localhost:8080 HOST=http://localhost:8080 go run main.go
```

## Example
This example is provided assuming the account package is hosted on a public repo called "form3-interview"

```Go
package main

import (
	"form3-interview/account"
	"form3-interview/models"
)


func main() {
  serverURL := os.Getenv("SERVER_URL")
  host := os.Getenv("HOST")
  
  var req account.ListRequest
  req.PageNumber = pageNumberInt
  req.PageSize = pageSizeInt
  req.Host = host
  
  resp, err := account.GetAccountList(serverURL, &req)
  if err != nil {
    fmt.Println("Error: ", err)
  } else {
    body, err := json.MarshalIndent(resp, "", "  ")
    if err != nil {
      fmt.Println("Error: ", err)
    }
    fmt.Println("List of Accounts:")
    fmt.Println(string(body))
  }
}
```

## Testing
The testing strategy for this project is to use 3 different types of testing:
* unit tests to test the httpclient package
* developer tests to test the account package. In this case it uses a mocked server with the expected responses recorded in JSON files
* integration tests to test the real API endpoints

To run the  tests for the client

```
cd account
go test -v
```

To run tests for the http client

```
cd httpclient
go test -v
```

To run the integration tests
```
cd integrationTests
SERVER_URL={your_server_url} go test 
```

## Issues
* When creating a new account a couple of fields marked as deprecated in the online docs are actually still available on the endpoint:
  * first_name
  * alternative_bank_account_names
  * bank_account_name
* When creating a new account some fields are not persisted in the database:
  * name
  * alternative_names
  * switched
* The IBAN field is processed by the API when creating an account (an error is returned when an invalid IBAN is sent) but is not persisted in the database.
* When trying to delete an existing account but using a wrong version a 404 reponse status is returned. A 409 Conflict response status is expected.

## References
This was my first experience with Go so I had to go through different resources to speed up my learning process. Here are some of websites I have used in the process.

* https://golang.org/doc/tutorial

* https://golang.org/pkg

* https://blog.logrocket.com/making-http-requests-in-go

* https://dev.to/plutov/writing-rest-api-client-in-go-3fkg

* https://tutorialedge.net/golang/reading-console-input-golang

* https://mholt.github.io/json-to-go

* https://blog.alexellis.io/golang-writing-unit-tests

* https://blog.questionable.services/article/testing-http-handlers-go

* https://www.ardanlabs.com/blog/2019/03/integration-testing-in-go-executing-tests-with-docker.html
