package repositories

import (
	"intoduction/app/models"

	"gorm.io/gorm"
)

type SkillRepository interface {
	InsertSkill(ct models.Skill) models.Skill
	DeleteSkill(ct models.Skill) models.Skill
	GetSkill() []models.Skill
}

type skillConnection struct{ connection *gorm.DB }

func NewSkillRepository(dbConn *gorm.DB) SkillRepository { return &skillConnection{connection: dbConn} }

func (db *skillConnection) InsertSkill(skill models.Skill) models.Skill {
	db.connection.Create(&skill)
	db.connection.Find(&skill)
	return skill
}

func (db *skillConnection) DeleteSkill(skill models.Skill) models.Skill {
	db.connection.Where("profile_id", skill.ProfileID).Delete(&skill)
	db.connection.Find(&skill)
	return skill
}

func (db *skillConnection) GetSkill() []models.Skill {
	var skills []models.Skill
	db.connection.Find(&skills)
	return skills
}
