package utils

import (
	"io"
	"os"
)

var (
	Open    = os.Open
	ReadAll = io.ReadAll
)

func ReadFile(filepath string) ([]byte, error) {
	file, err := Open(filepath)

	if err != nil {
		return nil, err
	}

	bytes, err := ReadAll(file)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	return bytes, nil
}
