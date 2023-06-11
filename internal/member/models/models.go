package models

type Member struct {
	ID        int    `json:"id" gorm:"column:ID_MEMBER"`
	Username  string `json:"username" gorm:"column:USERNAME"`
	Gender    string `json:"gender" gorm:"column:GENDER"`
	SkinType  string `json:"skinType" gorm:"column:SKINTYPE"`
	SkinColor string `json:"skinColor" gorm:"column:SKINCOLOR"`
}
