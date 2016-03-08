package main

import (
	"fmt"

	"github.com/takaaki-mizuno/goji-boilerplate/app"

	"flag"

	"github.com/hypebeast/gojistatic"
	"github.com/takaaki-mizuno/goji-boilerplate/app/services"
	"github.com/zenazn/goji"
)

func main() {
	services.ConfigService().LoadConfigFile("config.yaml")

	staticFolder := services.ConfigService().GetStaticFolder()
	fmt.Println("Static folder : " + staticFolder)
	goji.Use(gojistatic.Static(staticFolder, gojistatic.StaticOptions{SkipLogging: false}))

	port := services.ConfigService().GetServerPort()
	flag.Set("bind", port)

	app.App() // Start main app

	goji.Serve() // Start goji app
}
