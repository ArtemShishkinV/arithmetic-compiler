package config

import (
	"errors"
	"strings"
)

type Mode string

const (
	Lexical Mode = "lex"
	Syntax  Mode = "syn"
	Unknown Mode = "unknown"
)

func GetMode(mode string) (Mode, error) {
	mode = strings.ToLower(mode)
	if mode == string(Lexical) {
		return Lexical, nil
	}

	if mode == string(Syntax) {
		return Syntax, nil
	}

	return Unknown, errors.New("unknown operating mode")
}
