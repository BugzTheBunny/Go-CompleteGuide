package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("Failed to find file.")
	}

	textValue := string(data)
	value, err := strconv.ParseFloat(textValue, 64)

	if err != nil {
		return defaultValue, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func WriteFloatToFile(filePath string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filePath, []byte(valueText), 0644)
}
