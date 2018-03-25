package gopattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonFromReader(t *testing.T) {
	result := jsonFromReader()
	assert.Equal(t, uint64(1), result.ID)
	assert.Equal(t, "findcoo", result.Name)
}

func TestJsonFromVar(t *testing.T) {
	result := jsonFromVar()
	assert.Equal(t, uint64(2), result.ID)
	assert.Equal(t, "findcoo", result.Name)
}

func TestWriterToJSON(t *testing.T) {
	result := writerToJSON()
	assert.Equal(t, []byte(`{"id":1,"name":"findcoo"}`), result[:len(result)-1])
}

func TestStructToJSON(t *testing.T) {
	result := structToJSON()
	assert.Equal(t, []byte(`{"id":1,"name":"findcoo"}`), result)
}
