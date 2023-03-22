package category

import (
	"context"
	categorypb "stackoverflow/gunk/v1/category"
	"stackoverflow/usermgm/storage"
)

type CoreCategory interface {
	CreateCategory(u storage.Category) (*storage.Category, error)
	// ListCategory() ([]storage.Category, error)
	ListCategory(uf storage.UserFilter) ([]storage.Category, error)
	DeleteCategory(id int32) error
	GetCategoryByID(id int32) (*storage.Category, error)
	UpdateCategory(u storage.Category) (*storage.Category, error)
}

type CategorySvc struct {
	categorypb.UnimplementedCategoryServiceServer
	core CoreCategory
}

func NewCatrgorySvc(cc CoreCategory) *CategorySvc {
	return &CategorySvc{
		core: cc,
	}
}

func (cs CategorySvc) CreateCategory(ctx context.Context, r *categorypb.CreateCategoryRequest) (*categorypb.CreateCategoryResponse, error) {
	category := storage.Category{
		Name: r.GetName(),
	}

	if err := category.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := cs.core.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	return &categorypb.CreateCategoryResponse{
		Category: &categorypb.Category{
			ID:   int32(u.ID),
			Name: u.Name,
		},
	}, nil
}

func (cs CategorySvc) ListCategory(ctx context.Context, r *categorypb.ListCategoryRequest) (*categorypb.ListCategoryResponse, error) {
	uf := storage.UserFilter{
		SearchTerm: r.GetSearchTerm(),
		Offset:     int(r.GetOffset()),
		Limit:      int(r.GetLimit()),
	}
	categories, err := cs.core.ListCategory(uf)
	if err != nil {
		return nil, err
	}
	list := make([]*categorypb.Category, len(categories))
	for i, c := range categories {
		list[i] = &categorypb.Category{
			ID:   int32(c.ID),
			Name: c.Name,
		}
	}

	return &categorypb.ListCategoryResponse{
		Categorys: list,
	}, nil

}

func (cs CategorySvc) DeleteCategory(ctx context.Context, r *categorypb.DeleteCategoryRequest) (*categorypb.DeleteCategoryResponse, error) {
	err := cs.core.DeleteCategory(r.GetID())

	if err != nil {
		return nil, err
	}
	return &categorypb.DeleteCategoryResponse{}, nil
}

func (cs CategorySvc) EditCategory(ctx context.Context, r *categorypb.EditCategoryRequest) (*categorypb.EditCategoryResponse, error) {
	res, err := cs.core.GetCategoryByID(r.GetID())
	if err != nil {
		return nil, err
	}

	return &categorypb.EditCategoryResponse{
		Category: &categorypb.Category{
			ID: int32(res.ID),
			Name: res.Name,
		},
	},err
}



func (cs CategorySvc) UpdateCategory(ctx context.Context, r *categorypb.UpdateCategoryRequest) (*categorypb.UpdateCategoryResponce, error) {
	category := storage.Category{
		ID:        int(r.GetID()),
		Name:      r.GetName(),
	}
	if err := category.Validate(); err != nil {
		return nil, err
	}
	_, err := cs.core.UpdateCategory(category)
	if err != nil {
		return nil, err
	}

	return &categorypb.UpdateCategoryResponce{},nil
}
