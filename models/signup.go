package models

//Signup model
type Signup struct {
	ID              string
	Email           string `form:"email" json:"email" xml:"email"  binding:"required,email"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required,min=12,max=50"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" xml:"confirmPassword" binding:"required,min=12,max=50"`
}
