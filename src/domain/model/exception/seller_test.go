package exception

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SellerNotFoundError(t *testing.T) {
	eId := SellerNotFound{
		Id: 1,
	}
	e2Id := SellerNotFound{
		Id: 65,
	}
	assert.Equal(t, `seller with id 1 not found`, eId.Error())
	assert.Equal(t, `seller with id 65 not found`, e2Id.Error())
}
