package postgres

import (
	"fmt"
	"log"

	"stackoverflow/usermgm/storage"
)

const registerQuery = `INSERT INTO users (
	first_name,
	last_name,
	username,
	email,
	password,
	is_active,
	is_admin
) VALUES(
	:first_name,
	:last_name,
	:username,
	:email,
	:password,
	:is_active,
	:is_admin
) RETURNING *`

func (s PostgresStorage) Register(u storage.User) (*storage.User, error) {
	stmt, err := s.DB.PrepareNamed(registerQuery)
	if err != nil {
		return nil, err
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("unable to insert user into db")
	}
	return &u, nil
}

const getUserByUsernameQuery = `SELECT * FROM users WHERE username=$1 AND deleted_at IS NULL`

func (ps PostgresStorage) GetUserByUsername(usernanme string) (*storage.User, error) {
	var user storage.User
	if err := ps.DB.Get(&user, getUserByUsernameQuery, usernanme); err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("unable to find user by username")
	}

	return &user, nil
}





const listQuery = `SELECT * FROM users WHERE deleted_at IS NULL`

func (s PostgresStorage) ListUser() ([]storage.User, error) {

	var user []storage.User
	if err := s.DB.Select(&user, listQuery); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return user, nil
}


