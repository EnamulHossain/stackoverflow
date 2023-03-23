package postgres

import (
	"sort"
	"stackoverflow/usermgm/storage"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestRegister(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "Register_User_Success",
			in: storage.User{
				FirstName: "first",
				LastName:  "last",
				Email:     "test@example.com",
				Username:  "test",
				Password:  "12345678",
			},
			want: &storage.User{
				FirstName: "first",
				LastName:  "last",
				Email:     "test@example.com",
				Username:  "test",
				IsActive:  true,
				IsAdmin:   false,
			},
		},
		{
			name: "Register_User_EMIAL_UNIQUE_FAILED",
			in: storage.User{
				FirstName: "first2",
				LastName:  "last2",
				Email:     "test@example.com",
				Username:  "test2",
				Password:  "12345678",
			},
			wantErr: true,
		},
		{
			name: "Register_User_USRNAME_UNIQUE_FAILED",
			in: storage.User{
				FirstName: "first3",
				LastName:  "last3",
				Email:     "test3@example.com",
				Username:  "test",
				Password:  "12345678",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.Register(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestPostgresStorage_AdminRegister(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{

		{
			name: "Register_Admin_SUCESES",
			in: storage.User{
				FirstName: "firstf",
				LastName:  "last",
				Email:     "testf@example.com",
				Username:  "testf",
				Password:  "12345678",
			},
			want: &storage.User{
				FirstName: "firstf",
				LastName:  "last",
				Email:     "testf@example.com",
				Username:  "testf",
				IsActive:  true,
				IsAdmin:   true,
			},
		},
		{
			name: "Register_Admin_EMIAL_UNIQUE_FAILED",
			in: storage.User{
				FirstName: "firstf2",
				LastName:  "lastf2",
				Email:     "testf@example.com",
				Username:  "testf2",
				Password:  "12345678",
			},
			wantErr: true,
		},
		{
			name: "Register_Admin_USRNAME_UNIQUE_FAILED",
			in: storage.User{
				FirstName: "first3",
				LastName:  "last3",
				Email:     "test3@example.com",
				Username:  "testf",
				Password:  "12345678",
				IsActive:  true,
				IsAdmin:   true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.AdminRegister(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.AdminRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}


func TestPostgresStorage_GetUserByUsername(t *testing.T) {
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
	username := createUser.Username

	tests := []struct {
		name    string
		in      string
		want    *storage.User
		wantErr bool
	}{
		{
			name: "Get_User_By_UserName_SUCESES",
			in:   username,
			want: &storage.User{
				FirstName: "Sayeem",
				LastName:  "Mahmud",
				Email:     "Sayeem@gmail.com",
				Username:  "Sayeem",
				IsActive:  true,
				IsAdmin:   false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetUserByUsername(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}



func TestPostgresStorage_ListUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	users := []storage.User{
		{
			FirstName: "jabbar",
			LastName:  "khan",
			Email:     "jabbar@example.com",
			Username:  "jabbar",
			Password:  "12345678",
		},
		{
			FirstName: "ratul",
			LastName:  "khan",
			Email:     "ratul@example.com",
			Username:  "ratul",
			Password:  "12345678",
		},
		{
			FirstName: "pranto",
			LastName:  "khan",
			Email:     "pranto@example.com",
			Username:  "pranto",
			Password:  "12345678",
		},
	}

	for _, user := range users {
		_, err := s.Register(user)
		if err != nil {
			t.Fatalf("unable to create user for list user testing %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.UserFilter
		want    []storage.User
		wantErr bool
	}{
		{
			name: "LIST_ALL_USER_SUCCESS",
			in:   storage.UserFilter{},
			want: []storage.User{
				{
					FirstName: "jabbar",
					LastName:  "khan",
					Email:     "jabbar@example.com",
					Username:  "jabbar",
					IsActive:  true,
					IsAdmin:   false,
				},
				{
					FirstName: "ratul",
					LastName:  "khan",
					Email:     "ratul@example.com",
					Username:  "ratul",
					IsActive:  true,
					IsAdmin:   false,
				},
				{
					FirstName: "pranto",
					LastName:  "khan",
					Email:     "pranto@example.com",
					Username:  "pranto",
					IsActive:  true,
					IsAdmin:   false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}



func TestPostgresStorage_DeleteUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newUser := storage.User{
		FirstName: "karim",
		LastName:  "khan",
		Email:     "karim@example.com",
		Username:  "karim",
		Password:  "12345678",
	}

	user, err := s.Register(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	id := int32(user.ID)

	tests := []struct {
		name    string
		in      int32
		wantErr bool
	}{
		{
			name: "DELETE_USER_BY_ID_SUCEESS",
			in:   id,
		},
		{
			name:    "DELETE_USER_BY_ID_FAILED",
			in:      id,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}


func TestPostgresStorage_GetUserByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newUser := storage.User{
		FirstName: "Rahrim",
		LastName:  "khan",
		Email:     "rahrim@example.com",
		Username:  "Rahrim",
		Password:  "12345678",
	}

	user, err := s.Register(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	id := int32(user.ID)

	tests := []struct {
		name    string
		in      int32
		want    *storage.User
		wantErr bool
	}{
		{
			name: "Get_USER_BY_ID_SUCEESS",
			in:   id,
			want: &storage.User{
				FirstName: "Rahrim",
				LastName:  "khan",
				Email:     "rahrim@example.com",
				Username:  "Rahrim",
				Password:  "12345678",
				IsActive:  true,
				IsAdmin:   false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.GetUserByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}


func TestPostgresStorage_UpdateUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newUser := storage.User{
		FirstName: "rahim",
		LastName:  "khan",
		Email:     "rahim@example.com",
		Username:  "rahim",
		Password:  "12345678",
	}
	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "UPDATE_USER_SUCEESS",
			in: storage.User{
				ID:        0,
				FirstName: "rahimupdate",
				LastName:  "khanupdate",
				Email:     "rahim@example.com",
				Username:  "rahim",
			},
			want: &storage.User{
				FirstName: "rahimupdate",
				LastName:  "khanupdate",
				Email:     "rahim@example.com",
				Username:  "rahim",
			},
		},
	}

	user, err := s.Register(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.ID = user.ID
			got, err := s.UpdateUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
