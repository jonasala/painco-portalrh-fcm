package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"firebase.google.com/go/messaging"
)

var client *messaging.Client

func main() {
	portPtr := flag.String("p", "8290", "porta http")

	flag.Parse()

	fmt.Println(*portPtr)

	var err error
	client, err = messagingClient(context.Background())
	if err != nil {
		log.Fatalf("não foi possível recuperar o messaging client: %v", err)
	}

	router := httpServer()
	log.Printf("painco-portalrh-fcm escutando http na porta %v\n", *portPtr)
	http.ListenAndServe(":"+(*portPtr), router)
}
