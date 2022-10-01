package pkg

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
		line = strings.TrimRight(line, "\r\n")
		lines = append(lines, line)
	}

	return lines, nil
}
