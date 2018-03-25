package gopattern

import (
	"bytes"
	"encoding/json"
)

// DataSet data set for test
type DataSet struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func jsonFromReader() *DataSet {
	data := &DataSet{}
	json.NewDecoder(bytes.NewReader([]byte(`{"id": 1, "name": "findcoo"}`))).Decode(data)
	return data
}

func jsonFromVar() *DataSet {
	data := &DataSet{}
	json.Unmarshal([]byte(`{"id": 2, "name": "findcoo"}`), data)
	return data
}

func writerToJSON() []byte {
	data := &DataSet{
		ID:   1,
		Name: "findcoo",
	}

	buff := &bytes.Buffer{}
	json.NewEncoder(buff).Encode(data)
	return buff.Bytes()
}

func structToJSON() []byte {
	data := &DataSet{
		ID:   1,
		Name: "findcoo",
	}

	jsonText, _ := json.Marshal(data)
	return jsonText
}
