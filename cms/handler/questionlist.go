package handler

import (
	"net/http"
	questionpb "stackoverflow/gunk/v1/question"
)


type questionListForm struct {
	Question []Questionc
	SearchTerm string
}


func (h Handler) ListQuestion(w http.ResponseWriter, r *http.Request) {
	res, err := h.questionSvc.ListQuestion(r.Context(),&questionpb.ListQuestionRequest{})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data := []Questionc{}


	if res != nil {

		for _,q := range res.Questions{
			data = append(data, Questionc{
				ID:          int(q.ID),
				UserId:      int(q.UserId),
				CategoryId:  int(q.CategoryId),
				Title:       q.Title,
				Description: q.Description,
			})
		}
	}
	h.pareseQuestionListTemplate(w,questionListForm{
		Question: data,
	})

}


func (h Handler) ListQuestionForAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.questionSvc.ListQuestion(r.Context(),&questionpb.ListQuestionRequest{})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data := []Questionc{}


	if res != nil {

		for _,q := range res.Questions{
			data = append(data, Questionc{
				ID:          int(q.ID),
				UserId:      int(q.UserId),
				CategoryId:  int(q.CategoryId),
				Title:       q.Title,
				Description: q.Description,
			})
		}
	}
	h.pareseQuestionListForAllTemplate(w,questionListForm{
		Question: data,
	})

}