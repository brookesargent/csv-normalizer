package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type DataRow struct {
	Timestamp     string
	Address       string
	Zip           string
	FullName      string
	FooDuration   string
	BarDuration   string
	TotalDuration string
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
	csvwriter := csv.NewWriter(os.Stdout)

	var data [][]string
	for _, r := range rows {
		row := []string{r.Timestamp, r.Address, r.Zip, r.FullName, r.FooDuration, r.BarDuration, r.TotalDuration, r.Notes}
		data = append(data, row)
	}
	err := csvwriter.WriteAll(data) // calls Flush internally
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	records := readCSV()

	var rows []DataRow
	for _, record := range records {
		data := DataRow{
			Timestamp:     record[0],
			Address:       record[1],
			Zip:           record[2],
			FullName:      record[3],
			FooDuration:   record[4],
			BarDuration:   record[5],
			TotalDuration: record[6],
			Notes:         record[7],
		}
		rows = append(rows, data)
	}

	writeCSV(rows)
}
