package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/koodinikkarit/teppo/seppo_service"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("192.168.180.29:3214", grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("Grpc connection failed")
	}
	defer conn.Close()

	c := SeppoService.NewSeppoClient(conn)

	files, _ := ioutil.ReadDir("./lauludata2")

	numberOfSongs := len(files) - 1

	fmt.Println(numberOfSongs)

	for index, f := range files {
		if strings.Contains(f.Name(), ".txt") == true {
			splits := strings.Split(f.Name(), ".txt")
			name := splits[0]
			dat, _ := ioutil.ReadFile("./lauludata2/" + f.Name())
			text := string(dat)
			fmt.Println("Lähetetään laulu " + name + " " + strconv.Itoa(index) + "/" + strconv.Itoa(numberOfSongs))
			_, err := c.CreateVariation(context.Background(), &SeppoService.CreateVariationRequest{
				Name:            name,
				Text:            text,
				Origin:          "jyvaskyla",
				SongDatabaseIds: []uint32{1},
				TagIds:          []uint32{2},
			})
			if err != nil {
				fmt.Println("Laulun " + name + " lahetys epäonnistui " + err.Error())
			} else {
				fmt.Println("Laulu " + name + " lähetetty")
			}
		}
	}

}
