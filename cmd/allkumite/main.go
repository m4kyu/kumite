package main

import (
	"fmt"
	"github.com/m4kyu/kumite/pkg/api"
	"github.com/m4kyu/kumite/internal/db"
)

func main() {
	err := db.DB("./data/test.db")
	if err != nil {
		fmt.Println(err)
		return
	}


	champs, err := api.Champs()
	if err != nil {
		fmt.Println(err)
		return
	}

	db.AddChamps("./data/test.db", champs)
}


