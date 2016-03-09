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

func (this *configService) GetConfig(key string, defaultValue string) string {
	value, err := config.GetString(key)
	if err != nil {
		return defaultValue
	}
	return value
}
