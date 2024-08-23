package Domain

type Blog struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"foreignKey:ID"`
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type IBlogUseCase interface {
	CreateBlog(blog *Blog) error
	GetBlogById(id uint) (*Blog, error)
	GetAllBlogsByUserId(userId uint) ([]Blog, error)
	GetAllBlogs() ([]Blog, error)
	UpdateBlog(id uint, blog *Blog) error
	DeleteBlog(id uint) error
}
