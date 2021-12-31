package Config

import (
	"encoding/json"
	"fmt"
	"main/Shared/Model"
	"os"
)

func GetConfig() Model.Configuration {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Model.Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
