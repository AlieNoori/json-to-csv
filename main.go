package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func readJSON(filepath string) ([]map[string]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("json file does not exist: %s", err)
	}
	defer f.Close()

	var result []map[string]string
	if err := json.NewDecoder(f).Decode(&result); err != nil {
		return nil, fmt.Errorf("error while decoding json: %s", err)
	}

	if len(result) <= 0 {
		return nil, fmt.Errorf("empty json file")
	}

	return result, nil
}

func writeCSV(data []map[string]string, filepath string) error {
	values := make([][]string, 0, len(data))
	keys := make([]string, 0, len(data[0]))

	for k := range data[0] {
		keys = append(keys, k)
	}

	for _, v := range data {
		value := []string{}
		for _, k := range keys {
			value = append(value, v[k])
		}
		values = append(values, value)
	}

	fn, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Errro creating csv file:%s", err)
	}
	defer fn.Close()

	w := csv.NewWriter(fn)

	if err = w.Write(keys); err != nil {
		return fmt.Errorf("error writing header: %s", err)
	}

	if err = w.WriteAll(values); err != nil {
		return fmt.Errorf("error writing records: %s", err)
	}

	defer w.Flush()
	return nil
}

func main() {
	filepath := os.Args[1]
	if ok := strings.Contains(filepath, ".json"); !ok {
		filepath += ".json"
	}

	data, err := readJSON(filepath)
	if err != nil {
		log.Fatal(err)
	}

	outputfile := strings.Split(filepath, ".")[0] + ".csv"

	if err = writeCSV(data, outputfile); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully converted!")
}
