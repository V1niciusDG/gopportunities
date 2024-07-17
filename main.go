package main

import (
	"fmt"

	"github.com/V1niciusDG/gopportunities/config"
	"github.com/V1niciusDG/gopportunities/router"
)

func main() {

	// initialize configs
	err := config.Init()
	if err != nil{
		fmt.Println(err)
		return
	}

	//chama router
	router.Initialize()
}