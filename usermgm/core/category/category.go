package category

import (
	"fmt"
	"stackoverflow/usermgm/storage"
)

type CategoryStore interface {
	CreateCategory(storage.Category) (*storage.Category, error)
	// ListCategory() ([]storage.Category, error)
	ListCategory(uf storage.UserFilter) ([]storage.Category, error)
	DeleteCategory(id int32) error
	GetCategoryByID(id int32) (*storage.Category, error)
	UpdateCategory(u storage.Category) (*storage.Category, error)
}

type CoreCategory struct {
	store CategoryStore
}

func NewCoreCategory(us CategoryStore) *CoreCategory {
	return &CoreCategory{
		store: us,
	}
}

func (ca CoreCategory) CreateCategory(u storage.Category) (*storage.Category, error) {

	cc, err := ca.store.CreateCategory(u)
	if err != nil {
		return nil, err
	}
	if cc == nil {
		return nil, fmt.Errorf("unable to register")
	}
	return cc, nil
}

func (cc CoreCategory) ListCategory(uf storage.UserFilter) ([]storage.Category, error) {

	
	list, err := cc.store.ListCategory(uf)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (cc CoreCategory) DeleteCategory(id int32) error {
	err := cc.store.DeleteCategory(id)

	if err != nil {
		return err
	}

	return nil
}

func (cu CoreCategory) GetCategoryByID(id int32) (*storage.Category, error) {
	
	categories, err := cu.store.GetCategoryByID(id)
	if err != nil {
		return categories, err
	}

	return categories, err
}

func (cu CoreCategory) UpdateCategory(u storage.Category) (*storage.Category, error) {

	categories, err := cu.store.UpdateCategory(u)
	if err != nil {
		return nil, err
	}

	if categories == nil {
		return nil, fmt.Errorf("unable to update")
	}

	return categories, err

}
