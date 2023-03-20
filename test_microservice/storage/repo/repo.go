package repo

import (
	u "github.com/double/test_microservice/genproto/user"
)

type UserStorageI interface {
	CreateUser(*u.UserRequest) (*u.UserResponse, error)
	GetUserById(*u.UserId) (*u.UserResponse, error)
	GetUsersAll(*u.UserListReq) (*u.Users, error)
	UpdateUser(*u.UserUpdateReq) (*u.UserResponse, error)
	DeleteUser(*u.UserDeleteReq) (*u.Users, error)
	GetUserByPostId(id int64) (*u.UserResponseForPost, error)
}
