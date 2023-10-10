package main

import (
	"fmt"

	"github.com/italorafaeltavares/gopportunities.git/config"
	"github.com/italorafaeltavares/gopportunities.git/router"
)

func main() {

	// Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize Router
	router.Initialize()
}
