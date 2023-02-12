package dtos

type PortfolioCreateDTO struct {
	ID                   string `json:"id" form:"id"`
	PortfolioTitle       string `json:"portfolio_title" form:"portfolio_title" binding:"required"`
	PortfolioImage       string `json:"portfolio_image" form:"portfolio_image" binding:"required"`
	PortfolioDescription string `json:"portfolio_description" form:"portfolio_description" binding:"required"`
	PortfolioLink        string `json:"portfolio_link" form:"portfolio_link" binding:"required"`
	Status               uint8  `json:"status" form:"status" binding:"required"`
}

type PortfolioUpdateDTO struct {
	ID                   string `json:"id" form:"id"`
	PortfolioTitle       string `json:"portfolio_title" form:"portfolio_title" binding:"required"`
	PortfolioImage       string `json:"portfolio_image" form:"portfolio_image"`
	PortfolioDescription string `json:"portfolio_description" form:"portfolio_description" binding:"required"`
	PortfolioLink        string `json:"portfolio_link" form:"portfolio_link" binding:"required"`
	Status               uint8  `json:"status" form:"status" binding:"required"`
}

type PortfolioRestoreDTO struct {
	ID string `json:"id" form:"id"`
}
