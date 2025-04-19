package converter

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"
)

// ConvertCSVToJSONLines reads a CSV file and converts it to JSON Lines format.
// Each line in the output file represents a JSON object corresponding to a row in the CSV file.
// The first row of the CSV file is treated as the header, and each column is converted to a key-value pair in the JSON object.
func ConvertCSVToJSONLines(inputFile, outputFile string) error {
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	reader := csv.NewReader(inFile)
	headers, err := reader.Read()
	headers = sanitizeHeaders(headers)
	if err != nil {
		return err
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		obj := make(map[string]interface{})
		for i, key := range headers {
			if val, err := strconv.ParseFloat(record[i], 64); err == nil {
				obj[key] = val
			} else {
				obj[key] = record[i]
			}
		}

		line, err := json.Marshal(obj)
		if err != nil {
			return err
		}

		outFile.Write(line)
		outFile.Write([]byte("\n"))
	}

	return nil
}

// SanitizeHeaders removes the BOM from the first header if present.
func sanitizeHeaders(headers []string) []string {
	if len(headers) > 0 {
		headers[0] = strings.TrimPrefix(headers[0], "\uFEFF")
	}
	return headers
}
