package main

import (
	"flag"

	"github.com/kusold/facebreak/config"
	"github.com/kusold/facebreak/facebook"
)

// Config stores the flags

func main() {
	config := &config.Config{
		AccessToken: flag.String("access", "", "Visit https://developers.facebook.com/tools/explorer/ to get an Access Token"),
	}
	flag.Parse()

	var client backupClient
	client = facebook.NewClient(config)
	run(client)
}

type backupClient interface {
	Run() error
}

func run(client backupClient) {
	client.Run()
}
