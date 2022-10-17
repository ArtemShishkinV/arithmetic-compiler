package main

import (
	"arithmetic-compiler/internal/app"
	"arithmetic-compiler/internal/config"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	//args := os.Args[1:]
	args := []string{"sem", ".\\files\\source.txt", ".\\files\\tree.txt"}
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
