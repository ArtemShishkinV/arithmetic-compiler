package main

import (
	"arithmetic-translator/app/handlers"
	"arithmetic-translator/app/handlers/config"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	newConfig, err := config.NewConfig(args)

	if err != nil {
		fmt.Println(err)
		return
	}

	handler, err := handlers.NewHandler(newConfig)

	if err != nil {
		fmt.Println(err)
		return
	}

	if err = handler.Start(); err != nil {
		fmt.Println(err)
		return
	}
}
