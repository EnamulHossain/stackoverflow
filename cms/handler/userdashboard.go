package handler

import (
	"fmt"
	"log"
	"net/http"
	questionpb "stackoverflow/gunk/v1/question"
	userpb "stackoverflow/gunk/v1/user"
	"strconv"

	"github.com/justinas/nosurf"
)

type UserdashboardForm struct {
	IsAdmin   bool
	User      []User
	QList     []Questionc
	CSRFToken string
}

func (h Handler) Userdashboard(w http.ResponseWriter, r *http.Request) {

	userID := h.sessionManager.GetString(r.Context(), "userID")
	isadmin := h.sessionManager.GetBool(r.Context(), "IsAdmin")
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
	R = append(R, User{
		ID:        int(Res.User.ID),
		FirstName: Res.User.FirstName,
		LastName:  Res.User.LastName,
		Email:     Res.User.Email,
		Username:  Res.User.Username,
		IsActive:  Res.User.IsActive,
	})

	// question list of this user

	Ql, err := h.questionSvc.GetQueByUser(r.Context(), &questionpb.GetQueByUserRequest{
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
				PublishedAt: a.PublishedAt.AsTime(),
			})
		}
	}
	// question list of this user  END

	Data := UserdashboardForm{
		IsAdmin:   isadmin,
		User:      R,
		QList:     dat,
		CSRFToken: nosurf.Token(r),
	}

	h.pareseUserDashboardTemplate(w, Data)
}

func (h Handler) Publish(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	id := r.FormValue("ID")

	uid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	Res, err := h.questionSvc.PublishedQuestion(r.Context(), &questionpb.PublishedQuestionRequest{
		ID: int32(uid),
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(Res)
	http.Redirect(w, r, "/users/dashboard", http.StatusSeeOther)
	// h.pareseUserDashboardTemplate(w, r)
}
