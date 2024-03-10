package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type ModuleInterface interface {
	SwaggerModule()
}

func InitModule(router fiber.Router, server *server) ModuleInterface {
	return &module{}
}

type module struct {
	router fiber.Router
	server *server
	// TODO add middileware
}

// TODO Init middleware

// SwaggerModule implements ModuleInterface.
func (m *module) SwaggerModule() {
	router := m.router.Group("/swagger")

	// http://localhost:8080/v1/swagger
	router.Get("/*", swagger.HandlerDefault) // default

	router.Get("/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))
}
