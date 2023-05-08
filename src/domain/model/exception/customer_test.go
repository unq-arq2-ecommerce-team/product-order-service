package exception

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CustomerAlreadyExistError(t *testing.T) {
	e1 := CustomerAlreadyExist{
		Email: "e1",
	}
	e2 := CustomerAlreadyExist{
		Email: "e2",
	}
	assert.Equal(t, `customer with email e1 already exists`, e1.Error())
	assert.Equal(t, `customer with email e2 already exists`, e2.Error())
}

func Test_CustomerNotFoundError(t *testing.T) {
	eId := CustomerNotFound{
		Id: 1,
	}
	eEmail := CustomerNotFound{
		Email: "e1",
	}
	assert.Equal(t, `customer with id 1 not found`, eId.Error())
	assert.Equal(t, `customer with email e1 not found`, eEmail.Error())
}

func Test_CustomerCannotDeleteError(t *testing.T) {
	e1 := CustomerCannotDelete{
		Id: 1,
	}
	e2 := CustomerCannotDelete{
		Id: 2,
	}
	assert.Equal(t, `customer with id 1 cannot delete`, e1.Error())
	assert.Equal(t, `customer with id 2 cannot delete`, e2.Error())
}

func Test_CustomerCannotUpdateError(t *testing.T) {
	e1 := CustomerCannotUpdate{
		Id: 1,
	}
	e2 := CustomerCannotUpdate{
		Id: 2,
	}
	assert.Equal(t, `customer with id 1 cannot update`, e1.Error())
	assert.Equal(t, `customer with id 2 cannot update`, e2.Error())
}
