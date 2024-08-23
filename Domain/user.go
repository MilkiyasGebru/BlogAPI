package Domain

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}


type IUserUseCase interface {
	CreateUser(user *User) (error)
	LoginUser(user *User) (string,string,error)
	GetUserByEamil(eamil string) (User,error)
}