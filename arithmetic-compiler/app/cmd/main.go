package main

import (
	"arithmetic-compiler/internal/app"
	"arithmetic-compiler/internal/config"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
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
	fmt.Println(time.Since(start))
}
