package dtos

type ContentCreateDTO struct {
	ID                 string `json:"id" form:"id"`
	ContentName        string `json:"content_name" form:"content_name" binding:"required"`
	ContentTitle       string `json:"content_title" form:"content_title" binding:"required"`
	ContentImage       string `json:"content_image" form:"content_image"`
	ContentDescription string `json:"content_description" form:"content_description" binding:"required"`
}

type ContentUpdateDTO struct {
	ID                 string `json:"id" form:"id"`
	ContentName        string `json:"content_name" form:"content_name" binding:"required"`
	ContentTitle       string `json:"content_title" form:"content_title" binding:"required"`
	ContentImage       string `json:"content_image" form:"content_image"`
	ContentDescription string `json:"content_description" form:"content_description" binding:"required"`
}
