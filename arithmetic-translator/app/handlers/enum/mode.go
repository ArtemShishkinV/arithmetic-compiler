package enum

import "errors"

type Mode string

const (
	Translation Mode = "T"
	Generation  Mode = "G"
	unknown     Mode = "unknown"
)

func GetMode(mode string) (Mode, error) {
	if mode == "T" {
		return Translation, nil
	}

	if mode == "G" {
		return Generation, nil
	}

	return unknown, errors.New("unknown operating mode")
}
