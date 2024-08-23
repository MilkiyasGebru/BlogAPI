package Domain

type Comment struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `json:"description"`
	BlogID      uint
	UserID      uint
	User        User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Blog        Blog `gorm:"foreignKey:BlogID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
