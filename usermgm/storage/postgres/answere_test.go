package postgres

import (
	"sort"
	"stackoverflow/usermgm/storage"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestPostgresStorage_CreateAnswere(t *testing.T) {
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

	questionid := int32(question.ID)

	tests := []struct {
		name    string
		in      storage.Answere
		want    *storage.Answere
		wantErr bool
	}{
		{
			name: "Create Answere",
			in: storage.Answere{
				UserId:     int(UserId),
				QuestionId: int(questionid),
				Answere:    "Answere 1",
			},
			want: &storage.Answere{
				UserId:     int(UserId),
				QuestionId: int(questionid),
				Answere:    "Answere 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateAnswere(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateAnswere() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Answere{}, "ID", "IsCorrect", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateQuestion() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestPostgresStorage_ListAnswere(t *testing.T) {
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

	questionid := int32(question.ID)

	answersss := []storage.Answere{
		{
			UserId:     int(UserId),
			QuestionId: int(questionid),
			Answere:    "Answere1",
		},
		{
			UserId:     int(UserId),
			QuestionId: int(questionid),
			Answere:    "Answere2",
		},
		{
			UserId:     int(UserId),
			QuestionId: int(questionid),
			Answere:    "Answere3",
		},
	}

	for _, ans := range answersss {
		_, err := s.CreateAnswere(ans)
		if err != nil {
			t.Fatalf("PostgresStorage.CreateAnswereList() error = %v", err)
		}
	}

	tests := []struct {
		name    string
		in      int32
		inn     storage.UserFilter
		want    []storage.Answere
		wantErr bool
	}{
		{
			name: "List answere By Question",
			in:   int32(questionid),
			inn:  storage.UserFilter{},
			want: []storage.Answere{
				{
					UserId:     int(UserId),
					QuestionId: int(questionid),
					Answere:    "Answere1",
				},
				{
					UserId:     int(UserId),
					QuestionId: int(questionid),
					Answere:    "Answere2",
				},
				{
					UserId:     int(UserId),
					QuestionId: int(questionid),
					Answere:    "Answere3",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.ListAnswere(tt.in, tt.inn)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListAnswere() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Answere{}, "ID", "IsCorrect", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateAnswereList() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestPostgresStorage_DeleteAnswere(t *testing.T) {
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

	questionid := int32(question.ID)

	newans := storage.Answere{
		UserId:     int(UserId),
		QuestionId: int(questionid),
		Answere:    "Answere for delete",
	}
	ans, err := s.CreateAnswere(newans)
	if err != nil {
		t.Fatalf("PostgresStorage.AnswereDelete() error = %v", err)
	}

	id := int32(ans.ID)

	tests := []struct {
		name    string
		in      int32
		wantErr bool
	}{
		{
			name: "Delete Answere success",
			in:   id,
		},
		{
			name:    "Delete Answere Failed",
			in:      id,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.DeleteAnswere(tt.in); (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteAnswere() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresStorage_GetAnswereByID(t *testing.T) {
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

	questionid := int32(question.ID)

	newans := storage.Answere{
		UserId:     int(UserId),
		QuestionId: int(questionid),
		Answere:    "Answere for delete",
	}
	ans, err := s.CreateAnswere(newans)
	if err != nil {
		t.Fatalf("PostgresStorage.AnswereDelete() error = %v", err)
	}

	id := int32(ans.ID)

	tests := []struct {
		name    string
		in      int32
		want    *storage.Answere
		wantErr bool
	}{
		{
			name: "Get_answere_By_id",
			in:   int32(id),
			want: &storage.Answere{
				UserId:     int(UserId),
				QuestionId: int(questionid),
				Answere:    "Answere for delete",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetAnswereByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetAnswereByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Answere{}, "ID", "IsCorrect", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetQuestion() diff = %v", cmp.Diff(got, tt.want, opts...))
			}

		})
	}
}

func TestPostgresStorage_AnswereUpdate(t *testing.T) {
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

	questionid := int32(question.ID)

	newans := storage.Answere{
		UserId:     int(UserId),
		QuestionId: int(questionid),
		Answere:    "Answere for Update",
	}
	ans, err := s.CreateAnswere(newans)
	if err != nil {
		t.Fatalf("PostgresStorage.AnswereDelete() error = %v", err)
	}

	id := int32(ans.ID)

	tests := []struct {
		name    string
		in      *storage.Answere
		wantErr bool
	}{
		{
			name: "UPDATE_Answere_SUCEESS",
			in: &storage.Answere{
				ID:         int(id),
				UserId:     int(UserId),
				QuestionId: int(questionid),
				Answere:    "Answere  Updated",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.AnswereUpdate(*tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.AnswereUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPostgresStorage_CorrectAnswere(t *testing.T) {
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

	questionid := int32(question.ID)

	newans := storage.Answere{
		UserId:     int(UserId),
		QuestionId: int(questionid),
		Answere:    "Answere for Update",
		IsCorrect:  false,
	}
	ans, err := s.CreateAnswere(newans)
	if err != nil {
		t.Fatalf("PostgresStorage.AnswereDelete() error = %v", err)
	}

	id := int32(ans.ID)

	tests := []struct {
		name    string
		in      storage.Answere
		wantErr bool
	}{
		{
			name: "Answere Is Correct",
			in: storage.Answere{
				ID: int(id),
				IsCorrect:  true,
			},wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.CorrectAnswere(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CorrectAnswere() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
