package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/brookesargent/csv-normalizer/normalizations"
)

type DataRow struct {
	Timestamp     string
	Address       string
	Zip           string
	FullName      string
	FooDuration   float64
	BarDuration   float64
	TotalDuration float64
	Notes         string
}

func readCSV() [][]string {
	records, err := csv.NewReader(os.Stdin).ReadAll()
	if err != nil {
		fmt.Println("error")
	}
	return records
}

func writeCSV(rows []DataRow) {
	csvWriter := csv.NewWriter(os.Stdout)

	data := make([][]string, len(rows))
	for _, r := range rows {
		row := []string{r.Timestamp, r.Address, r.Zip, r.FullName, fmt.Sprintf("%f", r.FooDuration), fmt.Sprintf("%f", r.BarDuration), fmt.Sprintf("%f", r.TotalDuration), r.Notes}
		data = append(data, row)
	}
	err := csvWriter.WriteAll(data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	records := readCSV()

	rows := make([]DataRow, len(records))
	for i, record := range records {
		// don't normalize header
		if i == 0 {
			continue
		}

		formattedTimestamp, err := normalizations.FormatTimestamp(record[0])
		if err != nil {
			log.Println(fmt.Sprintf("Warning: row %d has an invalid timestamp that could not be parsed. It will be dropped from the output", i+1))
			continue
		}

		data := DataRow{
			Timestamp:   formattedTimestamp,
			Address:     record[1],
			Zip:         normalizations.FormatZipCode(record[2]),
			FullName:    record[3],
			FooDuration: normalizations.DurationToSeconds(record[4]),
			BarDuration: normalizations.DurationToSeconds(record[5]),
			Notes:       record[7],
		}
		data.TotalDuration = data.FooDuration + data.BarDuration
		rows = append(rows, data)
	}

	writeCSV(rows)
}
