package handler

import (
	"net/http"
)

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	
	h.pareseHomeTemplate(w, r)
}

