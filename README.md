# First time writing code in Go 
Please forgive my sins :-) 

# Description 
Take home exercise

# Project structure
   .
    ├── client                  # a Go client to inteface with form3 APIs. This implements some of the "Account" functionalities
    ├── cmd                     # Command line app. Useful to play with the client
    ├── internal                # packages used by the client that we don't want to expose. This includes a wrapper to help handling an http client
    └── scripts                 # docker compose file and DB scripts provided by form3

# Prerequisite
To run the command line tool and test the client you need

* Go https://golang.org/doc/install
* Docker https://www.docker.com/get-started

# Getting started 

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

# Testing
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



# References
This was my first experience with Go so I had to go through different resources to speed up my learning process. Here is a list of websites I have used in the process.

https://golang.org/doc/tutorial/
https://golang.org/pkg
https://blog.logrocket.com/making-http-requests-in-go/
https://dev.to/plutov/writing-rest-api-client-in-go-3fkg
https://tutorialedge.net/golang/reading-console-input-golang/
https://mholt.github.io/json-to-go/
https://blog.alexellis.io/golang-writing-unit-tests/
https://blog.questionable.services/article/testing-http-handlers-go/
