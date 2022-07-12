package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func getExecise() ([]Execise, error) {
	jsonFile, err := os.Open("exerciseDB.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	var execises []Execise
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(byteValue, &execises)
	return execises, nil
}

func findById(id string, execises []Execise) Execise {
	for _, val := range execises {
		if val.Id == id {
			return val
		}
	}
	return Execise{}
}
