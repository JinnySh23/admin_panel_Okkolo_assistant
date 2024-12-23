// ------------------------------------
// RR IT 2024
//
// ------------------------------------

//
// ----------------------------------------------------------------------------------
//
// 				CSV (The module for working with the CSV dataset)
//
// ----------------------------------------------------------------------------------

package rr_csv

import (
	// System Packages
	"encoding/csv"
	"fmt"
	"os"
)

// A function for reading data from a CSV file
func ReadCSVData(filePath string) ([][]string, error) {
	// Opening the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	// Reading data from a file
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return data, nil
}

// Function for adding a line to the end of a CSV file
func AppendToCSV(filePath string, question, answer string) error {
	// Opening the file in the add mode
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	// Creating a CSV recorder
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Writing a new line
	err = writer.Write([]string{question, answer})
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}

	return nil
}

// A function for deleting a line from a CSV file by index
func RemoveRowFromCSV(filePath string, rowIndex int) error {
	// Reading data from a file
	data, err := ReadCSVData(filePath)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	// We check that the row index is in the acceptable range.
	if rowIndex < 0 || rowIndex >= len(data) {
		return fmt.Errorf("индекс строки %d выходит за пределы файла", rowIndex)
	}

	// Deleting the row with the specified index
	data = append(data[:rowIndex], data[rowIndex+1:]...)

	// Overwriting the updated data file
	err = WriteCSVData(filePath, data)
	if err != nil {
		return fmt.Errorf("ошибка записи файла: %w", err)
	}

	return nil
}

// Function for writing data to a CSV file
func WriteCSVData(filePath string, data [][]string) error {
	// Opening the file for recording
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer file.Close()

	// Creating a CSV recorder
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Writing the data to a file
	for _, row := range data {
		err := writer.Write(row)
		if err != nil {
			return fmt.Errorf("ошибка записи строки в файл: %w", err)
		}
	}

	return nil
}
