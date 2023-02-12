package dtos

type ClientCreateDTO struct {
	ID          string `json:"id" form:"id"`
	ClientName  string `json:"client_name" form:"client_name" binding:"required"`
	ClientTitle string `json:"client_title" form:"client_title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Status      uint8  `json:"status" form:"status" binding:"required"`
	ClientImage string `json:"client_image" form:"client_image" binding:"required"`
}
type ClientUpdateDTO struct {
	ID          string `json:"id" form:"id"`
	ClientName  string `json:"client_name" form:"client_name" binding:"required"`
	ClientTitle string `json:"client_title" form:"client_title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Status      uint8  `json:"status" form:"status" binding:"required"`
	ClientImage string `json:"client_image" form:"client_image" binding:"required"`
}

type ClientRestoreDTO struct {
	ID string `json:"id" form:"id"`
}
