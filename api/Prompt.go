package api

type Prompt struct {
	From       string
	Service    string
	Parameters []string
}

func NewPrompt(from string, service string, parameters []string) *Prompt {
	return &Prompt{
		From:       from,
		Service:    service,
		Parameters: parameters,
	}
}
