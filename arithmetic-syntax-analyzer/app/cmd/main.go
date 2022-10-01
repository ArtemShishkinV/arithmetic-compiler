package main

import (
	"arithmetic-syntax-analyzer/internal/app"
	"arithmetic-syntax-analyzer/internal/config"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	defConfig, err := config.NewConfig(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	application, _ := app.NewApp(defConfig)
	if err = application.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
