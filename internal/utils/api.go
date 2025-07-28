package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

  "bytes"
)


func FetchAPI[T any](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	
	return marshalData[T](resp)
}


func PostAPI[T any](url string, data map[string]string) (*T, error) {
	jsonData, err := json.Marshal(data) 
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err;
	}	

	return marshalData[T](resp)
}


func ToStruct[T any](data *map[string]T) (*[]T) {
	elements := make([]T, len(*data))	
	for i, val := range *data {
		index, _ := strconv.Atoi(i)
		elements[index-1] = val
	}

	return &elements
}


func marshalData[T any](resp *http.Response) (*T, error) {
	defer resp.Body.Close()

	
	data2, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	
	var res T
	err = json.Unmarshal(data2, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}


