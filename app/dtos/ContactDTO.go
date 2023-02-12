package dtos

type ContactCreateDTO struct {
	ID             string `json:"id" form:"id"`
	ContactName    string `json:"contact_name" form:"contact_name" binding:"required"`
	ContactEmail   string `json:"contact_email" form:"contact_email" binding:"required"`
	ContactMessage string `json:"contact_message" form:"contact_message"`
	Status         uint8  `json:"status" form:"status" binding:"required"`
}

type ContactUpdateDTO struct {
	ID             string `json:"id" form:"id"`
	ContactName    string `json:"contact_name" form:"contact_name" binding:"required"`
	ContactEmail   string `json:"contact_email" form:"contact_email"`
	ContactMessage string `json:"contact_message" form:"contact_message"`
	Status         uint8  `json:"status" form:"status" binding:"required"`
}

type ContactRestoreDTO struct {
	ID string `json:"id" form:"id"`
}
