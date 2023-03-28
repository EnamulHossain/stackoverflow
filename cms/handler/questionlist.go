package handler

import (
	"log"
	"net/http"
	questionpb "stackoverflow/gunk/v1/question"
	"strconv"
)

const LimitPerPage = 10

type questionListForm struct {
	Question []Questionc
	SearchTerm string
}


func (h Handler) ListQuestion(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	var err error
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

	st := r.FormValue("SearchTerm")

	res, err := h.questionSvc.ListQuestion(r.Context(), &questionpb.ListQuestionRequest{
		SearchTerm: st,
		Limit:      LimitPerPage,
		Offset:     int32(offset),
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data := []Questionc{}


	if res != nil {

		for _, q := range res.Questions {
			data = append(data, Questionc{
				ID:          int(q.ID),
				UserId:      int(q.UserId),
				CategoryId:  int(q.CategoryId),
				Title:       q.Title,
				Description: q.Description,
				Name:        q.Name,
			})
		}
	}
	h.pareseQuestionListTemplate(w,questionListForm{
		Question: data,
	})

}

func (h Handler) ListQuestionForAll(w http.ResponseWriter, r *http.Request) {


	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	var err error
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

	st := r.FormValue("SearchTerm")

	res, err := h.questionSvc.ListQuestion(r.Context(), &questionpb.ListQuestionRequest{
		SearchTerm: st,
		Limit:      LimitPerPage,
		Offset:     int32(offset),
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data := []Questionc{}


	if res != nil {

		for _, q := range res.Questions {
			data = append(data, Questionc{
				ID:          int(q.ID),
				UserId:      int(q.UserId),
				CategoryId:  int(q.CategoryId),
				Title:       q.Title,
				Description: q.Description,
				Name:        q.Name,
			})
		}
	}
	h.pareseQuestionListForAllTemplate(w,questionListForm{
		Question: data,
	})

}