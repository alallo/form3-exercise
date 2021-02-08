package models

type AccountAttributes struct {
	// ISO 3166-1 code used to identify the domicile of the account, e.g. 'GB', 'FR'
	Country string `json:"country,omitempty"`

	// ISO 4217 code used to identify the base currency of the account, e.g. 'GBP', 'EUR'
	BaseCurrency string `json:"base_currency,omitempty"`

	// Account number. A unique account number will automatically be generated if not provided. If provided, the account number is not validated.
	AccountNumber string `json:"account_number,omitempty"`

	// Local country bank identifier. Format depends on the country. Used for most countries.
	BankID string `json:"bank_id,omitempty"`

	// Identifies the type of bank ID being used. Value depends on country attribute.
	BankIDCode string `json:"bank_id_code,omitempty"`

	// SWIFT BIC in either 8 or 11 character format e.g. 'NWBKGB22'
	Bic string `json:"bic,omitempty"`

	// IBAN of the account. Will be calculated from other fields if not supplied.
	Iban *string `json:"iban,omitempty"`

	// A free-format reference that can be used to link this account to an external system
	CustomerID string `json:"customer_id,omitempty"`

	// Name of the account holder, up to four lines possible.
	Name [4]string `json:"name"`

	// Alternative primary account names, only used for UK Confirmation of Payee
	AlternativeNames [3]string `json:"alternative_names"`

	// Classification of account, only used for Confirmation of Payee (CoP)
	AccountClassification *string `json:"account_classification,omitempty"`

	// Flag to indicate if the account is a joint account, only used for Confirmation of Payee (CoP)
	JointAccount bool `json:"joint_account,omitempty"`

	// Flag to indicate if the account has been switched away from this organisation, only used for Confirmation of Payee (CoP)
	Switched bool `json:"switched,omitempty"`

	// Flag to indicate if the account has opted out of account matching, only used for Confirmation of Payee
	AccountMatchingOptOut bool `json:"account_matching_opt_out,omitempty"`

	// Status of the account. Inferred from the status field of the newest Account Event resource associated with the account. Always confirmed for older accounts where no Account Event resources are present.ß
	Status string `json:"status,omitempty"`

	// Additional information to identify the account and account holder, only used for Confirmation of Payee (CoP)
	SecondaryIdentification string `json:"secondary_identification,omitempty"`
}
