package exception

import "fmt"

type CannotMapOrderState struct {
	State string
}

func (e CannotMapOrderState) Error() string {
	return fmt.Sprintf("cannot map order state %s", e.State)
}

type OrderNotFound struct {
	Id int64
}

func (e OrderNotFound) Error() string {
	return fmt.Sprintf("order with id %v not found", e.Id)
}

type OrderCannotUpdate struct {
	Id int64
}

func (e OrderCannotUpdate) Error() string {
	return fmt.Sprintf("order with id %v cannot update", e.Id)
}

type OrderInvalidTransitionState struct {
	Id int64
}

func (e OrderInvalidTransitionState) Error() string {
	return fmt.Sprintf("invalid transition state for order with id %v", e.Id)
}
