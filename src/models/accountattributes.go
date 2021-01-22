package models

type AccountAttributes struct {
	Country               string `json:"country"`
	BaseCurrency          string `json:"base_currency"`
	AccountNumber         string `json:"account_number"`
	BankID                string `json:"bank_id"`
	BankIDCode            string `json:"bank_id_code"`
	Bic                   string `json:"bic"`
	Iban                  string `json:"iban"`
	AccountClassification string `json:"account_classification"`
	JointAccount          bool   `json:"joint_account"`
	Switched              bool   `json:"switched"`
	AccountMatchingOptOut bool   `json:"account_matching_opt_out"`
	Status                string `json:"status"`
}
