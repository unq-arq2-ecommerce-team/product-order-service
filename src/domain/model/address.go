package model

import "github.com/cassa10/arq2-tp1/src/domain/util"

type Address struct {
	Street      string `json:"street" bson:"street" binding:"required"`
	City        string `json:"city" bson:"city" binding:"required"`
	State       string `json:"state" bson:"state" binding:"required"`
	Country     string `json:"country" bson:"country" binding:"required"`
	Observation string `json:"observation" bson:"observation"`
}

func (a Address) String() string {
	return util.ParseStruct("Address", a)
}
