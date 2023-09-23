package api

type mediaPayload struct {
	Link    string `json:"link"`
	Caption string `json:"caption"`
}

func NewMediaPayload(url string, caption string) *mediaPayload {
	return &mediaPayload{
		Link:    url,
		Caption: caption,
	}
}
