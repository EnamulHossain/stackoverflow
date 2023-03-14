package handler

import (
	"fmt"
	"net/http"
	userpb "stackoverflow/gunk/v1/user"
	"stackoverflow/usermgm/storage"
)

type Ulist struct {
	Users []storage.User
	SearchTerm string
}

func (c Handler) ListUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request handler")
	res, err := c.usermgmSvc.ListUser(r.Context(), &userpb.ListUserRequest{})
	if err!=nil {
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

	fmt.Println("---user list---")
	fmt.Printf("%#v", res.Users)
	fmt.Println("---user list---")
	c.pareseUserListTemplate(w, res)
}
