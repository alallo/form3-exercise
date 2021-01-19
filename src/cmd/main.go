package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"form3.com/accountlist"
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
		}
		if strings.Compare("3", text) == 0 {
			fmt.Println("List")
			var filters accountlist.Filters
			filters.PageNumber = 0
			filters.PageSize = 100
			filters.BankID = "1234"
			filters.AccountNumber = "898888"
			resp, err := accountlist.GetAccountList("http://localhost:8080/v1/organisation/accounts", &filters)
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
