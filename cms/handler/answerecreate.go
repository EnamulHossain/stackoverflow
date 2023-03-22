package handler

import (
	"fmt"
	"log"
	"net/http"
	answerepb "stackoverflow/gunk/v1/answere"
	questionpb "stackoverflow/gunk/v1/question"
	"strconv"
	"strings"

	// questionpb "stackoverflow/gunk/v1/question"

	// "github.com/go-chi/chi"
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
	FormError map[string]error
	CSRFToken string
	Form      Questionc
}

// func (h Handler) CreateAnswere(w http.ResponseWriter, r *http.Request) {

// Question, err := h.questionSvc.ListQuestion(r.Context(),&questionpb.ListQuestionRequest{})
// if err != nil {
// 	log.Fatalln(err)
// }

// data := []Questionc{}

// if Question != nil {

// 	for _, q := range Question.Questions {
// 		data = append(data,Questionc{
// 			ID:          int(q.ID),
// 			UserId:      int(q.UserId),
// 			CategoryId:  int(q.CategoryId),
// 			Title:       q.Title,
// 			Description: q.Description,
// 		})
// 	}
// }
// h.pareseCreateAnswereTemplate(w, AnswereForm{
// 	Question:  data,
// 	CSRFToken: nosurf.Token(r),
// })
// }

func (h Handler) CreateAnswere(w http.ResponseWriter, r *http.Request) {

	// zid := chi.URLParam(r, "ID")
	// fmt.Println("yyyyyyyyyyyyyyyyyyyyyyyyyyy",zid)

	Url := r.URL.Path
	id := strings.ReplaceAll(Url, "/users/answere/create/", "")

	uid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := h.questionSvc.EditQuestion(r.Context(), &questionpb.EditQuestionRequest{
		ID: int32(uid),
	})
	if err != nil {
		log.Fatalln(err)
	}

	Data := AnswereForm{
		CSRFToken: nosurf.Token(r),
		Form: Questionc{
			ID:          int(res.Question.ID),
			UserId:      int(res.Question.UserId),
			CategoryId:  int(res.Question.CategoryId),
			Title:       res.Question.Title,
			Description: res.Question.Description,
		},
	}

	h.pareseCreateAnswereTemplate(w, Data)

}

func (h Handler) CreateAnswerePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println(r.PostForm)

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

	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	fmt.Println(uID,ac.QuestionId,ac.Answere)
	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

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

	h.pareseCreateAnswereTemplate(w, form)
}
