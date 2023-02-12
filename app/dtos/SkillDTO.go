package dtos

type SkillCreateDTO struct {
	ProfileID int    `json:"profile_id" form:"profile_id"`
	Level     string `json:"level" form:"level" binding:"required"`
	Skill     string `json:"skill" form:"skill" binding:"required"`
}
