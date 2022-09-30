package app

import (
	"arithmetic-syntax-analyzer/internal/config"
	"arithmetic-syntax-analyzer/internal/handlers"
	"arithmetic-syntax-analyzer/pkg"
	"errors"
)

type App struct {
	cfg *config.Config
}

func NewApp(config *config.Config) (App, error) {
	return App{
		cfg: config,
	}, nil
}

func (a *App) Run() error {
	expressions, err := pkg.ReadFileLines(a.cfg.SrcFileName)
	if err != nil {
		return err
	}
	if len(expressions) != 1 {
		return errors.New("source file must contain only one string")
	}

	analyzer := handlers.NewLexicalAnalyzer()
	tokens, tableVars, err := analyzer.Start(expressions[0])
	if err != nil {
		return err
	}
	if err = pkg.WriteFile(tokens, a.cfg.OutTokensFileName); err != nil {
		return err
	}
	if err = pkg.WriteFile(tableVars, a.cfg.OutSymbolsFileName); err != nil {
		return err
	}
	return nil
}
