package exception

import "fmt"

type SellerNotFound struct {
	Id int64
}

func (e SellerNotFound) Error() string {
	return fmt.Sprintf("seller with id %v not found", e.Id)
}
