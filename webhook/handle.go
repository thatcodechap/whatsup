package webhook

import (
	"encoding/json"
	"net/http"
	"whatsup/config"
)

func Handle(res http.ResponseWriter, req *http.Request) {
	if isVerificationRequest(req) {
		verifyWebhook(req, res)
	} else if isJson(req) {
		var notification Notification
		json.NewDecoder(req.Body).Decode(&notification)
		handleNotification(&notification)
	}
}

func isVerificationRequest(req *http.Request) bool {
	query := req.URL.Query()
	if req.Method == "GET" && query.Has("hub.mode") && query.Has("hub.challenge") && query.Has("hub.verify_token") {
		return true
	}
	return false
}

func verifyWebhook(req *http.Request, res http.ResponseWriter) {
	token := req.URL.Query().Get("hub.verify_token")
	challenge := req.URL.Query().Get("hub.challenge")
	if token == config.VERIFY_TOKEN {
		res.Write([]byte(challenge))
	} else {
		res.WriteHeader(400)
	}
}

func isJson(req *http.Request) bool {
	if req.Method == "POST" && req.Header.Get("Content-Type") == "application/json" {
		return true
	}
	return false
}
