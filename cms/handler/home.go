package handler

import (
	"net/http"
)

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	
	h.pareseHomeTemplate(w, r)
}

func (h Handler) About(w http.ResponseWriter, r *http.Request) {
	
	h.pareseAboutTemplate(w, r)
}