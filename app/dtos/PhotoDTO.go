package dtos

type PhotoCreateDTO struct {
	Base64img string `json:"base64img"`
}

type PhotoUpdateDTO struct {
	ID string `json:"id" form:"id"`
}
