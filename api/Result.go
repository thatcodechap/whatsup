package api

type Result struct {
	Type    string
	Payload interface{}
}

func NewResult(Type string, payload interface{}) *Result {
	return &Result{
		Type:    Type,
		Payload: payload,
	}
}
