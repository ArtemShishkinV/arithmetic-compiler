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
	expressions, err := pkg.ReadFileLines(a.cfg.Files[0])

	if err != nil {
		return err
	}
	if len(expressions) != 1 {
		return errors.New("source file must contain only one string")
	}

	analyzer := handlers.NewHandler(*a.cfg)
	results, err := analyzer.Start(expressions[0])

	if err != nil {
		return err
	}
	if err := a.outResultToFiles(results); err != nil {
		return err
	}
	return nil
}

func (a *App) outResultToFiles(results [][]string) error {
	if err := pkg.WriteFile(results[0], a.cfg.Files[1]); err != nil {
		return err
	}
	if a.cfg.Mode == config.Lexical {
		if err := pkg.WriteFile(results[1], a.cfg.Files[2]); err != nil {
			return err
		}
	}
	return nil
}
