package main

import (
	"fmt"
	"github.com/disism/oauth2/oauth2"
)

func main() {
	if err := oauth2.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
