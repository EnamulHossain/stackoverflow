package user

import (
	"context"
	"fmt"

	userpb "stackoverflow/gunk/v1/user"
	"stackoverflow/usermgm/storage"
)

type CoreUser interface {
	Register(storage.User) (*storage.User, error)
	Login(storage.Login) (*storage.User, error)
	ListUser() ([]storage.User, error)
}

type UserSvc struct {
	userpb.UnimplementedUserServiceServer
	core CoreUser
}

func NewUserSvc(cu CoreUser) *UserSvc {
	return &UserSvc{
		core: cu,
	}
}

func (us UserSvc) Register(ctx context.Context, r *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	user := storage.User{
		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		Email:     r.GetEmail(),
		Username:  r.GetUsername(),
		Password:  r.GetPassword(),
		IsAdmin:   false,
		IsActive:  true,

	}

	if err := user.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := us.core.Register(user)
	if err != nil {
		return nil, err
	}

	return &userpb.RegisterResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			IsActive:  u.IsActive,
			IsAdmin:   false,
		},
	}, nil
}

func (us UserSvc) Login(ctx context.Context, r *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	login := storage.Login{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}

	if err := login.Validate(); err != nil {
		return nil, err
	}

	u, err := us.core.Login(login)
	if err != nil {
		return nil, err
	}

	return &userpb.LoginResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			IsActive:  u.IsActive,
			IsAdmin:   u.IsAdmin,
		},
	}, nil
}

func (us UserSvc) ListUser(ctx context.Context, req *userpb.ListUserRequest) (*userpb.ListUserResponse, error) {
	
		users, err :=us.core.ListUser()
		if err != nil {
			return nil, err
		}
		fmt.Println("..................",users)
	
		list := make([]*userpb.User, len(users))
		for i, u := range users {
			list[i] = &userpb.User{
				ID:        int32(u.ID),
				FirstName: u.FirstName,
				LastName:  u.LastName,
				Username:  u.Username,
				Email:     u.Email,
				IsActive:  u.IsActive,
				IsAdmin:   u.IsAdmin,
			}
		}


		fmt.Println("..................",list)
		
		return &userpb.ListUserResponse{
			Users: list,
		}, nil


}