package main

import (
	"fmt"
	"encoding/json"
	"os"
)

var settings struct {
	ServerMode bool `json:"serverMode"`
	SourceDir  string `json:"sourceDir"`
	TargetDir  string `json:"targetDir"`
	Array     []string `json: "array"`
}

func main() {

	// then config file settings

	configFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&settings); err != nil {
		panic(err)
	}

	fmt.Printf("%v %s %s", settings.ServerMode, settings.SourceDir, settings.TargetDir)
	fmt.Println(settings.Array)





	return
}