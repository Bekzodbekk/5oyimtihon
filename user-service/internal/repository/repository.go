package repository

import (
	"context"

	"github.com/Bekzodbekk/5oyimtihonProto/genproto/user"
)

type UserRepository interface {
	Register(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserResp, error)
	Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error)
	Verify(ctx context.Context, req *user.VerifyReq) (*user.VerifyResp, error)

	UpdateUser(ctx context.Context, req *user.UpdateUserReq) (*user.UpdateUserResp, error)
	DeleteUser(ctx context.Context, req *user.DeleteUserReq) (*user.DeleteUserResp, error)
	GetUserById(ctx context.Context, req *user.GetUserByIdReq) (*user.GetUserByIdResp, error)
}
