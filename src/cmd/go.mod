module main

go 1.15

replace form3.com/httpclient => ../internal/httpclient

replace form3.com/account => ../client/account

replace form3.com/models => ../models

require (
	form3.com/account v0.0.0-00010101000000-000000000000
	form3.com/models v0.0.0-00010101000000-000000000000
	github.com/go-delve/delve v1.6.0 // indirect
	github.com/google/uuid v1.2.0
)
