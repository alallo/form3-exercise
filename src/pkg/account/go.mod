module form3.com/account

go 1.15

replace form3.com/httpclient => ../../internal/httpclient

replace form3.com/models => ../../models

require (
	form3.com/httpclient v0.0.0-00010101000000-000000000000
	form3.com/models v0.0.0-00010101000000-000000000000
)
