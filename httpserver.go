package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	rendererr "github.com/jonasala/render-errors"
)

func httpServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(10 * time.Second))

	router.Post("/nova-notificacao", func(w http.ResponseWriter, r *http.Request) {
		mensagem := &Mensagem{}
		if err := render.Bind(r, mensagem); err != nil {
			render.Render(w, r, rendererr.ErrInvalidRequest(err))
			return
		}
		response, err := enviarMensagem(r.Context(), mensagem)
		if err != nil {
			log.Panicf("erro ao enviar notificação: %v\n", err)
		}
		w.Write([]byte(response))
	})

	return router
}
