package config

import (
	"errors"
	"strings"
)

type Mode string

const (
	Lexical    Mode = "lex"
	Syntax     Mode = "syn"
	Semantic   Mode = "sem"
	Generator1 Mode = "gen1"
	Generator2 Mode = "gen2"
	Unknown    Mode = "unknown"
)

func GetMode(mode string) (Mode, error) {
	mode = strings.ToLower(mode)
	if mode == string(Lexical) {
		return Lexical, nil
	}

	if mode == string(Syntax) {
		return Syntax, nil
	}

	if mode == string(Semantic) {
		return Semantic, nil
	}

	if mode == string(Generator1) {
		return Generator1, nil
	}

	if mode == string(Generator2) {
		return Generator2, nil
	}

	return Unknown, errors.New("unknown operating mode")
}
