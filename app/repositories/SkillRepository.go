package repositories

import (
	"intoduction/app/models"

	"gorm.io/gorm"
)

type SkillRepository interface {
	InsertSkill(ct models.Skill) models.Skill
	DeleteSkill(ct models.Skill) models.Skill
	GetSkill(id int) []models.Skill
}

type skillConnection struct{ connection *gorm.DB }

func NewSkillRepository(dbConn *gorm.DB) SkillRepository { return &skillConnection{connection: dbConn} }

func (db *skillConnection) InsertSkill(skill models.Skill) models.Skill {
	db.connection.Create(&skill)
	db.connection.Find(&skill)
	return skill
}

func (db *skillConnection) DeleteSkill(skill models.Skill) models.Skill {
	db.connection.Debug().Where("profile_id", skill.ProfileID).Where("id", skill.ID).Delete(&skill)
	db.connection.Select("profile_id").Find(&skill)
	res := models.Skill{
		ProfileID: skill.ProfileID,
	}
	return res
}

func (db *skillConnection) GetSkill(id int) []models.Skill {
	var skills []models.Skill
	db.connection.Select("id", "skill", "level").Omit("created_at", "updated_at").Where("profile_id", id).Find(&skills)
	return skills
}
