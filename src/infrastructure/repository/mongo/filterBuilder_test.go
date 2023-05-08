package mongo

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func Test_GivenNewFilterBuilder_WhenBuild_ReturnAnEmptyBson(t *testing.T) {
	filterBuilder := NewFilterBuilder()
	assert.Equal(t, bson.M{}, filterBuilder.Build())
}

func Test_GivenNewFilterBuilder_WhenAppendOperationFilterWithABson_ReturnThisBson(t *testing.T) {
	filterBuilder := NewFilterBuilder()

	op := "op"
	filterBuilder.AppendOperationFilter(op, bson.M{"a": "1"})

	assert.Equal(t, bson.M{"a": "1"}, filterBuilder.Build())
}

func Test_GivenAFilterBuilderWithBsonX_WhenAppendOperationFilterWithABsonY_ReturnABsonNestedWithThatOperationAndBsonWithXAndY(t *testing.T) {
	filterBuilder := NewFilterBuilder()
	filterBuilder.AppendOperationFilter("", bson.M{"a": "1"})

	op := "op"
	filterBuilder.AppendOperationFilter(op, bson.M{"b": "2"})

	assert.Equal(t, bson.M{op: bson.A{bson.M{"a": "1"}, bson.M{"b": "2"}}}, filterBuilder.Build())
}

func Test_GivenAFilterBuilderWithBsonNestedWithOp1WithBsonXAndY_WhenAppendAndOpFilterWithBsonZ_ReturnABsonNestedWithThatAndOpAndWithBsonZAndBsonWithXAndY(t *testing.T) {
	op1 := "op1"
	filterBuilder := NewFilterBuilder()
	filterBuilder.AppendOperationFilter("", bson.M{op1: bson.A{bson.M{"a": "1"}, bson.M{"b": "2"}}})

	filterBuilder.AppendAndOpFilter(bson.M{"c": "3"})
	assert.Equal(t, bson.M{"$and": bson.A{bson.M{op1: bson.A{bson.M{"a": "1"}, bson.M{"b": "2"}}}, bson.M{"c": "3"}}}, filterBuilder.Build())
}

func Test_GivenAFilterBuilderWithBsonX_WhenAppendAndOpFilterIfWithABsonYAndFalseCondition_ReturnSameInitialBson(t *testing.T) {
	filterBuilder := NewFilterBuilder()
	filterBuilder.AppendOperationFilter("", bson.M{"a": "1"})

	filterBuilder.AppendAndOpFilterIf(false, bson.M{"b": "2"})

	assert.Equal(t, bson.M{"a": "1"}, filterBuilder.Build())
}

func Test_GivenAFilterBuilderWithBsonX_WhenAppendAndOpFilterIfWithABsonYAndTrueCondition_ReturnABsonNestedWithAndOpAndBsonWithXAndY(t *testing.T) {
	filterBuilder := NewFilterBuilder()
	filterBuilder.AppendOperationFilter("", bson.M{"a": "1"})

	filterBuilder.AppendAndOpFilterIf(true, bson.M{"b": "2"})

	assert.Equal(t, bson.M{"$and": bson.A{bson.M{"a": "1"}, bson.M{"b": "2"}}}, filterBuilder.Build())
}
