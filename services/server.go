package seppo

import (
	"github.com/koodinikkarit/seppo/db"
)

func CreateSeppoServer() {
	databaseService := SeppoDB.CreateDb("jaska", "asdf321", "localhost", "3306", "seppo")
	go databaseService.Start()
	CreateSeppoService("3214", databaseService)
}
