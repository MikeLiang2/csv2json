# csv2json

A command-line tool written in Go to convert CSV to JSON Lines format.

# Features

- Converts standard comma-delimited CSV into JSON Lines format
- Preserves numeric and string values automatically
- Ignores trailing blank lines in CSV
- Simple CLI interface
- Includes unit tests and expected output validation
- Ready to be integrated into other Go projects

---

## Build Instructions

# Compile the tool
go build -o csv2json
Or using the compiled file "csv2json"

# Run the tool manually
./csv2json convert testdata/sample.csv testdata/output.jl

## Usage
csv2json convert [input.csv] [output.jl]

# Example:
csv2json convert testdata/sample.csv testdata/output.jl

## Testing
To run unit tests:
go test ./...

## Sample Input (CSV)
value,income,age,rooms,bedrooms,pop,hh
452600,8.3252,41,880,129,322,126
358500,8.3014,21,7099,1106,2401,1138

## Sample Output (JL)
{"value":452600,"income":8.3252,"age":41,"rooms":880,"bedrooms":129,"pop":322,"hh":126}
{"value":358500,"income":8.3014,"age":21,"rooms":7099,"bedrooms":1106,"pop":2401,"hh":1138}

## AI Assistant Disclosure
This project was developed with the support of ChatGPT.
Guidance include Go project structure, testing design, and documentation generation.
All core logic, coding decisions, and testing were understood, reviewed, and manually verified.

