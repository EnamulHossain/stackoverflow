package handler

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	answerepb "stackoverflow/gunk/v1/answere"
	categorypb "stackoverflow/gunk/v1/category"
	questionpb "stackoverflow/gunk/v1/question"
	userpb "stackoverflow/gunk/v1/user"

	"github.com/Masterminds/sprig"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/form"
	"google.golang.org/grpc"
)

type usermgmService struct {
	userpb.UserServiceClient
}

type categoryService struct {
	categorypb.CategoryServiceClient
}

type questionService struct {
	questionpb.QuestionServiceClient
}

type answereService struct {
	answerepb.AnswereServiceClient
}

type Handler struct {
	sessionManager *scs.SessionManager
	decoder        *form.Decoder
	usermgmSvc     usermgmService
	categorySvc    categoryService
	questionSvc    questionService
	answereSvc     answereService
	Templates      *template.Template
	staticFiles    fs.FS
	templateFiles  fs.FS
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

type ErrorPage struct {
	Code    int
	Message string
}

func (h Handler) Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	ep := ErrorPage{
		Code:    code,
		Message: error,
	}

	tf := "default"
	switch code {
	case 400, 401, 402, 403, 404:
		tf = "4xx"
	case 500, 501, 503:
		tf = "5xx"
	}

	tpl := fmt.Sprintf("%s.html", tf)
	t := h.Templates.Lookup(tpl)
	if t == nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	if err := t.Execute(w, ep); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)

	}
}

func NewHandler(sm *scs.SessionManager, formDecoder *form.Decoder, usermgmConn *grpc.ClientConn, staticFiles, templateFiles fs.FS) *chi.Mux {
	h := &Handler{
		sessionManager: sm,
		decoder:        formDecoder,
		usermgmSvc:     usermgmService{userpb.NewUserServiceClient(usermgmConn)},
		categorySvc:    categoryService{categorypb.NewCategoryServiceClient(usermgmConn)},
		questionSvc:    questionService{questionpb.NewQuestionServiceClient(usermgmConn)},
		answereSvc:     answereService{answerepb.NewAnswereServiceClient(usermgmConn)},
		staticFiles:    staticFiles,
		templateFiles:  templateFiles,
	}

	h.ParseTemplates()
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(Method)

	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Get("/", h.Home)
		r.Get("/about", h.About)
		r.Get("/register", h.Register)
		r.Post("/register", h.RegisterPost)
		r.Get("/login", h.Login)
		r.Post("/login", h.LoginPostHandler)
		r.Get("/logout", h.LogoutHandler)

		r.Get("/question/list", h.ListQuestionForAll)

	})

	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.FS(h.staticFiles))))

	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Route("/users", func(r chi.Router) {

			r.Get("/dashboard", h.Userdashboard)

			r.Get("/question/create", h.CreateQuestion)
			r.Post("/question/store", h.CreateQuestionPost)
			r.Post("/question/publish", h.Publish)

			r.Get("/answere/create/{{.ID}}", h.CreateAnswere)
			r.Post("/answere/store", h.CreateAnswerePost)
			r.Post("/answere/iscorrect", h.AnswereISPost)

		})

		r.Get("/logout", h.LogoutHandler)
	})

	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Use(h.AdminAuthentication)

		r.Route("/admin", func(r chi.Router) {
			r.Get("/list", h.ListUser)
			r.Get("/r", h.AdminReg)
			r.Post("/register", h.AdminRegisterPost)

			r.Get("/catergoy/create", h.CreateCategory)
			r.Post("/category/store", h.CreateCategoryPost)
			r.Get("/catergoy/list", h.ListCategory)

			r.Get("/question/list", h.ListQuestion)

		})
	})

	return r
}

func Method(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			switch strings.ToLower(r.PostFormValue("_method")) {
			case "put":
				r.Method = http.MethodPut
			case "patch":
				r.Method = http.MethodPatch
			case "delete":
				r.Method = http.MethodDelete
			default:
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (h Handler) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := h.sessionManager.GetString(r.Context(), "userID")
		uID, err := strconv.Atoi(userID)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if uID <= 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h Handler) AdminAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		IsAdmin := h.sessionManager.GetBool(r.Context(), "IsAdmin")
		if !IsAdmin {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) ParseTemplates() error {
	templates := template.New("web-templates").Funcs(template.FuncMap{
		"calculatePreviousPage": func(currentPageNumber int) int {
			if currentPageNumber == 1 {
				return 0
			}

			return currentPageNumber - 1
		},

		"calculateNextPage": func(currentPageNumber, totalPage int) int {
			if currentPageNumber == totalPage {
				return 0
			}

			return currentPageNumber + 1
		},
	}).Funcs(sprig.FuncMap())

	tmpl := template.Must(templates.ParseFS(h.templateFiles, "*/*/*.html", "*/*.html", "*.html"))
	if tmpl == nil {
		log.Fatalln("unable to parse templates")
	}

	h.Templates = tmpl
	return nil
}
