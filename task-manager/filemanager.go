package main

import (
	"encoding/json"
	"os"
)

/*
1. read file and return task array
2. update file with new data
*/

func extract_data(filename string) ([]Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var data []Task
	var d Task
	decoder.Token()
	for decoder.More() {
		err = decoder.Decode(&d)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}

func updata_data(filename string, t []Task) error {
	wfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer wfile.Close()
	jsonEncoder := json.NewEncoder(wfile)
	if err = jsonEncoder.Encode(t); err != nil {
		return err
	}
	return nil
}
