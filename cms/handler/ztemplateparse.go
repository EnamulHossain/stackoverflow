package handler

import (
	"log"
	"net/http"
)


func (h Handler) pareseRegisterTemplate(w http.ResponseWriter, form RegisterUserForm) {
	t := h.Templates.Lookup("register.html")
	if t == nil {
		log.Println("unable to lookup register template")
		h.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, form); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}


func (c Handler) pareseHomeTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("home.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}




func (c Handler) pareseUserListTemplate(w http.ResponseWriter, data any) {
	t := c.Templates.Lookup("list-user.html")
	if t == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}