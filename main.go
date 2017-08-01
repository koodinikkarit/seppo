package main

import (
	"github.com/koodinikkarit/seppo/services"
)

func main() {
	config := seppo.LoadConfig("config.yml")

	seppo.CreateSeppoServer(config)
	for {

	}
}
