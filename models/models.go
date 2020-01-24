package models

// NOT USED, NOT USED, NOT USED
// CompanyType description
type CompanyType struct {
	BusinessPartner string `json:"BusinessPartner"`
	DateUsed        string `json:"dateUsed"`
	BankID          string `json:"bankID"`
	NotLoggedIn     bool   `json:"notLoggedIn"`
	Period          string `json:"Period"`
}

// CompaniesType description
type CompaniesType struct {
	Companies []CompanyType
}
