package api

type textPayload struct {
	Body    string `json:"body"`
	Preview string `json:"preview_url"`
}

func NewTextPayload(text string) *textPayload {
	return &textPayload{
		Body:    text,
		Preview: "true",
	}
}
