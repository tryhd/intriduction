package dtos

type SkillCreateDTO struct {
	ID        string `json:"id" form:"id"`
	SkillName string `json:"skill_name" form:"skill_name" binding:"required"`
}
