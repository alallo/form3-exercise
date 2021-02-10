FROM golang:latest 
ENV GO111MODULE=auto

WORKDIR /go/src/form3-interview/
COPY . .

RUN go get github.com/google/uuid

ENTRYPOINT  ["go", "test", "-v", "./integrationTests", "-coverprofile", "cover.out"]