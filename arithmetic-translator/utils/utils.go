package utils

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"reflect"
	"strings"
)

func RandMapKey(m interface{}) interface{} {
	mapKeys := reflect.ValueOf(m).MapKeys()
	return mapKeys[rand.Intn(len(mapKeys))].Interface()
}

func WriteFile(strings []string, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, item := range strings {
		if _, err = file.WriteString(item + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func ReadFileLines(fileName string) ([]string, error) {
	f, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	var lines []string

	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		line = strings.TrimRight(line, "\n")
		lines = append(lines, line)
	}

	return lines, nil
}
