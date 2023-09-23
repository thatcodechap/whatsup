package main

import (
	"net/http"
	"whatsup/webhook"
)

const PORT string = "8080"

func main() {
	http.HandleFunc("/webhook", webhook.Handle)
	http.ListenAndServe(":"+PORT, nil)
}
