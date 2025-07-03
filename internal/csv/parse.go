package csv

import (
	csv_reader "encoding/csv"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

func ParseFromUrl(url string) (*CSVData, error) {
	if !strings.Contains(url, "://") {
		url = "https://" + url
	}
	curl, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(curl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch CSV from URL: " + url)
	}

	csvData, err := parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return csvData, nil
}
func Parse(path string) (*CSVData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	csvData, err := parse(file)
	if err != nil {
		return nil, err
	}
	return csvData, nil
}

func parse(r io.Reader) (*CSVData, error) {
	reader := csv_reader.NewReader(r)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var csvData = &CSVData{}

	data := []map[string]interface{}{}
	csvData.headers = records[0]
	rows := records[1:]

	for _, row := range rows {
		if len(row) != len(csvData.headers) {
			return nil, csv_reader.ErrFieldCount
		}
		record := make(map[string]interface{})
		for i, header := range csvData.headers {
			record[header] = row[i]

			if record[header] == "" {
				record[header] = nil
			}
		}
		data = append(data, record)
	}

	csvData.rawRecords = data

	return csvData, nil
}
