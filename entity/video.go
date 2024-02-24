package entity

type Person struct {
	FirsName string `json:"firstname" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Age      int    `json:"age" binding:"gte=1,lte=115"`
	Email    string `json:"email" validate:"required,email"`
}
type Video struct {
	Title       string `json:"title" binding:"min=2,max=30" validate:"is-cool"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
