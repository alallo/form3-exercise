package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"form3-interview/account"
	"form3-interview/models"

	"github.com/google/uuid"
)

func main() {
	serverURL := os.Getenv("SERVER_URL")
	host := os.Getenv("HOST")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Form3 Accounts API Console")
	fmt.Println("---------------------")
	fmt.Println("Select one of the following options and press enter:")
	fmt.Println("1. Create a new account")
	fmt.Println("2. Fetch an existing account")
	fmt.Println("3. List of Accounts")
	fmt.Println("4. Delete an account")
	fmt.Println("Enter to exit")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("1", text) == 0 {
			fmt.Println("Create")
			for {
				var newAccountAttrs models.AccountAttributes

				fmt.Print("Country: ")
				country, _ := reader.ReadString('\n')
				country = strings.Replace(country, "\n", "", -1)
				newAccountAttrs.Country = country

				fmt.Print("BaseCurrency: ")
				baseCurrency, _ := reader.ReadString('\n')
				baseCurrency = strings.Replace(baseCurrency, "\n", "", -1)
				newAccountAttrs.BaseCurrency = baseCurrency

				fmt.Print("BankID: ")
				bankID, _ := reader.ReadString('\n')
				bankID = strings.Replace(bankID, "\n", "", -1)
				newAccountAttrs.BankID = bankID

				fmt.Print("BankIDCode: ")
				bankIDCode, _ := reader.ReadString('\n')
				bankIDCode = strings.Replace(bankIDCode, "\n", "", -1)
				newAccountAttrs.BankIDCode = bankIDCode

				fmt.Print("Bic: ")
				bic, _ := reader.ReadString('\n')
				bic = strings.Replace(bic, "\n", "", -1)
				newAccountAttrs.Bic = bic

				fmt.Print("AccountNumber: ")
				accountNumber, _ := reader.ReadString('\n')
				accountNumber = strings.Replace(accountNumber, "\n", "", -1)
				newAccountAttrs.AccountNumber = accountNumber

				fmt.Print("CustomerID: ")
				customerID, _ := reader.ReadString('\n')
				customerID = strings.Replace(customerID, "\n", "", -1)
				newAccountAttrs.CustomerID = customerID

				var newAccount models.Account
				newAccount.ID = uuid.New()
				newAccount.Type = "accounts"
				newAccount.OrganisationID = uuid.New()
				newAccount.Attributes = &newAccountAttrs

				var newData account.Data
				newData.Account = &newAccount

				var req account.CreateRequest
				req.Host = host
				req.Data = &newData

				createAccount(serverURL, req)
				break
			}
		}
		if strings.Compare("2", text) == 0 {
			fmt.Println("Fecth")
			for {
				fmt.Print("Account ID: ")
				accountID, _ := reader.ReadString('\n')
				accountID = strings.Replace(accountID, "\n", "", -1)

				var req account.FetchRequest
				req.AccountID = accountID
				req.Host = host
				fecthAccount(host, serverURL, req)
				break
			}
		}
		if strings.Compare("3", text) == 0 {
			fmt.Println("List")
			for {
				fmt.Print("Page Number: ")
				pageNumber, _ := reader.ReadString('\n')
				pageNumber = strings.Replace(pageNumber, "\n", "", -1)
				pageNumberInt, _ := strconv.Atoi(pageNumber)

				fmt.Print("Page Size: ")
				pageSize, _ := reader.ReadString('\n')
				pageSize = strings.Replace(pageSize, "\n", "", -1)
				pageSizeInt, _ := strconv.Atoi(pageSize)

				var req account.ListRequest
				req.PageNumber = pageNumberInt
				req.PageSize = pageSizeInt
				req.Host = host
				accountList(serverURL, req)
				break
			}

		}
		if strings.Compare("4", text) == 0 {
			fmt.Println("Delete")
			for {
				fmt.Print("Account ID: ")
				accountID, _ := reader.ReadString('\n')
				accountID = strings.Replace(accountID, "\n", "", -1)

				fmt.Print("Version: ")
				version, _ := reader.ReadString('\n')
				version = strings.Replace(version, "\n", "", -1)
				versionInt, _ := strconv.Atoi(version)

				var req account.DeleteRequest
				req.Host = host
				req.AccountID = accountID
				req.Version = versionInt
				deleteAccount(host, serverURL, req)
				break
			}

		}
		if strings.Compare("", text) == 0 {
			fmt.Println("Bye!")
			break
		}
	}

}

func createAccount(serverURL string, req account.CreateRequest) {
	resp, err := account.CreateAccount(serverURL, &req)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		body, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(string(body))
	}
}

func accountList(serverURL string, req account.ListRequest) {
	resp, err := account.GetAccountList(serverURL, &req)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		body, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(string(body))
	}
}

func deleteAccount(host string, serverURL string, req account.DeleteRequest) {

	err := account.DeleteAccount(serverURL, &req)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Account deleted succesfuly")
	}
}

func fecthAccount(host string, serverURL string, req account.FetchRequest) {
	resp, err := account.GetAccount(serverURL, &req)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		body, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(string(body))
	}
}
