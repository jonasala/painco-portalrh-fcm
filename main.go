package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"firebase.google.com/go/messaging"
)

var client *messaging.Client

func main() {
	httpPort := os.Getenv("PORTALRH_FCM_HTTP_PORT")
	if _, err := strconv.Atoi(httpPort); err != nil {
		log.Fatalf("erro com a env PORTALRH_FCM_HTTP_PORT: %v", err)
	}

	var err error
	client, err = messagingClient(context.Background())
	if err != nil {
		log.Fatalf("não foi possível recuperar o messaging client: %v", err)
	}

	router := httpServer()
	log.Printf("painco-portalrh-fcm escutando http na porta %v\n", httpPort)
	http.ListenAndServe(":"+httpPort, router)
}
