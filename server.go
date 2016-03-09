package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/hypebeast/gojistatic"
	"github.com/takaaki-mizuno/goji-boilerplate/app"
	"github.com/takaaki-mizuno/goji-boilerplate/app/services"
	"github.com/zenazn/goji"
)

func LogSetupAndDestruct() func() {
	logFilePath := services.ConfigService().GetConfig("systemLog:path", "/var/log/boilerplate.log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Panicln(err)
	}

	log.SetFormatter(&log.TextFormatter{DisableColors: true})
	log.SetOutput(logFile)
	log.SetLevel(log.DebugLevel)

	return func() {
		e := logFile.Close()
		if e != nil {
			fmt.Fprintf(os.Stderr, "Problem closing the log file: %s\n", e)
		}
	}
}

func LoadFileConfig() {
	services.ConfigService().LoadConfigFile("config.yaml")
}

func main() {
	LoadFileConfig()
	defer LogSetupAndDestruct()()
	log.Info("App start")

	staticFolder := services.ConfigService().GetConfig("staticFolder", "public")
	goji.Use(gojistatic.Static(staticFolder, gojistatic.StaticOptions{SkipLogging: false}))
	log.Debug("Static folder : " + staticFolder)

	port := services.ConfigService().GetConfig("port", ":8000")
	flag.Set("bind", port)

	app.App()

	goji.Serve()
}
