package handler

import (
	"net/http"
	userpb "stackoverflow/gunk/v1/user"
)

type Ulist struct {
	Users []User
	SearchTerm string
}


type User struct{
	ID        int         
	FirstName string       
	LastName  string       
	Email     string      
	Username  string       
	Password  string      
	IsActive  bool        
	IsAdmin   bool        
}

// type UserTempData struct {
// 	UserData PublicTempData
// }

func (c Handler) ListUser(w http.ResponseWriter, r *http.Request) {
	res, err := c.usermgmSvc.ListUser(r.Context(), &userpb.ListUserRequest{})
	if err!=nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	data := []User{}

if res != nil {

	for _,u := range res.Users {
		data = append(data, User{
			ID:        int(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Username:  u.Username,
			IsActive:  u.IsActive,
			IsAdmin:   u.IsAdmin,
		})
	}
}	
	c.pareseUserListTemplate(w, Ulist{
		Users:      data,
	})
}
