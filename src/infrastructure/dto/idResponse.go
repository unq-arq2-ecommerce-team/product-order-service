package dto

type IdResponse struct {
	Id int64 `json:"id"`
}

func NewIdResponse(id int64) *IdResponse {
	return &IdResponse{Id: id}
}
