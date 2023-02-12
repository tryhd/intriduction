package repositories

import (
	"intoduction/app/models"

	"gorm.io/gorm"
)

type SkillRepository interface {
	InsertSkill(ct models.Skill) models.Skill
	DeleteSkill(ct models.Skill) models.Skill
	AllSkill() []models.Skill
	FindSkillByID(skillId string) models.Skill
}

type skillConnection struct{ connection *gorm.DB }

func NewSkillRepository(dbConn *gorm.DB) SkillRepository { return &skillConnection{connection: dbConn} }

func (db *skillConnection) InsertSkill(skill models.Skill) models.Skill {
	db.connection.Create(&skill)
	db.connection.Find(&skill)
	return skill
}

func (db *skillConnection) DeleteSkill(skill models.Skill) models.Skill {
	db.connection.Delete(&skill)
	db.connection.Find(&skill)
	return skill
}

func (db *skillConnection) AllSkill() []models.Skill {
	var skills []models.Skill
	db.connection.Find(&skills)
	return skills
}

func (db *skillConnection) FindSkillByID(skillID string) models.Skill {
	var skill models.Skill
	db.connection.Where("id =?", skillID).First(&skill)
	return skill
}
