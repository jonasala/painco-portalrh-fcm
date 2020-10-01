package main

import "net/http"

//Mensagem representa as informações recebidas pelo servidor HTTP
type Mensagem struct {
	Titulo   string `json:"titulo"`
	Conteudo string `json:"conteudo"`
}

//Bind satisfaz a interface render.Binder
func (m *Mensagem) Bind(r *http.Request) error {
	return nil
}
