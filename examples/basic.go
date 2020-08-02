package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rsclarke/go-upcloud/upcloud"
)

func main() {

	tp := upcloud.BasicAuthTransport{
		Username: os.Getenv("UPCLOUD_USERNAME"),
		Password: os.Getenv("UPCLOUD_PASSWORD"),
	}

	client := upcloud.NewClient(tp.Client())
	ctx := context.Background()

	account, _, err := client.Accounts.GetAccountDetails(ctx, "someUser")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v", account)
}
