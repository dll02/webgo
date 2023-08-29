package user

import "github.com/dll02/webgo/framework"

type UserService struct {
	container framework.Container
}

func NewUserService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	return &UserService{container: container}, nil
}

func (s *UserService) GetUsers() []User {
	return []User{
		{
			ID:   1,
			Name: "user1",
		},
		{
			ID:   2,
			Name: "user2",
		},
	}
}
