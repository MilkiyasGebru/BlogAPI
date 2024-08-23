package usecases

import (
	"blogAPI/Domain"
	"errors"
)

type IBlogRepository interface {
	CreateBlog(blog *Domain.Blog) error
	GetBlogById(id uint) (*Domain.Blog, error)
	GetAllBlogsByUserId(userId uint) ([]Domain.Blog, error)
	GetAllBlogs() ([]Domain.Blog, error)
	UpdateBlog(id uint, blog *Domain.Blog) error
	DeleteBlog(id uint) error
}

type blogUseCase struct {
	blogRepo IBlogRepository
}

func NewBlogUseCase(br *IBlogRepository) Domain.IBlogUseCase {
	return &blogUseCase{
		blogRepo: *br,
	}
}

func (br *blogUseCase) CreateBlog(blog *Domain.Blog) error {
	return errors.New("Error")
}

func (br *blogUseCase) GetBlogById(id uint) (*Domain.Blog, error) {
	return nil, errors.New("Not Found")
}

func (br *blogUseCase) GetAllBlogsByUserId(id uint) ([]Domain.Blog, error) {
	//return make([]Domain.Blog, 0), errors.New("Empty")
	return []Domain.Blog{}, errors.New("Empty")
}

func (br *blogUseCase) GetAllBlogs() ([]Domain.Blog, error) {
	return []Domain.Blog{}, errors.New("Empty")
}

func (br *blogUseCase) UpdateBlog(id uint, blog *Domain.Blog) error {
	return errors.New("Not Implemented")
}

func (br *blogUseCase) DeleteBlog(id uint) error {
	return errors.New("Not Implemented")
}
