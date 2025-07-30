

package main

import (
	"fmt"
	"github.com/m4kyu/kumite/pkg/api"
)

func main() {
	count, err := api.ParticipantsCount("champ_ua_06_25")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	fmt.Printf("Participants count: %v\n", count)
}


