package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"form3.com/account"
)

func main() {

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
		}
		if strings.Compare("2", text) == 0 {
			fmt.Println("Fecth")
			var req account.AccountFetchRequest
			req.AccountId = "c93d6404-8990-4c6b-81f8-7ce67533733d"
			req.Host = "api.form3.tech"
			resp, err := account.GetAccount("http://localhost:8080", &req)
			if err != nil {
				fmt.Println("Error: ", err)
				break
			} else {
				fmt.Println(resp.ID)
			}
		}
		if strings.Compare("3", text) == 0 {
			fmt.Println("List")
			var req account.AccountListRequest
			req.PageNumber = 0
			req.PageSize = 100
			req.Host = "api.form3.tech"
			req.BankID = []string{"1234", "456", "8963"}
			req.AccountNumber = []string{"898888, 11111, 2222"}
			resp, err := account.GetAccountList("http://localhost:8080", &req)
			if err != nil {
				fmt.Println("Error: ", err)
				break
			} else {
				fmt.Println(resp.Accounts[0].Type)
			}

		}
		if strings.Compare("4", text) == 0 {
			fmt.Println("Delete")
		}
		if strings.Compare("", text) == 0 {
			fmt.Println("Bye!")
			break
		}
	}

}
