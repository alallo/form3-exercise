# First time writing code in Go 
Name: Alessandro Lallo (alessandro.lallo@gmail.com)
Please forgive my sins :-) 

![Image of cat](https://i.pinimg.com/474x/77/ad/93/77ad9387b0e57423b3e00b28116cd393.jpg)

## Description 
Take home exercise

## Project structure

Folder Name | Description
------------ | -------------
client | a Go client to inteface with form3 APIs. This implements some of the "Account" functionalities
cmd | Command line app. Useful to play with the client
internal | packages used by the client that we don't want to expose. This includes a wrapper to help handling an http client
scripts | docker compose file and DB scripts provided by form3

## Prerequisite
To run the command line tool and test the client you need

* Go https://golang.org/doc/install
* Docker https://www.docker.com/get-started

## Try it

First of all use the docker compose file from the scripts folder to spin up the form3 API

```
cd scripts
docker-compose up
```

Now you can run the command line tool and interact with the client running

```
SERVER_URL={your_server_url} HOST={your_host} go run main.go
```

because we are using the docker-compose file and not the real API endpoints, this is what we need

```
cd cmd
SERVER_URL=http://localhost:8080 HOST=http://localhost:8080 go run main.go
```

## Example
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
To run tests for the client

```
cd client/account
go test
```

To run tests for the http client

```
cd internal/httpclient
go test
```



## References
This was my first experience with Go so I had to go through different resources to speed up my learning process. Here is a list of websites I have used in the process.

* https://golang.org/doc/tutorial

* https://golang.org/pkg

* https://blog.logrocket.com/making-http-requests-in-go

* https://dev.to/plutov/writing-rest-api-client-in-go-3fkg

* https://tutorialedge.net/golang/reading-console-input-golang

* https://mholt.github.io/json-to-go

* https://blog.alexellis.io/golang-writing-unit-tests

* https://blog.questionable.services/article/testing-http-handlers-go
