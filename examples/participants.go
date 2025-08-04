package main

import (
	"fmt"
	"github.com/m4kyu/kumite/pkg/api"
)

func main() {
	fighters, err := api.Participants("champ_ua_06_25")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}


	for _, val := range *fighters {
		fmt.Printf("Name: %v. Age: %v. Categories: %v\n", val.FIO, val.Age, val.Categories)
	}
}


