package postgres

import (
	"database/sql"
	"sort"
	"stackoverflow/usermgm/storage"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestPostgresStorage_CreateQuestion(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	user := storage.User{
		FirstName: "Sayeem",
		LastName:  "Mahmud",
		Email:     "Sayeem@gmail.com",
		Username:  "Sayeem",
		Password:  "12345678",
	}

	createUser, err := s.Register(user)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	UserId := int32(createUser.ID)

	newCat := storage.Category{
		Name: "Phy",
	}
	category, err := s.CreateCategory(newCat)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	catid := int32(category.ID)

	tests := []struct {
		name    string
		in      storage.Question
		want    *storage.Question
		wantErr bool
	}{
		{name: "Create Question",
			in: storage.Question{
				UserId:      int(UserId),
				CategoryId:  int(catid),
				Title:       "Question Title",
				Description: "Question Description",
			},
			want: &storage.Question{
				UserId:      int(UserId),
				CategoryId:  int(catid),
				Title:       "Question Title",
				Description: "Question Description",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateQuestion(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Question{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateQuestion() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestPostgresStorage_DeleteQuestion(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	user := storage.User{
		FirstName: "Sayeem",
		LastName:  "Mahmud",
		Email:     "Sayeem@gmail.com",
		Username:  "Sayeem",
		Password:  "12345678",
	}

	createUser, err := s.Register(user)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	UserId := int32(createUser.ID)

	newCat := storage.Category{
		Name: "Phy",
	}
	category, err := s.CreateCategory(newCat)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	catid := int32(category.ID)

	newquestion := storage.Question{
		UserId:      int(UserId),
		CategoryId:  int(catid),
		Title:       "Question Title",
		Description: "Question Description",
	}

	question, err := s.CreateQuestion(newquestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuestion() error = %v", err)
	}

	id := int32(question.ID)

	tests := []struct {
		name    string
		in      int32
		wantErr bool
	}{
		{
			name: "Delete Question success",
			in:   id,
		},
		{
			name:    "Delete Question Failed",
			in:      id,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.DeleteQuestion(tt.in); (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteQuestion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresStorage_GetQuestionByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	user := storage.User{
		FirstName: "Sayeem",
		LastName:  "Mahmud",
		Email:     "Sayeem@gmail.com",
		Username:  "Sayeem",
		Password:  "12345678",
	}

	createUser, err := s.Register(user)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	UserId := int32(createUser.ID)

	newCat := storage.Category{
		Name: "Phy",
	}
	category, err := s.CreateCategory(newCat)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	catid := int32(category.ID)

	newquestion := storage.Question{
		UserId:      int(UserId),
		CategoryId:  int(catid),
		Title:       "Question Title",
		Description: "Question Description",
	}

	question, err := s.CreateQuestion(newquestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuestion() error = %v", err)
	}

	id := int32(question.ID)

	tests := []struct {
		name    string
		in      int32
		want    *storage.Question
		wantErr bool
	}{
		{
			name: "Get Question By ID",
			in:   id,
			want: &storage.Question{
				UserId:      int(UserId),
				CategoryId:  int(catid),
				Title:       "Question Title",
				Description: "Question Description",
			}, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetQuestionByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetQuestionByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Question{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateQuestion() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestPostgresStorage_QuestionUpdate(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	user := storage.User{
		FirstName: "Sayeem",
		LastName:  "Mahmud",
		Email:     "Sayeem@gmail.com",
		Username:  "Sayeem",
		Password:  "12345678",
	}

	createUser, err := s.Register(user)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	UserId := int32(createUser.ID)

	newCat := storage.Category{
		Name: "Phy",
	}
	category, err := s.CreateCategory(newCat)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	catid := int32(category.ID)

	newquestion := storage.Question{
		UserId:      int(UserId),
		CategoryId:  int(catid),
		Title:       "Question Title",
		Description: "Question Description",
	}

	question, err := s.CreateQuestion(newquestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuestion() error = %v", err)
	}

	id := int32(question.ID)

	tests := []struct {
		name    string
		in      storage.Question
		wantErr bool
	}{
		{
			name: "UPDATE_Question_SUCEESS",
			in: storage.Question{
				ID:          int(id),
				UserId:      int(UserId),
				CategoryId:  int(catid),
				Title:       "Question Title Update",
				Description: "Question Description Update",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.QuestionUpdate(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.QuestionUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPostgresStorage_QuestionPublished(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	user := storage.User{
		FirstName: "Sayeem",
		LastName:  "Mahmud",
		Email:     "Sayeem@gmail.com",
		Username:  "Sayeem",
		Password:  "12345678",
	}

	createUser, err := s.Register(user)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	UserId := int32(createUser.ID)

	newCat := storage.Category{
		Name: "Phy",
	}
	category, err := s.CreateCategory(newCat)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	catid := int32(category.ID)

	newquestion := storage.Question{
		UserId:      int(UserId),
		CategoryId:  int(catid),
		Title:       "Question Title",
		Description: "Question Description",
	}

	question, err := s.CreateQuestion(newquestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuestion() error = %v", err)
	}

	id := int32(question.ID)

	tests := []struct {
		name    string
		in      storage.Question
		wantErr bool
	}{
		{
			name: "UPDATE_Question_SUCEESS",
			in: storage.Question{
				ID:          int(id),
				PublishedAt: sql.NullTime{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.QuestionPublished(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.QuestionPublished() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}


func TestPostgresStorage_ListQuestionByUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	user := storage.User{
		FirstName: "Sayeem",
		LastName:  "Mahmud",
		Email:     "Sayeem@gmail.com",
		Username:  "Sayeem",
		Password:  "12345678",
	}

	createUser, err := s.Register(user)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	UserId := int32(createUser.ID)

	newCat := storage.Category{
		Name: "Phy",
	}
	category, err := s.CreateCategory(newCat)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	catid := int32(category.ID)

	newquestions := []storage.Question{
		{
			UserId:      int(UserId),
			CategoryId:  int(catid),
			Title:       "Question Title",
			Description: "Question Description",
		},
		{
			UserId:      int(UserId),
			CategoryId:  int(catid),
			Title:       "Question Title1",
			Description: "Question Description1",
		},
		{
			UserId:      int(UserId),
			CategoryId:  int(catid),
			Title:       "Question Title2",
			Description: "Question Description2",
		},
	}

	for _, que := range newquestions {
		_, err := s.CreateQuestion(que)
		if err != nil {
			t.Fatalf("PostgresStorage.CreateQuestionList() error = %v", err)
		}
	}

	tests := []struct {
		name    string
		in      int32
		want    []storage.Question
		wantErr bool
	}{
		{
			name: "List_Question_By_User",
			in:   int32(UserId),
			want: []storage.Question{
				{
					UserId:      int(UserId),
					CategoryId:  int(catid),
					Title:       "Question Title",
					Description: "Question Description",
				},
				{
					UserId:      int(UserId),
					CategoryId:  int(catid),
					Title:       "Question Title1",
					Description: "Question Description1",
				},
				{
					UserId:      int(UserId),
					CategoryId:  int(catid),
					Title:       "Question Title2",
					Description: "Question Description2",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.ListQuestionByUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListQuestionByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Question{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.ListQuestionByUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}







func TestPostgresStorage_ListQuestion(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	user := storage.User{
		FirstName: "Sayeem",
		LastName:  "Mahmud",
		Email:     "Sayeem@gmail.com",
		Username:  "Sayeem",
		Password:  "12345678",
	}

	createUser, err := s.Register(user)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	UserId := int32(createUser.ID)

	newCat := storage.Category{
		Name: "Phy",
	}
	category, err := s.CreateCategory(newCat)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	catid := int32(category.ID)
	catname := category.Name

	newquestions := []storage.Question{
		{
			UserId: int(UserId),
			CategoryId: int(catid),
			Name:  catname,
			Title: "Question Title",
			Description: "Question Description",
		},
		{
			UserId:      int(UserId),
			CategoryId:  int(catid),
			Title:       "Question Title1",
			Description: "Question Description1",
		},
		{
			UserId:      int(UserId),
			CategoryId:  int(catid),
			Title:       "Question Title2",
			Description: "Question Description2",
		},
	}

	for _, que := range newquestions {
		_, err := s.CreateQuestion(que)
		if err != nil {
			t.Fatalf("PostgresStorage.CreateQuestionList() error = %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.UserFilter
		want    []storage.Question
		wantErr bool
	}{
		{
			name: "Question List",
			in:   storage.UserFilter{
				SearchTerm: "",
				Offset:     0,
				Limit:      10,
			},
			want: []storage.Question{
				{
					UserId: int(UserId),
					CategoryId: int(catid),
					Name:  catname,
					Title: "Question Title",
					Description: "Question Description",
				},
				{
					UserId: int(UserId),
					CategoryId: int(catid),
					Name:  catname,
					Title: "Question Title1",
					Description: "Question Description1",
				},
				{
					UserId:      int(UserId),
					CategoryId:  int(catid),
					Name:  catname,
					Title:       "Question Title2",
					Description: "Question Description2",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListQuestion(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Question{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateQuestion() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}