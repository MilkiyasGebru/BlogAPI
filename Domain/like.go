package Domain

type Like struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `gorm:"foreignKey:ID" json:"userID"`
	BlogID uint `gorm:"foreignKey:ID" json:"blogID"`
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Blog   Blog `gorm:"foreignKey:BlogID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
