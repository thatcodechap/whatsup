package help

import (
	"whatsup/api"
)

type service struct{}

var Service service

var help string = `Hey there 👋🏻
Hope you are doing good 🤍
To use the services start the message with a dot -
*.<service> <message to give the service>*
Example: *.helloworld* hello

To get help : .<service> help

Services Available:
*.helloworld* - ChatGPT in whatsapp 😎`

func (service) Execute(*api.Prompt) *api.Result {
	text := api.NewTextPayload(help)
	return api.NewResult("text", text)
}

func (service) Help() *api.Result {
	text := api.NewTextPayload("Okay I can help!")
	return api.NewResult("text", text)
}
