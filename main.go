package main

// 39342 total customer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"segment/version"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// Banner text
const (
	// TETRACON banner
	TETRACON = `
_________    __
|__    __|   | |
   |  |  ___ | |_   ____  ___   ___ ___  _ __ 
   |  | / _ \|  _| /  __|/ _ \ / __/ _ \| '_ \
   |  | \ __/| |_  | |  | (_| | (_| (_) | | | | 
   |__| \___| \__| |_|   \__,_|\___\___/|_| |_| 
Server-version: %s Model-version: %s Model-date: %s
`
)

//
var (
	srv  bool
	vrsn bool
	date bool
)

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

// init documwentation
func init() {

	// instanciate a new logger
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.Level = logrus.DebugLevel
	color.Set(color.FgHiGreen)
	fmt.Fprint(os.Stderr, fmt.Sprintf(TETRACON, version.ServerVersion(), version.ModelVersion(), version.ModelDate()))
	color.Unset()
}

// main
func main() {
	fmt.Println("Hello world!")

	//company.BankID = true

	fileName := "companies.csv"
	startTime1 := time.Now()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//
	lineCount := 0
	//companies := []models.CompaniesType{}

	scanner := bufio.NewScanner(file)
	companies := CompaniesType{}
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		items := strings.Split(scanner.Text(), ";")
		company := CompanyType{}
		for i := range items {
			//fmt.Println(items[i])
			company.BusinessPartner = items[0]
			company.DateUsed = items[1]
			company.BankID = items[2]
			//sNotLoggedIn := items[3]
			if strings.Contains(items[3], "TRUE") {
				company.NotLoggedIn = true
			} else {
				company.NotLoggedIn = false
			}
			company.Period = items[i]
		}
		companies.AddItem(company) // Add company

		/*
			fmt.Printf("BP: %s, date: %s - BankID: %s - NotLoggedIn: %b\r\n",
				company.BusinessPartner,
				company.DateUsed,
				company.BankID,
				company.NotLoggedIn)
		*/
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//printCompanies(companies)
	printNumberOfNotLoggedIn(companies)
	printNumberOfNoBankID(companies)
	printLoggedInPeriod(companies, "201601")
	printLoggedInPeriod(companies, "201602")
	printLoggedInPeriod(companies, "201603")
	printLoggedInPeriod(companies, "201604")
	printLoggedInPeriod(companies, "201605")
	printLoggedInPeriod(companies, "201606")
	printLoggedInPeriod(companies, "201607")
	printLoggedInPeriod(companies, "201608")
	printLoggedInPeriod(companies, "201609")
	printLoggedInPeriod(companies, "201610")
	printLoggedInPeriod(companies, "201611")
	printLoggedInPeriod(companies, "201612")
	printLoggedInPeriod(companies, "201701")
	printLoggedInPeriod(companies, "201702")
	printLoggedInPeriod(companies, "201703")
	printLoggedInPeriod(companies, "201704")
	printLoggedInPeriod(companies, "201705")
	printLoggedInPeriod(companies, "201706")
	printLoggedInPeriod(companies, "201707")
	printLoggedInPeriod(companies, "201708")
	printLoggedInPeriod(companies, "201709")
	printLoggedInPeriod(companies, "201710")
	printLoggedInPeriod(companies, "201711")
	printLoggedInPeriod(companies, "201712")
	printLoggedInPeriod(companies, "201801")
	printLoggedInPeriod(companies, "201802")
	printLoggedInPeriod(companies, "201803")
	printLoggedInPeriod(companies, "201804")
	printLoggedInPeriod(companies, "201805")
	printLoggedInPeriod(companies, "201806")
	printLoggedInPeriod(companies, "201807")
	printLoggedInPeriod(companies, "201808")
	printLoggedInPeriod(companies, "201809")
	printLoggedInPeriod(companies, "201810")
	printLoggedInPeriod(companies, "201811")
	printLoggedInPeriod(companies, "201812")
	printLoggedInPeriod(companies, "201901")
	printLoggedInPeriod(companies, "201902")
	printLoggedInPeriod(companies, "201903")
	printLoggedInPeriod(companies, "201904")
	printLoggedInPeriod(companies, "201905")
	printLoggedInPeriod(companies, "201906")
	printLoggedInPeriod(companies, "201907")
	printLoggedInPeriod(companies, "201908")
	printLoggedInPeriod(companies, "201909")
	printLoggedInPeriod(companies, "201910")
	printLoggedInPeriod(companies, "201911")
	printLoggedInPeriod(companies, "201912")
	printLoggedInPeriod(companies, "202001")
	printNAPeriod(companies, "#N/A")
	fmt.Printf("- File: %s, # of lines: %d, processing time: %s \r\n",
		fileName, lineCount, time.Since(startTime1))

	//
}

// AddItem description
func (companies *CompaniesType) AddItem(company CompanyType) {

	companies.Companies = append(companies.Companies, company)

}

// numberOfNotSignedIn description
func printNumberOfNotLoggedIn(companies CompaniesType) {

	n := 0
	for _, company := range companies.Companies {
		if company.NotLoggedIn {
			n++
		}
	}
	fmt.Printf("Customer: - Number of customer not LoggedIn to IB: {%d}\r\n", n)

}

// numberOfNoBankID description
func printNumberOfNoBankID(companies CompaniesType) {

	n := 0
	for _, company := range companies.Companies {
		if company.BankID == "0" {
			n++
		}
	}
	fmt.Printf("Customer: - Number of customer without BankID: {%d}\r\n", n)
}

// printLoggedInPeriod documentation
func printLoggedInPeriod(companies CompaniesType, period string) {

	n := 0
	for _, company := range companies.Companies {
		if company.Period == period {
			n++
		}
	}
	fmt.Printf("Period: %s - Number of digital customer : {%d}\r\n", period, n)
}

// printNAPeriod documentation
func printNAPeriod(companies CompaniesType, period string) {

	n := 0
	for _, company := range companies.Companies {
		if company.Period == period || company.Period == "190001" {
			n++
		}
	}
	fmt.Printf("Period: %s - Number of non digital customer : {%d}\r\n", period, n)
}

// printCompanies documentation
func printCompanies(companies CompaniesType) {

	//company := CompanyType{}

	for i, company := range companies.Companies {
		fmt.Printf("id: {%d} Company: {%s} Period: {%s} BankID: {%s} NotLoggedIn: {%b}\r\n",
			i, company.BusinessPartner, company.Period, company.BankID,
			company.NotLoggedIn)
	}
}
