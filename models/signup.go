package models

import "encoding/json"

//Signup model
type Signup struct {
	ID              json.Number `form:"id" json:"id" xml:"id" binding:"gte=0,required"`
	Email           string      `form:"email" json:"email" xml:"email"  binding:"required,email"`
	Password        string      `form:"password" json:"password" xml:"password" binding:"required,min=12,max=50"`
	ConfirmPassword string      `form:"confirmPassword" json:"confirmPassword" xml:"confirmPassword" binding:"required,min=12,max=50"`
}
