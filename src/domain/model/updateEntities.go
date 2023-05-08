package model

type UpdateCustomer struct {
	Lastname  string `json:"lastname" bson:"lastname" binding:"required"`
	Firstname string `json:"firstname" bson:"firstname" binding:"required"`
	Email     string `json:"email" bson:"email" binding:"required,email"`
}

type UpdateSeller struct {
	Name  string `json:"name" bson:"name" binding:"required"`
	Email string `json:"email" bson:"email" binding:"required,email"`
}

type UpdateProduct struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Category    string  `json:"category" binding:"required"`
}
