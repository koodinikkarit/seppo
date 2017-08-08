package main

import (
	"github.com/koodinikkarit/seppo/services"
)

func main() {
	config := seppo.GetConfig()

	seppo.CreateSeppoServer(config)
	for {

	}
}
