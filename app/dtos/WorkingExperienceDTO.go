package dtos

type WorkingExperienceCreateDTO struct {
	ProfileID         int    `json:"profileId"`
	WorkingExperience string `json:"workingExperience"`
}

type WorkingExperienceUpdateDTO struct {
	ProfileID         int    `json:"profileId"`
	WorkingExperience string `json:"workingExperience"`
}
