package user

import (
	"context"

	userpb "stackoverflow/gunk/v1/user"
	"stackoverflow/usermgm/storage"
)

type CoreUser interface {
	Register(storage.User) (*storage.User, error)
	AdminRegister(u storage.User) (*storage.User, error)
	Login(storage.Login) (*storage.User, error)
	ListUser() ([]storage.User, error)
	DeleteUser(id int32) error
	GetUserByID(id int32) (*storage.User, error)
	UpdateUser(u storage.User) (*storage.User, error)
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
			IsAdmin:   false,
			IsActive:  u.IsActive,
		},
	}, nil
}

func (us UserSvc) AdminRegister(ctx context.Context, r *userpb.AdminRegisterRequest) (*userpb.AdminRegisterResponse, error) {
	user := storage.User{
		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		Email:     r.GetEmail(),
		Username:  r.GetUsername(),
		Password:  r.GetPassword(),
		IsAdmin:   true,
		IsActive:  true,
	}

	if err := user.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := us.core.AdminRegister(user)
	if err != nil {
		return nil, err
	}

	return &userpb.AdminRegisterResponse{
		User: &userpb.User{
			ID:        int32(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			IsActive:  true,
			IsAdmin:   true,
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

	users, err := us.core.ListUser()
	if err != nil {
		return nil, err
	}

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

	return &userpb.ListUserResponse{
		Users: list,
	}, nil

}

func (us UserSvc) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {

	err := us.core.DeleteUser(req.GetID())

	if err != nil {
		return nil, err
	}

	return &userpb.DeleteUserResponse{}, nil
}

func (us UserSvc) EditUser(ctx context.Context, r *userpb.EditUserRequest) (*userpb.EditUserResponse, error) {
	res, err := us.core.GetUserByID(r.GetID())
	if err != nil {
		return nil, err
	}

	return &userpb.EditUserResponse{
		User: &userpb.User{
			ID:        int32(res.ID),
			FirstName: res.FirstName,
			LastName:  res.LastName,
			Username:  res.Username,
			Email:     res.Email,
			IsActive:  res.IsActive,
			IsAdmin:   res.IsAdmin,
		},
	}, err

}

func (us UserSvc) UpdateUser(ctx context.Context, r *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {

	user := storage.User{
		ID:        int(r.GetID()),
		FirstName: r.GetFirstName(),
		LastName:  r.GetLastName(),
		Email:     r.GetEmail(),
		Username:  r.GetUsername(),
		Password:  r.GetPassword(),
		IsActive:  r.GetIsActive(),
		IsAdmin:   r.GetIsAdmin(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	_, err := us.core.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{}, nil

}
