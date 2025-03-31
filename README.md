# JSON to CSV Converter

A simple command-line tool written in Go that converts JSON files to CSV format.

## Overview

This tool reads a JSON file containing an array of objects and converts it to a CSV file with matching headers and data. It handles various data types (strings, numbers, booleans) and automatically converts them to appropriate string representations in the CSV output.

## Features

- Converts JSON arrays of objects to CSV format
- Preserves field order in the output CSV
- Simple command-line interface

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

The JSON file should contain an array of objects. The program supports various data types and will convert them to appropriate string representations.

Example input (data.json):
```json
[
  {
    "name": "John",
    "age": 30,
    "city": "New York",
    "active": true
  },
  {
    "name": "Jane",
    "age": 25,
    "city": "San Francisco",
    "active": false
  }
]
```

## Output Format

The program will create a CSV file with:
- A header row containing all the keys from the JSON objects
- Data rows containing the corresponding values converted to strings

Example output (data.csv):
```
name,age,city,active
John,30,New York,true
Jane,25,San Francisco,false
```
