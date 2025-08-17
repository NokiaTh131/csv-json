package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"io"
	"log/slog"
	"os"
	"regexp"
)

func main() {
	inputFile := flag.String("f", "", "Input file")
	outputFile := flag.String("o", "output.json", "Output file")
	csvRx := regexp.MustCompile(`^.*\.csv$`)
	jsonRx := regexp.MustCompile(`^.*\.json$`)

	flag.Parse()

	if *inputFile == "" {
		slog.Error("Input file not specified")
		return
	}

	if !csvRx.MatchString(*inputFile) {
		slog.Error("Invalid input file extension", "ext", *inputFile)
		return
	}

	if !jsonRx.MatchString(*outputFile) {
		slog.Error("Invalid output file extension", "ext", *outputFile)
		return
	}

	file, err := os.Open(*inputFile)
	if err != nil {
		slog.Error("Error opening file", "err", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		slog.Error("Error reading CSV headers", "err", err)
		return
	}

	var data []map[string]string

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			slog.Error("Error reading CSV row", "err", err)
			return
		}
		rowMap := make(map[string]string)
		for i, header := range headers {
			if i < len(row) {
				rowMap[header] = row[i]
			} else {
				rowMap[header] = ""
			}
		}
		data = append(data, rowMap)
	}
	json, err := json.Marshal(data)
	if err != nil {
		slog.Error("Error marshaling JSON", "err", err)
		return
	}

	file, err = os.Create(*outputFile)
	if err != nil {
		slog.Error("Error creating output file", "err", err)
		return
	}
	defer file.Close()

	_, err = file.Write(json)
	if err != nil {
		slog.Error("Error writing to output file", "err", err)
		return
	}
	slog.Info("Successfully wrote JSON to output file", "file", *outputFile)
}
