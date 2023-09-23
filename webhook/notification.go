package webhook

import (
	"strings"
	"whatsup/api"
	"whatsup/config"
	"whatsup/servicemanager"
	"whatsup/utils"
)

type Notification struct {
	Entry [1]struct {
		Changes [1]struct {
			Value struct {
				Contacts [1]map[string]interface{} `json:"contacts"`
				Messages [1]map[string]interface{} `json:"messages"`
				Statuses [1]map[string]interface{} `json:"statuses"`
			} `json:"value"`
		} `json:"changes"`
	} `json:"entry"`
}

func handleNotification(notification *Notification) {
	message := parseMessage(notification)
	if message == nil {
		return
	}
	from := message["from"].(string)
	if message["type"] == "text" {
		payload := message["text"].(map[string]interface{})
		handleTextMessage(from, payload)
	}
}

func parseMessage(notification *Notification) (message map[string]interface{}) {
	defer func() {
		if x := recover(); x != nil {
			message = nil
		}
	}()
	return notification.Entry[0].Changes[0].Value.Messages[0]
}

func parsePrompt(text string) (string, []string) {
	text = utils.RemoveMultipleSpaces(text)
	words := strings.Split(text, " ")
	return words[0][1:], words[1:]
}

func handleTextMessage(from string, payload map[string]interface{}) {
	var prompt *api.Prompt
	if payload["body"].(string)[0] == config.PREFIX {
		service, parameters := parsePrompt(payload["body"].(string))
		prompt = api.NewPrompt(from, service, parameters)
	} else {
		prompt = api.NewPrompt(from, "help", nil)
	}
	servicemanager.Run(prompt)
}
