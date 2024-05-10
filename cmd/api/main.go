package main

import (
	"log"
	"main/cmd/api/docs"
	"main/pkg/config"

	di "main/pkg/di"
)

func main() {
	docs.SwaggerInfo.Title = "inventory_management"

	docs.SwaggerInfo.Version = "1.0"

	docs.SwaggerInfo.Host = "localhost:1233"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
	config, configErr := config.LoadConfig()
	if configErr != nil {

		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
