package service

import (
	"context"
	"user-service/internal/repository"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/user"
)

type Service struct {
	*user.UnimplementedUserServiceServer
	userRepo repository.UserRepository
}

func NewService(user repository.UserRepository) *Service {

	return &Service{
		userRepo: user,
	}
}

func (s *Service) Register(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserResp, error) {
	return s.userRepo.Register(ctx, req)
}
func (s *Service) Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error) {
	return s.userRepo.Login(ctx, req)
}
func (s *Service) Verify(ctx context.Context, req *user.VerifyReq) (*user.VerifyResp, error) {
	return s.userRepo.Verify(ctx, req)
}
func (s *Service) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	return s.userRepo.UpdateUser(ctx, req)
}
func (s *Service) DeleteUser(ctx context.Context, req *user.DeleteUserReq) (*user.DeleteUserResp, error) {
	return s.userRepo.DeleteUser(ctx, req)
}
func (s *Service) GetUserById(ctx context.Context, req *user.GetUserByIdReq) (*user.GetUserByIdResp, error) {
	return s.userRepo.GetUserById(ctx, req)
}
