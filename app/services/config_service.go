package services

import (
	"fmt"

	"github.com/tsuru/config"
)

type configService struct {
}

func ConfigService() *configService {
	return &configService{}
}

func (this *configService) LoadConfigFile(path string) {
	defer func() {
		if i := recover(); i != nil {
			fmt.Println("Error loading file config : ", i)
		}
	}()
	err := config.ReadConfigFile(path)
	if err != nil {
		panic(err)
	}
}

func (this *configService) GetServerPort() string {
	value, err := config.GetString("port")
	if err == nil {
		return value
	}
	return ":8000"
}

func (this *configService) GetStaticFolder() string {
	value, err := config.GetString("staticFolder")
	if err == nil {
		return value
	}
	return "public"
}
