package Controller

import (
	"main/Config"
	"main/Model"
)

func GetConfig() Model.Configuration {
	return Config.GetConfig()
}
