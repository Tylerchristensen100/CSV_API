package csv

import "fmt"

func (data *CSVData) Sort(keys ...string) error {
	key, err := data.getKey(keys...)
	if err != nil {
		return err
	}

	var sortedData = make(map[string][]interface{}, 1)
	for _, record := range data.rawRecords {

		value, exists := record[key]
		if !exists || value == nil {
			return fmt.Errorf("key %q not found or nil in record", key)
		}

		strVal, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string for key %q, got %T", key, value)
		}

		sortedData[strVal] = append(sortedData[strVal], record)
	}

	data.records = sortedData
	return nil
}
