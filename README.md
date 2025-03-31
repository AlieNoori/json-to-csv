# JSON to CSV Converter

A simple command-line tool written in Go that converts JSON files to CSV format.

## Overview

This tool reads a JSON file containing an array of objects with string values and converts it to a CSV file with matching headers and data.

## Requirements

- Go 1.23.5 or later

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/json-to-csv.git
   cd json-to-csv
   ```

2. Build the executable:
   ```
   go build
   ```

## Usage

Run the program with the JSON filename as a command-line argument:

```
./json-to-csv data
```

Or with the full filename:

```
./json-to-csv data.json
```

The program will generate a CSV file with the same name but a `.csv` extension (e.g., `data.csv`).

## Input Format

The JSON file should contain an array of objects where all values are strings.

Example input (data.json):
```json
[
  {
    "name": "John",
    "age": "30",
    "city": "New York"
  },
  {
    "name": "Jane",
    "age": "25",
    "city": "San Francisco"
  }
]
```

## Output Format

The program will create a CSV file with:
- A header row containing all the keys from the JSON objects
- Data rows containing the corresponding values

Example output (data.csv):
```
name,age,city
John,30,New York
Jane,25,San Francisco
```

