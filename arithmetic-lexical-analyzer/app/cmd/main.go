package main

import (
	"arithmetic-lexical-analyzer/internal/app"
	"arithmetic-lexical-analyzer/internal/config"
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
	fmt.Printf("#running\nConfig: %s, %s, %s\n",
		defConfig.SrcFileName, defConfig.OutTokensFileName, defConfig.OutSymbolsFileName)

	application, _ := app.NewApp(defConfig)
	if err = application.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
