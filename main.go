package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"firebase.google.com/go/messaging"
)

var client *messaging.Client

func main() {
	f, err := os.OpenFile("painco-portalrh-fcm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("não foi possível criar o arquivo de log: %v\n", err)
	}
	log.SetOutput(f)

	portPtr := flag.String("p", "8290", "porta http")
	flag.Parse()

	client, err = messagingClient(context.Background())
	if err != nil {
		log.Fatalf("não foi possível recuperar o messaging client: %v", err)
	}

	router := httpServer()
	log.Printf("painco-portalrh-fcm escutando http na porta %v\n", *portPtr)
	http.ListenAndServe(":"+(*portPtr), router)
}
