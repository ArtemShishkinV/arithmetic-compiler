package main

import (
	"arithmetic-lexical-analyzer/internal/config"
	"arithmetic-lexical-analyzer/internal/handlers"
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

	handler := handlers.NewHandler(defConfig)
	if err = handler.Start(); err != nil {
		fmt.Println(err)
		return
	}
}
