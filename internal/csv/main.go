package csv

import (
	"encoding/json"
	"fmt"
	"slices"
)

type CSVData struct {
	records    map[string][]interface{}
	headers    []string
	rawRecords []map[string]interface{}
}

func (data *CSVData) JSON() ([]byte, error) {
	if data.records == nil {
		data.Sort()
	}

	jsonData, err := json.Marshal(data.records)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (data *CSVData) getKey(keys ...string) (string, error) {
	var key string
	if len(keys) == 0 {
		key = data.headers[0]
	} else {
		key = keys[0]
		if !slices.Contains(data.headers, key) {
			return "", fmt.Errorf("key %s not found in headers", key)
		}
	}
	return key, nil
}
