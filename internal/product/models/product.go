package models

type Product struct {
	ID    int     `json:"id" gorm:"column:ID_PRODUCT"`
	Name  string  `json:"productName" gorm:"column:PRODUCT_NAME"`
	Price float64 `json:"price" gorm:"column:PRICE"`
}

type ProductWithReview struct {
	Product *Product  `json:"product" gorm:"embedded" `
	Reviews []*Review `json:"review" gorm:"embedded" `
}

type LikeReview struct {
	ReviewID int `gorm:"column:ID_REVIEW" json:"reviewId"`
	MemberID int `gorm:"column:ID_MEMBER" json:"memberId"`
}

type ReviewData struct {
	ID          int    `gorm:"column:ID_REVIEW" json:"reviewId"`
	ProductID   int    `gorm:"column:ID_PRODUCT" json:"productId"`
	MemberID    int    `gorm:"column:ID_MEMBER" json:"memberId"`
	Username    string `gorm:"column:username" json:"username"`
	LikeCount   int    `gorm:"column:like_count" json:"likeCount"`
	Description string `gorm:"column:DESC_REVIEW" json:"descReview"`
	Gender      string `gorm:"column:gender" json:"gender"`
	SkinType    string `gorm:"column:skintype" json:"skinType"`
	SkinColor   string `gorm:"column:skincolor" json:"skinColor"`
}

type Review struct {
	ReviewData
}

func (Review) TableName() string {
	return "review_products"
}
