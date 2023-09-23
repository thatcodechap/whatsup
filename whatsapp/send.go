package whatsapp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"whatsup/config"
)

func SendMessage(message map[string]interface{}) {
	body, err := json.Marshal(message)
	if err == nil {
	}
	request, err := http.NewRequest("POST", config.ENDPOINT, bytes.NewReader(body))
	request.Header.Set("Authorization", "Bearer "+config.ACCESS_TOKEN)
	request.Header.Set("Content-Type", "application/json")
	http.DefaultClient.Do(request)
}

func CreateMessage(to string, Type string, payload interface{}) map[string]interface{} {
	message := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                to,
		"type":              Type,
	}
	message[Type] = payload
	return message
}
