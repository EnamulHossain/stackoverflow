package handler

import (
	"net/http"
	categorypb "stackoverflow/gunk/v1/category"
)

type categorylistForm struct {
	Category   []CategoryCreate
	SearchTerm string
}

func (h Handler) ListCategory(w http.ResponseWriter, r *http.Request) {
	res, err := h.categorySvc.ListCategory(r.Context(), &categorypb.ListCategoryRequest{})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data := []CategoryCreate{}

	if res != nil {

		for _,c := range res.Categorys {
			data = append(data, CategoryCreate{
				ID:   int(c.ID),
				Name: c.Name,
			})
		}
	}
	h.pareseCatagoryListTemplate(w,categorylistForm{
		Category: data,
	})
}
