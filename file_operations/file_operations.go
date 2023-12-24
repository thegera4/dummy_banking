package file_operations

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

// NOTE: to export this functions and make them available to other files, you have to
//name them, starting ith a capital letter. It is the "export" keyword equivalente in GO.

func WriteFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("Error while reading the file...")
	}

	readText := string(data)
	parsedValue, err := strconv.ParseFloat(readText, 64)

	if err != nil {
		return defaultValue, errors.New("Error while parsing the value...")	
	}

	return parsedValue, nil
}