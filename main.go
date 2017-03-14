package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc(pathV1Message, MessageHandler).Methods(http.MethodPost)

	n := negroni.Classic()

	n.UseHandler(router)
	n.Run(":3000")
}
