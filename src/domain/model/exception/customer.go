package exception

import "fmt"

type CustomerAlreadyExist struct {
	Email string
}

func (e CustomerAlreadyExist) Error() string {
	return fmt.Sprintf("customer with email %s already exists", e.Email)
}

type CustomerNotFound struct {
	Id    int64
	Email string
}

func (e CustomerNotFound) Error() string {
	if e.Id != 0 {
		return fmt.Sprintf("customer with id %v not found", e.Id)
	}
	return fmt.Sprintf("customer with email %v not found", e.Email)
}

type CustomerCannotDelete struct {
	Id int64
}

func (e CustomerCannotDelete) Error() string {
	return fmt.Sprintf("customer with id %v cannot delete", e.Id)
}

type CustomerCannotUpdate struct {
	Id int64
}

func (e CustomerCannotUpdate) Error() string {
	return fmt.Sprintf("customer with id %v cannot update", e.Id)
}
