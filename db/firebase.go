package db

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
)

var ctx = context.Background()
var opt = option.WithCredentialsFile("db/key.json") // SDK key

func Firebase() {
	config := &firebase.Config{ProjectID: "<FIREBASE_PROJECT_ID>"}
	println("[HJLOG] config : ", config)

	var app, err = firebase.NewApp(ctx, nil, opt)
	fmt.Printf("[HJLOG] app : %s , err : %s\n", app, err)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	fmt.Printf("[HJLOG] client : %s , err : %s\n", client, err)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

}
