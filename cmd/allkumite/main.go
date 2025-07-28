

package main

import (
	"fmt"
	"github.com/m4kyu/kumite/pkg/api"
)

func main() {
	orgs, err := api.Champs()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	for _, i := range *orgs {
		fmt.Printf("Name: %v. ID: %v\n", i.Tittle, i.ID)
	}
}


