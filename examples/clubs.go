package main

import (
	"fmt"
	"github.com/m4kyu/kumite/pkg/api"
)

func main() {
	clubs, err := api.ChampClubs("champ_ua_06_25")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}


	for _, info := range *clubs {
		fmt.Printf("ID: %v. Name: %v. Count: %v. Logo: %v\n", info.ID, info.Name, info.Count, info.Logo)
	}
}


