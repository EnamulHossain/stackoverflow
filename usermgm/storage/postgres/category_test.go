package postgres

import (
	"sort"
	"stackoverflow/usermgm/storage"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestPostgresStorage_CreateCategory(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	tests := []struct {
		name    string
		in      storage.Category
		want    *storage.Category
		wantErr bool
	}{
		{
			name: "Create_Category_Success",
			in: storage.Category{
				Name: "Math1",
			},
			want: &storage.Category{
				Name: "Math1",
			},
		},
		{
			name: "Create_Unique_Category",
			in: storage.Category{
				Name: "Math1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateCategory(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateCategory() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestPostgresStorage_DeleteCategory(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCat := storage.Category{
		Name: "Phy",
	}
	category, err := s.CreateCategory(newCat)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}

	id := int32(category.ID)

	tests := []struct {
		name    string
		in      int32
		wantErr bool
	}{
		{
			name: "DELETE_Category_BY_ID_SUCEESS",
			in:   id,
		},
		{
			name:    "DELETE_Category_BY_ID_FAILED",
			in:      id,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.DeleteCategory(tt.in); (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresStorage_GetCategoryByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCat := storage.Category{
		ID:   1,
		Name: "Phy",
	}
	category, err := s.CreateCategory(newCat)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	id := int32(category.ID)

	tests := []struct {
		name    string
		in      int32
		want    *storage.Category
		wantErr bool
	}{
		{
			name: "Get_Category_BY_ID_SUCEESS",
			in:   id,
			want: &storage.Category{
				ID:   1,
				Name: "Phy",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetCategoryByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetCategoryByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Category{}, "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetCategoryByID() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestPostgresStorage_UpdateCategory(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	newCat := storage.Category{
		Name: "Phy",
	}
	cate, err := s.CreateCategory(newCat)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUser() error = %v", err)
	}
	id := cate.ID

	tests := []struct {
		name    string
		in      storage.Category
		wantErr bool
	}{
		{
			name: "UPDATE_Category_SUCEESS",
			in: storage.Category{
				ID:   id,
				Name: "PhyUpdate"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := s.UpdateCategory(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestPostgresStorage_ListCategory(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	categories := []storage.Category{
		{
			Name: "Math",
		},
		{
			Name: "Ict",
		},
		{
			Name: "Phy",
		},
	}

	for _, cat := range categories {
		_, err := s.CreateCategory(cat)
		if err != nil {
			t.Fatalf("unable to create  list user testing %v", err)
		}
	}
	tests := []struct {
		name    string
		in      storage.UserFilter
		want    []storage.Category
		wantErr bool
	}{
		{
			name: "List Category Success",
			in: storage.UserFilter{},
			want: []storage.Category{
				{
					Name: "Math",
				},
				{
					Name: "Ict",
				},
				{
					Name: "Phy",
				},
				
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListCategory(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.ListCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}
			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.ListCategory() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
