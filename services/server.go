package seppo

import (
	"fmt"

	"github.com/koodinikkarit/seppo/db"
)

func work(createSongChannel chan seppo.CreateSongInput) {
	fmt.Println("work")
	for {
		fmt.Println("Luodaan laulu ", <-createSongChannel)
	}
}

func work2(createSongChannel chan seppo.CreateSongInput) {
	createSongChannel <- seppo.CreateSongInput{
		Name: "Laulu2",
	}
}

func CreateSeppoServer() {
	databaseService := seppo.CreateDb("jaska", "asdf321", "localhost", "3306", "seppo")
	go databaseService.Start()

	createSongChannel := make(chan seppo.CreateSongInput)
	go CreateSeppoService("3214", databaseService)

	go work(createSongChannel)

	go work2(createSongChannel)

	fmt.Println("Luotu sepposerver")
}
