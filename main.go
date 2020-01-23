package main

import (
	"bufio"
	"fmt"
	"ftg-count/models"
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
	//var companies []models.CompanyType
	var company models.CompanyType

	//company.BankID = true

	fileName := "companies.csv"
	startTime1 := time.Now()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		items := strings.Split(scanner.Text(), ";")
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
		fmt.Printf("BP: %s, date: %s - BankID: %s - NotLoggedIn: %b\r\n",
			company.BusinessPartner,
			company.DateUsed,
			company.BankID,
			company.NotLoggedIn)
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("- File: %s, # of lines: %d, processing time: %s \r\n",
		fileName, lineCount, time.Since(startTime1))
	//
}
