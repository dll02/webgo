package user

const UserKey = "App:user"

type Service interface {
	// 获取全部用户
	GetUsers() []User
}

type User struct {
	ID   int
	Name string
}
