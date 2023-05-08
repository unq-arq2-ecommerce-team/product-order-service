package dto

type NextIdResponse struct {
	LastErrorObject struct {
		N               int  `json:"n"`
		UpdatedExisting bool `json:"updatedExisting"`
	} `json:"lastErrorObject"`
	Value struct {
		Id  string `json:"_id"`
		Seq int64  `json:"seq" binding:"required"`
	} `json:"value"`
	Ok float64 `json:"ok"`
}
