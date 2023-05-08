package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Address_String(t *testing.T) {
	address1 := Address{
		Street:      "Fake street 123",
		City:        "La Plata",
		State:       "Buenos Aires",
		Country:     "Argentina",
		Observation: "",
	}
	address2 := Address{
		Street:      "Fake street 321",
		City:        "Miami",
		State:       "Florida",
		Country:     "USA",
		Observation: "puerta negra",
	}
	assert.Equal(t, `[Address]{"street":"Fake street 123","city":"La Plata","state":"Buenos Aires","country":"Argentina","observation":""}`, address1.String())
	assert.Equal(t, `[Address]{"street":"Fake street 321","city":"Miami","state":"Florida","country":"USA","observation":"puerta negra"}`, address2.String())
}
