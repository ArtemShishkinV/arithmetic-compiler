package main

import (
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
	fmt.Printf("#running\nConfig: %s, %s, %s",
		defConfig.SrcFileName, defConfig.OutTokensFileName, defConfig.OutSymbolsFileName)
}
