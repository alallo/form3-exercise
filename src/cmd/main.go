package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"form3.com/account"
	"form3.com/models"
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

			var newAccountAttrs models.AccountAttributes
			newAccountAttrs.Country = "GB"
			newAccountAttrs.BaseCurrency = "GBP"
			newAccountAttrs.BankID = "400300"
			newAccountAttrs.BankIDCode = "GBSDC"
			newAccountAttrs.Bic = "NWBKGB22"

			var newAccount models.Account
			newAccount.ID = uuid.New()
			newAccount.Type = "accounts"
			newAccount.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
			newAccount.Attributes = &newAccountAttrs

			var newData account.Data
			newData.Account = &newAccount

			var req account.AccountCreateRequest
			req.Host = host
			req.Data = &newData

			resp, err := account.CreateAccount(serverURL, &req)
			if err != nil {
				fmt.Println("Error: ", err)
				break
			} else {
				fmt.Println(resp.ID)
			}
		}
		if strings.Compare("2", text) == 0 {
			fmt.Println("Fecth")
			for {
				fmt.Print("Account ID: ")
				accountID, _ := reader.ReadString('\n')
				accountID = strings.Replace(accountID, "\n", "", -1)
				fecthAccount(host, serverURL, accountID)
				break
			}

		}
		if strings.Compare("3", text) == 0 {
			fmt.Println("List")
			var req account.AccountListRequest
			req.PageNumber = 0
			req.PageSize = 100
			req.Host = host
			req.BankID = []string{"1234", "456", "8963"}
			req.AccountNumber = []string{"898888, 11111, 2222"}
			resp, err := account.GetAccountList(serverURL, &req)
			if err != nil {
				fmt.Println("Error: ", err)
				break
			} else {
				body, err := json.MarshalIndent(resp, "", "  ")
				if err != nil {
					fmt.Println("Error: ", err)
				}
				fmt.Println(string(body))
			}

		}
		if strings.Compare("4", text) == 0 {
			fmt.Println("Delete")
			var req account.AccountDeleteRequest
			req.Host = host
			req.AccountID = "c93d6404-8990-4c6b-81f8-7ce67533733d"
			req.Version = 0
			err := account.DeleteAccount(serverURL, &req)
			if err != nil {
				fmt.Println("Error: ", err)
				break
			} else {
				fmt.Println("Account deleted succesfuly")
			}

		}
		if strings.Compare("", text) == 0 {
			fmt.Println("Bye!")
			break
		}
	}

}

func fecthAccount(host string, serverURL string, accountID string) {
	var req account.AccountFetchRequest
	req.AccountId = accountID
	req.Host = host
	resp, err := account.GetAccount(serverURL, &req)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(resp.ID)
	}
}
