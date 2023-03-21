package handler

import (
	"log"
	"net/http"
	categorypb "stackoverflow/gunk/v1/category"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	// "github.com/go-playground/form"
	"github.com/justinas/nosurf"
)

type CategoryCreate struct {
	ID        int         
	Name      string      
}

func (c CategoryCreate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required.Error("The name field is required."),
		),
	)
}

type CategoryForm struct {
	Category  CategoryCreate
	FormError map[string]error
	CSRFToken string
}

func (h Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	h.pareseCatagoryCreateTemplate(w, CategoryForm{
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) CreateCategoryPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var form CategoryForm
	var cc CategoryCreate
	if err := h.decoder.Decode(&cc, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	form.Category = cc
	if err := cc.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			for key, val := range vErr {
				form.FormError[strings.Title(key)] = val
			}
		}
		h.pareseCatagoryCreateTemplate(w, form)
		return
	}

	_, err := h.categorySvc.CreateCategory(r.Context(), &categorypb.CreateCategoryRequest{
		Name: cc.Name,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/catergoy/list", http.StatusSeeOther)
}
