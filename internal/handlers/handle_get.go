package handlers

import (
	"fmt"
	"github.com/izzyyhh/zok/internal/httpclient"
	"os"
)

func handleGet(args []string) {
	if len(args) < 2 {
		fmt.Println("zok expected an url to request for")
		os.Exit(1)
	}

	url := args[1]

	body, err := httpclient.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
