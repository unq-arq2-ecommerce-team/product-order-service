package exception

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ProductNotFoundError(t *testing.T) {
	e1 := ProductNotFound{
		Id: 2,
	}
	e2 := ProductNotFound{
		Id: 3,
	}
	assert.Equal(t, `product with id 2 not found`, e1.Error())
	assert.Equal(t, `product with id 3 not found`, e2.Error())
}

func Test_ProductCannotDeleteError(t *testing.T) {
	e1 := ProductCannotDelete{
		Id: 2,
	}
	e2 := ProductCannotDelete{
		Id: 4,
	}
	assert.Equal(t, `product with id 2 cannot delete`, e1.Error())
	assert.Equal(t, `product with id 4 cannot delete`, e2.Error())
}

func Test_ProductCannotUpdateError(t *testing.T) {
	e1 := ProductCannotUpdate{
		Id: 2,
	}
	e2 := ProductCannotUpdate{
		Id: 4,
	}
	assert.Equal(t, `product with id 2 cannot update`, e1.Error())
	assert.Equal(t, `product with id 4 cannot update`, e2.Error())
}

func Test_ProductWithNoStockError(t *testing.T) {
	e1 := ProductWithNoStock{
		Id: 2,
	}
	e2 := ProductWithNoStock{
		Id: 4,
	}
	assert.Equal(t, `product with id 2 have no stock`, e1.Error())
	assert.Equal(t, `product with id 4 have no stock`, e2.Error())
}
