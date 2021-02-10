module account

go 1.15

replace form3-interview/httpclient => ../httpclient

replace form3-interview/models => ../models

require (
	form3-interview/httpclient v0.0.0-00010101000000-000000000000
	form3-interview/models v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.2.0
)
