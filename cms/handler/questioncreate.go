package handler

import (
	"log"
	"net/http"
	categorypb "stackoverflow/gunk/v1/category"
	questionpb "stackoverflow/gunk/v1/question"
	"strconv"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)

type Questionc struct {
	ID          int
	UserId      int
	CategoryId  int
	Title       string
	Description string
	Name        string
	PublishedAt time.Time
}

func (q Questionc) Validate() error {
	return validation.ValidateStruct(&q,
		validation.Field(&q.Title,
			validation.Required.Error("The Title field is required."),
		),
	)
}

type QuestionForm struct {
	Category  []CategoryCreate
	Question  Questionc
	FormError map[string]error
	CSRFToken string
}

func (h Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {

	Category, err := h.categorySvc.ListCategory(r.Context(), &categorypb.ListCategoryRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	data := []CategoryCreate{}

	if Category != nil {

		for _, u := range Category.Categorys {
			data = append(data, CategoryCreate{
				ID:   int(u.GetID()),
				Name: u.Name,
			})
		}
	}
	h.pareseCreateQuestionTemplate(w, QuestionForm{
		Category:  data,
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) CreateQuestionPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var form QuestionForm
	var qc Questionc

	if err := h.decoder.Decode(&qc, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	form.Question = qc

	if err := qc.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			for key, val := range vErr {
				form.FormError[strings.Title(key)] = val
			}
		}
		h.pareseCreateQuestionTemplate(w, form)
		return
	}

	userID := h.sessionManager.GetString(r.Context(), "userID")

	uID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	_, err = h.questionSvc.CreateQuestion(r.Context(), &questionpb.CreateQuestionRequest{
		UserId:      int32(uID),
		CategoryId:  int32(qc.CategoryId),
		Title:       qc.Title,
		Description: qc.Description,
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users/dashboard", http.StatusSeeOther)
}
