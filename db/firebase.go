package db

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Firebase() {
	// SDK key
	opt := option.WithCredentialsFile("key.json")

	var app, err = firebase.NewApp(context.Background(), nil, opt)
	fmt.Printf("HJLOG : app : %s , err : %s\n", app, err)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(context.Background())
	fmt.Printf("HJLOG : client : %s , err : %s\n", client, err)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

}
