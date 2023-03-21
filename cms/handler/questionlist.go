package handler

import (
	"net/http"
	questionpb "stackoverflow/gunk/v1/question"
)


type questionListForm struct {
	Question []Question
	SearchTerm string
}


func (h Handler) ListQuestion(w http.ResponseWriter, r *http.Request) {
	res, err := h.questionSvc.ListQuestion(r.Context(),&questionpb.ListQuestionRequest{})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data := []Question{}


	if res != nil {

		for _,q := range res.Questions{
			data = append(data, Question{
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