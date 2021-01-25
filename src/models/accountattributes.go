package models

type AccountAttributes struct {
	// ISO 3166-1 code used to identify the domicile of the account, e.g. 'GB', 'FR'
	Country string `json:"country"`

	// ISO 4217 code used to identify the base currency of the account, e.g. 'GBP', 'EUR'
	BaseCurrency string `json:"base_currency"`

	// Account number. A unique account number will automatically be generated if not provided. If provided, the account number is not validated.
	AccountNumber string `json:"account_number"`

	// Local country bank identifier. Format depends on the country. Used for most countries.
	BankID string `json:"bank_id"`

	// Identifies the type of bank ID being used. Value depends on country attribute.
	BankIDCode string `json:"bank_id_code"`

	// SWIFT BIC in either 8 or 11 character format e.g. 'NWBKGB22'
	Bic string `json:"bic"`

	// IBAN of the account. Will be calculated from other fields if not supplied.
	Iban string `json:"iban"`

	// A free-format reference that can be used to link this account to an external system
	CustomerID string `json:"customer_id"`

	// Name of the account holder, up to four lines possible.
	Name [4]string `json:"name"`

	// Alternative primary account names, only used for UK Confirmation of Payee
	AlternativeNames [3]string `json:"alternative_names"`

	// Classification of account, only used for Confirmation of Payee (CoP)
	AccountClassification string `json:"account_classification"`

	// Flag to indicate if the account is a joint account, only used for Confirmation of Payee (CoP)
	JointAccount bool `json:"joint_account"`

	// Flag to indicate if the account has been switched away from this organisation, only used for Confirmation of Payee (CoP)
	Switched bool `json:"switched"`

	// Flag to indicate if the account has opted out of account matching, only used for Confirmation of Payee
	AccountMatchingOptOut bool `json:"account_matching_opt_out"`

	// Status of the account. Inferred from the status field of the newest Account Event resource associated with the account. Always confirmed for older accounts where no Account Event resources are present.ß
	Status string `json:"status"`

	// Additional information to identify the account and account holder, only used for Confirmation of Payee (CoP)
	SecondaryIdentification string `json:"secondary_identification"`

	// Account holder identification if the account holder is a private entity
	PrivateIdentification string `json:"private_identification"`

	// The birth date of the account holder
	PrivateIdentificationBirthDate string `json:"private_identification.birth_date"`

	// The birth country of the account holder
	PrivateIdentificationBirthCountry string `json:"private_identification.birth_country"`

	// The number of the document used to identify the account holder
	PrivateIdentificationIdentification string `json:"private_identification.identification"`

	// The street name and house number of the postal address of the account holder.
	PrivateIdentificationAddress string `json:"private_identification.address"`

	// The city where the postal address of the account holder is located
	PrivateIdentificationCity string `json:"private_identification.city"`

	// The country where the postal address of the account holder is located
	PrivateIdentificationCountry string `json:"private_identification.country"`

	// Account holder identification if the account holder is an organisation
	OrganisationIdentification string `json:"organisation_identification"`

	// The registration number used to identify the account holding organisation.
	OrganisationIdentificationIdentification string `json:"organisation_identification.identification"`

	// The street name and house number of the postal address of the account holding organisation.
	OrganisationIdentificationAddress string `json:"organisation_identification.address"`

	// The city where the postal address of the account holding organisation is located
	OrganisationIdentificationCity string `json:"organisation_identification.city"`

	// The country where the postal address of the account holding organisation is located
	OrganisationIdentificationCountry string `json:"organisation_identification.country"`

	// The name of a person representing the account holding organisation
	OrganisationIdentificationActorsName string `json:"organisation_identification.actors.name"`

	// The birth date of the person named in organisation_identification.actors.name
	OrganisationIdentificationActorsBirthDate string `json:"organisation_identification.actors.birth_date"`

	// The country of residency of the person named in organisation_identification.actors.name
	OrganisationIdentificationActorsResidency string `json:"organisation_identification.actors.residency"`

	// The Account Event resources related to this account. Not present in List calls.
	RelationshipAccountEvents string `json:"relationships.account_events"`

	// The master account related to this account
	RelationshipsMasterAccount string `json:"relationships.master_account"`
}
