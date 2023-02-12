package dtos

type EducationCreateDTO struct {
	ProfileID   int    `json:"profileId"`
	School      string `json:"school"`
	Degree      string `json:"degree"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	City        string `json:"city"`
	Description string `json:"description"`
}

type EducationUpdateDTO struct {
	ID string `json:"id" form:"id"`
}
