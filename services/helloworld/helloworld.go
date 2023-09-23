package helloworld

import (
	"whatsup/api"
)

type service struct{}

var Service service // implement Service interface

func (service) Execute(*api.Prompt) *api.Result {
	text := api.NewTextPayload("Hello World!")
	return api.NewResult("text", text)
}

func (service) Help() *api.Result {
	text := api.NewTextPayload("Okay I can help!")
	return api.NewResult("text", text)
}
