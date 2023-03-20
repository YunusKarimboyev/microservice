package service

import (
	"context"

	u "github.com/double/test_microservice/genproto/user"
	"github.com/double/test_microservice/pkg/logger"
	"github.com/double/test_microservice/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	storage storage.IStorage
	Logger  logger.Logger
}

func NewUserService(db *sqlx.DB, log logger.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		Logger:  log,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *u.UserRequest) (*u.UserResponse, error) {
	res, err := s.storage.User().CreateUser(req)
	if err != nil {
		s.Logger.Error("error insert user", logger.Any("Error insert user", err))
		return &u.UserResponse{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}
	return res, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *u.UserId) (*u.UserResponse, error) {
	res, err := s.storage.User().GetUserById(req)
	if err != nil {
		s.Logger.Error("error get user", logger.Any("Error get user", err))
		return &u.UserResponse{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}

	return res, nil
}

func (s *UserService) GetUsersAll(ctx context.Context, req *u.UserListReq) (*u.Users, error){
	res, err := s.storage.User().GetUsersAll(req)
	if err != nil {
		s.Logger.Error("error get all user", logger.Any("error get all user", err))
		return &u.Users{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}
	return res, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *u.UserUpdateReq) (*u.UserResponse, error){
	res, err := s.storage.User().UpdateUser(req)
	if err != nil{
		s.Logger.Error("error update user", logger.Any("error update user", err))
		return &u.UserResponse{}, status.Error(codes.Internal, "something went wrong, please check user info")
	} 
	return res, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *u.UserDeleteReq) (*u.Users, error){
	res, err := s.storage.User().DeleteUser(req)
	if err != nil {
		s.Logger.Error("error delete user", logger.Any("error delete user", err))
		return &u.Users{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}
	return res, nil
}

func (s *UserService) GetUserByPostId(ctx context.Context, req *u.PostId) (*u.UserResponseForPost, error){
	res, err := s.storage.User().GetUserByPostId(req.PostId)
	if err != nil {
		s.Logger.Error("error delete user", logger.Any("error delete user", err))
		return &u.UserResponseForPost{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}
	return res, nil
}
