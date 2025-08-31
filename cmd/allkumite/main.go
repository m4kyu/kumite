package main

import (
	"fmt"
	"github.com/m4kyu/kumite/pkg/api"
)

func main() {
	data, err := api.ChampExt("champ_ua_06_25")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}


	fmt.Println(data.InfoUA)
	fmt.Println(data.RatingChamp)
	//for _, val := range *data {
//		fmt.Printf("ID: %v. URL: %v\n", val.TatamiID, val.URL)
	//}
}


