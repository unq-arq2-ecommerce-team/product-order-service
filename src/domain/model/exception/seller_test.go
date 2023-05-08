package exception

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SellerAlreadyExistError(t *testing.T) {
	e1 := SellerAlreadyExist{
		Name: "e1",
	}
	e2 := SellerAlreadyExist{
		Name: "e2",
	}
	assert.Equal(t, `seller with name e1 already exists`, e1.Error())
	assert.Equal(t, `seller with name e2 already exists`, e2.Error())
}

func Test_SellerNotFoundError(t *testing.T) {
	eId := SellerNotFound{
		Id: 1,
	}
	eName := SellerNotFound{
		Name: "e1",
	}
	assert.Equal(t, `seller with id 1 not found`, eId.Error())
	assert.Equal(t, `seller with name e1 not found`, eName.Error())
}

func Test_SellerCannotDeleteError(t *testing.T) {
	e1 := SellerCannotDelete{
		Id: 1,
	}
	e2 := SellerCannotDelete{
		Id: 2,
	}
	assert.Equal(t, `seller with id 1 cannot delete`, e1.Error())
	assert.Equal(t, `seller with id 2 cannot delete`, e2.Error())
}

func Test_SellerCannotUpdateError(t *testing.T) {
	e1 := SellerCannotUpdate{
		Id: 1,
	}
	e2 := SellerCannotUpdate{
		Id: 2,
	}
	assert.Equal(t, `seller with id 1 cannot update`, e1.Error())
	assert.Equal(t, `seller with id 2 cannot update`, e2.Error())
}
