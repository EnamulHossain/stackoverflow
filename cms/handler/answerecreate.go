package handler

import (
	"fmt"
	"log"
	"net/http"
	answerepb "stackoverflow/gunk/v1/answere"
	questionpb "stackoverflow/gunk/v1/question"
	"strconv"
	"strings"


	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)

type Answere struct {
	ID         int
	UserId     int
	QuestionId int
	Answere    string
	IsCorrect  bool
}

func (a Answere) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Answere,
			validation.Required.Error("The answere field is required."),
		),
	)
}

type AnswereForm struct {
	Question  []Questionc
	Answere   Answere
	An        []Answere
	FormError map[string]error
	CSRFToken string
	Form      Questionc
	USERID int
}

func (h Handler) CreateAnswere(w http.ResponseWriter, r *http.Request) {

	// zid := chi.URLParam(r, "ID")
	// its not working

	Url := r.URL.Path
	id := strings.ReplaceAll(Url, "/users/answere/create/", "")

	uid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	res, err := h.questionSvc.EditQuestion(r.Context(), &questionpb.EditQuestionRequest{
		ID: int32(uid),
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	// ANSWERE LIST Start

	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	CurrentPage := 1
	pn := r.FormValue("page")
	if pn != "" {
		CurrentPage, err = strconv.Atoi(pn)
		if err != nil {
			CurrentPage = 1
		}
	}

	offset := 0
	if CurrentPage > 1 {
		offset = (CurrentPage * LimitPerPage) - LimitPerPage
	}

	
	re, err := h.answereSvc.AnswereList(r.Context(), &answerepb.AnswereListRequest{
		ID:     int32(uid),
		Limit:  LimitPerPage,
		Offset: int32(offset),
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	dat := []Answere{}

	if re != nil {
		for _, a := range re.Answere {
			dat = append(dat, Answere{
				ID:         int(a.ID),
				UserId:     int(a.UserId),
				QuestionId: int(a.QuestionId),
				Answere:    a.Answere,
				IsCorrect:  a.IsCorrect,
			})
		}
	}
	// ANSWERE LIST End



	// Auth ID
	userID := h.sessionManager.GetString(r.Context(), "userID")

	uID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	// Auth ID

	Data := AnswereForm{
		CSRFToken: nosurf.Token(r),
		Form: Questionc{
			ID:          int(res.Question.ID),
			UserId:      int(res.Question.UserId),
			CategoryId:  int(res.Question.CategoryId),
			Title:       res.Question.Title,
			Description: res.Question.Description,
		},
		An: dat,
		USERID: uID,
	}
	h.pareseCreateAnswereTemplate(w, Data)
}


func (h Handler) CreateAnswerePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}


	var form AnswereForm
	var ac Answere

	if err := h.decoder.Decode(&ac, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	form.Answere = ac

	if err := ac.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			for key, val := range vErr {
				form.FormError[strings.Title(key)] = val
			}
		}
		h.pareseCreateAnswereTemplate(w, form)
		return
	}
	userID := h.sessionManager.GetString(r.Context(), "userID")

	uID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	_, err = h.answereSvc.AnswereCreate(r.Context(), &answerepb.AnswereCreateRequest{
		UserId:     int32(uID),
		QuestionId: int32(ac.QuestionId),
		Answere:    ac.Answere,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/users/answere/create/%d", ac.QuestionId), http.StatusSeeOther)
}
