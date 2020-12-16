package service

import (
	"context"

	pb "week04/api/demo"
	"week04/internal/dao"
)

type UserService struct {
	dao *dao.UserDao
}

func NewUserService(dao *dao.UserDao) *UserService {
	return &UserService{dao: dao}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := s.dao.Query(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{Id: user.ID, Name: user.Name}, nil
}
