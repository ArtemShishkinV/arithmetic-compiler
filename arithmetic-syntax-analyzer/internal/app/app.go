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
	if results == nil {
		return errors.New("implement syntax-analyzer")
	}
	if err = pkg.WriteFile(results[0], a.cfg.Files[1]); err != nil {
		return err
	}
	if err = pkg.WriteFile(results[1], a.cfg.Files[2]); err != nil {
		return err
	}
	return nil
}
