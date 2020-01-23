package models

// CompanyType description
type CompanyType struct {
	BusinessPartner string `json:"BusinessPartner"`
	DateUsed        string `json:"dateUsed"`
	BankID          string `json:"bankID"`
	NotLoggedIn     bool   `json:"notLoggedIn"`
	Period          string `json:"Period"`
}
