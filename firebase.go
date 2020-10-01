package main

import (
	"context"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func messagingClient(ctx context.Context) (*messaging.Client, error) {
	opt := option.WithCredentialsFile("dev-painco-portal-rh-firebase-adminsdk-p87gf-5367ee5ab2.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	return app.Messaging(ctx)
}

func enviarMensagem(ctx context.Context, mensagem *Mensagem) (string, error) {
	ttl, _ := time.ParseDuration("72h")

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: mensagem.Titulo,
			Body:  mensagem.Conteudo,
		},
		Android: &messaging.AndroidConfig{
			TTL: &ttl,
		},
		Topic: "app",
	}

	return client.Send(ctx, message)
}
