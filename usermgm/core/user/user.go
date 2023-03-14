package user

import (
	"fmt"

	"stackoverflow/usermgm/storage"

	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	Register(storage.User) (*storage.User, error)
	GetUserByUsername(string) (*storage.User, error)
	ListUser() ([]storage.User, error)
	DeleteUser(id int32) error
	AdminRegister(u storage.User) (*storage.User, error)
}

type CoreUser struct {
	store UserStore
}

func NewCoreUser(us UserStore) *CoreUser {
	return &CoreUser{
		store: us,
	}
}

func (cu CoreUser) Register(u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u.Password = string(hashPass)
	ru, err := cu.store.Register(u)
	if err != nil {
		return nil, err
	}

	if ru == nil {
		return nil, fmt.Errorf("unable to register")
	}

	return ru, nil
}



func (cu CoreUser) AdminRegister(u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u.Password = string(hashPass)
	ru, err := cu.store.AdminRegister(u)
	if err != nil {
		return nil, err
	}

	if ru == nil {
		return nil, fmt.Errorf("unable to register")
	}

	return ru, nil
}




func (cu CoreUser) Login(l storage.Login) (*storage.User, error) {
	u, err := cu.store.GetUserByUsername(l.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password)); err != nil {
		return nil, err
	}
	return u, nil
}

func (cu CoreUser) ListUser() ([]storage.User, error) {

	list, err := cu.store.ListUser()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (cu CoreUser) DeleteUser(id int32) error{
	fmt.Println("id in core ", id)
	err := cu.store.DeleteUser(id)
	if err != nil {
		fmt.Println("errrrrrrrrrrrrrr", err)
		return err
	}

	return nil
}
