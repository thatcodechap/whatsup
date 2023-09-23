package api

type Service interface {
	Execute(*Prompt) *Result
	Help() *Result
}
