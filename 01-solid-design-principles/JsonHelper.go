package main

import "encoding/json"

// JsonHelper is responsible for json manipulation operations
type JsonHelper struct {
	data interface{}
}

// ToJSON converts data to JSON format
// data is the struct / map / slice type variable
func (j *JsonHelper) ToJSON() (s []byte, err error) {
	js, err := json.Marshal(j.data)
	if err != nil {
		return nil, err
	}
	var dat interface{}
	if err = json.Unmarshal(js, &dat); err != nil {
		return nil, err
	}
	mapB, err := json.Marshal(dat)
	if err != nil {
		return nil, err
	}
	return mapB, nil
}
