package servicemanager

import (
	"whatsup/api"
	"whatsup/whatsapp"
)

func sendErrorMessage(to string) {
	errorMessage := whatsapp.CreateMessage(to, "text", api.NewTextPayload("Service Not Avalalbe"))
	whatsapp.SendMessage(errorMessage)
}

func Run(prompt *api.Prompt) {
	defer func() {
		if x := recover(); x != nil {
			sendErrorMessage(prompt.From)
		}
	}()
	service := Registry[prompt.Service]
	var result *api.Result
	if len(prompt.Parameters) == 1 && prompt.Parameters[0] == "help" {
		result = service.Help()
	} else {
		result = service.Execute(prompt)
	}
	message := whatsapp.CreateMessage(prompt.From, result.Type, result.Payload)
	whatsapp.SendMessage(message)
}
