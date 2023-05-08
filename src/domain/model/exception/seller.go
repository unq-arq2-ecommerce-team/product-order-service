package exception

import "fmt"

type SellerAlreadyExist struct {
	Name string
}

func (e SellerAlreadyExist) Error() string {
	return fmt.Sprintf("seller with name %s already exists", e.Name)
}

type SellerNotFound struct {
	Id   int64
	Name string
}

func (e SellerNotFound) Error() string {
	if e.Id != 0 {
		return fmt.Sprintf("seller with id %v not found", e.Id)
	}
	return fmt.Sprintf("seller with name %v not found", e.Name)
}

type SellerCannotDelete struct {
	Id int64
}

func (e SellerCannotDelete) Error() string {
	return fmt.Sprintf("seller with id %v cannot delete", e.Id)
}

type SellerCannotUpdate struct {
	Id int64
}

func (e SellerCannotUpdate) Error() string {
	return fmt.Sprintf("seller with id %v cannot update", e.Id)
}
