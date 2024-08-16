package filewriter

import (
	"os"
	"strconv"
)

var filepath string

func WriteIntToFile(value int) error {
	data := strconv.Itoa(value)
	err := os.WriteFile(filepath, []byte(data), 0644)
	return err
}

func ReadIntFromFile() (int, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}

	return value, nil
}

func Init(newpath string) {
	filepath = newpath
}
