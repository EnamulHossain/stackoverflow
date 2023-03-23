package storage

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
	ListUser   []User
}

type User struct {
	ID        int          `json:"id" form:"-" db:"id"`
	FirstName string       `json:"first_name" db:"first_name"`
	LastName  string       `json:"last_name" db:"last_name"`
	Email     string       `json:"email" db:"email"`
	Username  string       `json:"username" db:"username"`
	Password  string       `json:"password" db:"password"`
	IsActive  bool         `json:"is_active" db:"is_active"`
	IsAdmin   bool         `json:"is_admin" db:"is_admin"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName,
			validation.Required.Error("The first name field is required."),
		),
		validation.Field(&u.LastName,
			validation.Required.Error("The last name field is required."),
		),
		validation.Field(&u.Username,
			validation.Required.When(u.ID == 0).Error("The username field is required."),
		),
		validation.Field(&u.Email,
			validation.Required.When(u.ID == 0).Error("The email field is required."),
			is.Email.Error("The email field must be a valid email."),
		),
		validation.Field(&u.Password,
			validation.Required.When(u.ID == 0).Error("The password field is required."),
		),
	)
}

type Login struct {
	Username string
	Password string
}

func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Username,
			validation.Required.Error("The username field is required."),
		),
		validation.Field(&l.Password,
			validation.Required.Error("The password field is required."),
		),
	)
}

type Category struct {
	ID        int          `json:"id" form:"-" db:"id"`
	Name      string       `json:"name" db:"name"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (c Category) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name,
			validation.Required.Error("The name field is required."),
		),
	)
}

type Question struct {
	ID          int          `json:"id" form:"-" db:"id"`
	UserId      int          `json:"user_id" db:"user_id"`
	CategoryId  int          `json:"category_id" db:"category_id"`
	Name        string       `json:"name" db:"name"`
	Title       string       `json:"title" db:"title"`
	Description string       `json:"description" db:"description"`
	PublishedAt sql.NullTime `json:"published_at" db:"published_at"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (q Question) Validate() error {
	return validation.ValidateStruct(&q,
		validation.Field(&q.UserId,
			validation.Required.When(q.ID == 0).Error("The User Id  is required."),
		),
		validation.Field(&q.CategoryId,
			validation.Required.When(q.ID == 0).Error("The CategoryId  is required."),
		),
		validation.Field(&q.Title,
			validation.Required.Error("The Title field is required."),
		),
		validation.Field(&q.Description,
			validation.Required.Error("The Description is required."),
		),
	)
}

type Answere struct {
	ID         int          `json:"id" form:"-" db:"id"`
	UserId     int          `json:"user_id" db:"user_id"`
	QuestionId int          `json:"question_id" db:"question_id"`
	Answere    string       `json:"answere" db:"answere"`
	IsCorrect  bool         `json:"is_correct" db:"is_correct"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (a Answere) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.UserId,
			validation.Required.Error("The User Id  is required."),
		),
		validation.Field(&a.QuestionId,
			validation.Required.Error("The QuestionId  is required."),
		),
		validation.Field(&a.Answere,
			validation.Required.Error("The answere field is required."),
		),
	)
}
