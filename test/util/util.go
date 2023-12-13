package util

import (
	"encoding/json"
	"os"
	"testing"
)

// CreateJSONFile creates a file
func CreateJSONFile(t *testing.T, fName string, data interface{}) {
	var dataByte []byte
	var err error
	dataByte, err = json.MarshalIndent(data, "", "\t")

	if err != nil {
		t.Fatalf("couldnt marshal data to json. error [%s]", err)
		return
	}
	err = os.WriteFile(fName, dataByte, 0600)
	if err != nil {
		t.Fatalf("couldnt create file [%s] error [%s]", fName, err)
	}
}

// ReadJSONFile reads filename fName and unmarshall json to data
func ReadJSONFile[T any](t *testing.T, fName string) T {
	dataByte, err := os.ReadFile(fName) // #nosec
	var data T
	if err != nil {
		t.Fatalf("couldnt read file. error [%s]", err)
		return data
	}
	err = json.Unmarshal(dataByte, &data)
	if err != nil {
		t.Fatalf("couldnt unmarshal data to json. error [%s]", err)
	}
	return data
}
