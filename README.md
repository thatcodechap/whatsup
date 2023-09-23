# Whatsup
## whatsapp multi service chatbot framework
### Overview
This framework is useful for creating simple whatsapp chatbots with multiple services.
### How to use
- Configure the bot by adding credentials in `config.go` under `config` package
- Add required services
- Add the newly created to registry
#### Follow the [Meta Documentations](https://developers.facebook.com/docs/whatsapp/business-management-api/get-started#system-user-access-tokens) to get your credentials
### How to add a service
- Creat a package under `services` folder
- Declare a package varialbe `Service` of any type
- Implement `Service` interface for `Service` Object
  ```
  type Service interface {
      Execute (*Prompt) *Result
	  Help() *Result
  }
  ```
  `Service`, `Result` and `Prompt` types are defined under `api` package
- Import your package and add to registry map in `registry.go` under `servicemanager`` package
### Example template for a service
```
package newservice

import "whatsup/api"

type service struct{}
var Service service

func (service) Execute(prompt *api.Prompt) *api.Result {
	text := api.NewTextPayload("Hey There")
	return api.NewResult("text", text)
}

func (service) Help() *api.Result {
	text := api.NewTextPayload("Okay I can help!")
	return api.NewResult("text", text)
}
```
### How it works
- When a message is send with a prefix, the service is extracted from it
- Then the service registry is looked up to check whether the service exists
- If the service exists then it is executed and the result from the service is sent back to the user