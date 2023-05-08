package mongo

import "go.mongodb.org/mongo-driver/bson"

const (
	mongoAndOp = "$and"
	mongoGTEOp = "$gte"
	mongoLTEOp = "$lte"
)

type FilterBuilder struct {
	filter bson.M
}

func NewFilterBuilder() *FilterBuilder {
	return &FilterBuilder{filter: nil}
}

func (b *FilterBuilder) AppendAndOpFilter(filter bson.M) *FilterBuilder {
	return b.AppendOperationFilter(mongoAndOp, filter)
}

func (b *FilterBuilder) AppendAndOpFilterIf(condition bool, filter bson.M) *FilterBuilder {
	if condition {
		return b.AppendAndOpFilter(filter)
	}
	return b
}

func (b *FilterBuilder) AppendOperationFilter(op string, filter bson.M) *FilterBuilder {
	if b.filter == nil {
		b.filter = filter
	} else if filter != nil {
		b.filter = bson.M{op: bson.A{b.filter, filter}}
	}
	return b
}

func (b *FilterBuilder) Build() bson.M {
	if b.filter == nil {
		return bson.M{}
	}
	return b.filter
}
