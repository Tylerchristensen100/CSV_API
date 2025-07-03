package csv

func (data *CSVData) Filter(key string, value interface{}) {
	var filteredData []map[string]interface{}
	if key != "" && value != nil {
		for _, record := range data.rawRecords {
			if recordValue, exists := record[key]; exists && recordValue == value {
				filteredData = append(filteredData, record)
			}
		}
		data.rawRecords = filteredData
	}
}
