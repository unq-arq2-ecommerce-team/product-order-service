package exception

import "fmt"

type ProductNotFound struct {
	Id int64
}

func (e ProductNotFound) Error() string {
	return fmt.Sprintf("product with id %v not found", e.Id)
}

type ProductCannotDelete struct {
	Id int64
}

func (e ProductCannotDelete) Error() string {
	return fmt.Sprintf("product with id %v cannot delete", e.Id)
}

type ProductCannotUpdate struct {
	Id int64
}

func (e ProductCannotUpdate) Error() string {
	return fmt.Sprintf("product with id %v cannot update", e.Id)
}

type ProductWithNoStock struct {
	Id int64
}

func (e ProductWithNoStock) Error() string {
	return fmt.Sprintf("product with id %v have no stock", e.Id)
}
