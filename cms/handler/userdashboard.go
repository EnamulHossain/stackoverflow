package handler

import (
	"log"
	"net/http"
	questionpb "stackoverflow/gunk/v1/question"
	userpb "stackoverflow/gunk/v1/user"
	"strconv"
)

type UserdashboardForm struct{
	User []User
	QList []Questionc
}

func (h Handler) Userdashboard(w http.ResponseWriter, r *http.Request) {


	userID := h.sessionManager.GetString(r.Context(), "userID")

	uID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}


	Res, err := h.usermgmSvc.EditUser(r.Context(), &userpb.EditUserRequest{
		ID: int32(uID),
	})
	if err != nil {
		log.Fatalln(err)
	}

	R := []User{}
	R= append(R, User{
		ID:        int(Res.User.ID),
		FirstName: Res.User.FirstName,
		LastName:  Res.User.LastName,
		Email:     Res.User.Email,
		Username:  Res.User.Username,
		IsActive:  Res.User.IsActive,
	})


	// question list of this user

	Ql, err := h.questionSvc.GetQueByUser(r.Context(),&questionpb.GetQueByUserRequest{
		ID: int32(uID),
	})
	if err != nil {
		log.Fatalln(err)
	}


	dat := []Questionc{}

	if Ql != nil {
		for _, a := range Ql.Questions {
			dat = append(dat, Questionc{
				ID:          int(a.ID),
				UserId:      int(a.UserId),
				CategoryId:  int(a.CategoryId),
				Title:       a.Title,
				Description: a.Description,
			})
		}
	}
	// question list of this user  END


	Data := UserdashboardForm{
		User:  R,
		QList: dat,
	}

	h.pareseUserDashboardTemplate(w, Data)
}