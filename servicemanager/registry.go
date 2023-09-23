package servicemanager

import (
	"whatsup/api"
	"whatsup/services/helloworld"
	"whatsup/services/help"
)

var Registry = map[string]api.Service{
	"helloworld": helloworld.Service,
	"help":       help.Service,
}
