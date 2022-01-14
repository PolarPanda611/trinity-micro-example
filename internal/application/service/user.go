// Author: Daniel TAN
// Date: 2021-08-18 23:47:20
// LastEditors: Daniel TAN
// LastEditTime: 2021-10-02 01:26:04
// FilePath: /trinity-micro/example/crud/internal/application/service/user.go
// Description:
package service

import (
	"context"
	"sync"

	"trinity-micro-api/internal/application/dto"
	"trinity-micro-api/internal/application/repository"

	"github.com/PolarPanda611/trinity-micro"
	"github.com/PolarPanda611/trinity-micro/core/e"
)

func init() {
	trinity.RegisterInstance("UserService", &sync.Pool{
		New: func() interface{} {
			return new(userServiceImpl)
		},
	})
}

type UserService interface {
	GetUserID(ctx context.Context, req *dto.GetUserByIDRequest) (*dto.UserInfoResponse, error)
	ListUser(ctx context.Context, req *dto.ListUserRequest) (*dto.ListUserResponse, error)
	CreateUser(ctx context.Context, createUserForm *dto.CreateUserRequest) (*dto.UserInfoResponse, error)
	UpdateUser(ctx context.Context, updateUserForm *dto.UpdateUserRequest) error
}

type userServiceImpl struct {
	UserRepo repository.UserRepository `container:"autowire:true;resource:UserRepository"`
}

func (s *userServiceImpl) ListUser(ctx context.Context, req *dto.ListUserRequest) (*dto.ListUserResponse, error) {
	users, err := s.UserRepo.ListUser(ctx, req.Tenant, req.ParsePageQuery())
	if err != nil {
		return nil, err
	}
	userCount, err := s.UserRepo.CountUser(ctx, req.Tenant, req.ParseQuery())
	if err != nil {
		return nil, err
	}
	return dto.NewListUserResponse(users, req.PageSize, req.PageNum, userCount), nil
}

func (s *userServiceImpl) GetUserID(ctx context.Context, req *dto.GetUserByIDRequest) (*dto.UserInfoResponse, error) {
	user, err := s.UserRepo.GetUserByID(ctx, req.Tenant, req.ID)
	if err != nil {
		return nil, err
	}
	return dto.NewUserInfoResponse(user), nil
}

func (s *userServiceImpl) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserInfoResponse, error) {
	if err := req.CreateUserFrom.Validate(); err != nil {
		return nil, e.NewError(e.Info, e.ErrInvalidRequest, err.Error())
	}
	user, err := s.UserRepo.CreateUser(ctx, req.Tenant, req.Parse())
	if err != nil {
		return nil, err
	}
	return dto.NewUserInfoResponse(user), nil
}

func (s *userServiceImpl) UpdateUser(ctx context.Context, req *dto.UpdateUserRequest) error {
	if err := req.UpdateUserForm.Validate(); err != nil {
		return e.NewError(e.Info, e.ErrInvalidRequest, err.Error())
	}
	return s.UserRepo.UpdateUser(ctx, req.Tenant, req.ID, req.Version, req.UpdateUserForm)
}
