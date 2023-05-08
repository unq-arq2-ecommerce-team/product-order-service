package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ParseStruct_WithNil_Should_ReturnNullString(t *testing.T) {
	assert.Equal(t, `[]null`, ParseStruct("", nil))
}

func Test_ParseStruct_WithSimpleAnomStruct_Should_Ok(t *testing.T) {
	nameStruct := "AnomStruct"
	anomStruct := struct {
		SomeInt   int
		SomeStr   string
		NullField *string
	}{
		SomeInt: 321,
		SomeStr: "str",
	}
	assert.Equal(t, `[AnomStruct]{"SomeInt":321,"SomeStr":"str","NullField":null}`, ParseStruct(nameStruct, anomStruct))
}

func Test_ParseStruct_WithCollectionsOrNested_Should_ReturnOk(t *testing.T) {
	nameStruct := "ComplexStruct"
	anomStruct := struct {
		SomeArr    []string
		SomeMap    map[string]string
		SomeStruct struct{ SomeStr string }
	}{
		SomeArr:    []string{"a", "b"},
		SomeMap:    map[string]string{"a": "1", "b": "2", "c": "3"},
		SomeStruct: struct{ SomeStr string }{SomeStr: "sarasa"},
	}
	assert.Equal(t, `[ComplexStruct]{"SomeArr":["a","b"],"SomeMap":{"a":"1","b":"2","c":"3"},"SomeStruct":{"SomeStr":"sarasa"}}`, ParseStruct(nameStruct, anomStruct))
}

func Test_ParseStruct_WithNoCommonEncondingJSON_Should_ReturnEmptyString(t *testing.T) {
	nameFn := "Func_IdString"
	idString := func(x string) string { return x }
	assert.Equal(t, "", ParseStruct(nameFn, idString))
}
