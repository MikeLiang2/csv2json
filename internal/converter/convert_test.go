package converter

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// readJSONLines reads a .jl (JSON Lines) file and parses each line into a map[string]interface{}.
// It returns a slice of maps, where each map represents one JSON object.
func readJSONLines(path string) ([]map[string]interface{}, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var lines []map[string]interface{}
	for _, line := range strings.Split(string(content), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(line), &obj)
		if err != nil {
			return nil, err
		}
		lines = append(lines, obj)
	}

	return lines, nil
}

// mapsEqual compares two maps for deep equality.
// It assumes values are primitive types (numbers, strings).
func mapsEqual(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

// TestConvertCSVToJSONLines test for ConvertCSVToJSONLines.
// It verifies that the output .jl file is structurally equivalent to the expected JSON Lines file.
func TestConvertCSVToJSONLines(t *testing.T) {
	input := "../../testdata/sample.csv"
	output := "../../testdata/test_output.jl"
	expected := "../../testdata/expected.jl"

	os.Remove(output)
	err := ConvertCSVToJSONLines(input, output)
	if err != nil {
		t.Fatalf("Conversion failed: %v", err)
	}

	expectedData, err := readJSONLines(expected)
	if err != nil {
		t.Fatalf("Failed to read expected file: %v", err)
	}

	actualData, err := readJSONLines(output)
	if err != nil {
		t.Fatalf("Failed to read actual output: %v", err)
	}

	if len(expectedData) != len(actualData) {
		t.Fatalf("Line count mismatch: expected %d, got %d", len(expectedData), len(actualData))
	}

	for i := range expectedData {
		if !mapsEqual(expectedData[i], actualData[i]) {
			t.Errorf("Mismatch at line %d:\nExpected: %+v\nGot:      %+v\n", i+1, expectedData[i], actualData[i])
		}
	}
}
