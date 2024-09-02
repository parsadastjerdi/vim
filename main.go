package main

import (
	govapi "github.com/yldshv/go-valorant-api"
)

func main() {
	vapi := govapi.New()

	account, err := vapi.GetAccountByName(govapi.GetAccountByNameParams{
		Name: "todo",
		Tag:  "todo",
	})

	if err != nil {
		print("Error: {}", err)
	}

	print("account {}", account.Data.LastUpdate)
}
