package handlers

import (
	"flag"
	"fmt"
	"os"

	"github.com/izzyyhh/zok/internal/httpclient"
)

func handlePost(args []string) {
	if len(args) < 2 {
		fmt.Println("zok expected an url to request for")
		os.Exit(1)
	}

	url := args[len(args)-1]

	postFlagSet := flag.NewFlagSet(POST, flag.ExitOnError)
	postBody := postFlagSet.String("body", "", "HTTP body appended to the request")
	err := postFlagSet.Parse(args[1 : len(args)-1])

	if err != nil {
		fmt.Println("could not parse body")
		os.Exit(1)
	}

	body, err := httpclient.Post(&url, postBody)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
