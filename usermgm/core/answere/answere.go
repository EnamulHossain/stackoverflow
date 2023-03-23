package answere

import (
	"fmt"
	"stackoverflow/usermgm/storage"
)

type AnswereStore interface {
	CreateAnswere(u storage.Answere) (*storage.Answere, error)
	// ListAnswere(id int32) ([]storage.Answere, error)
	ListAnswere(id int32,uf storage.UserFilter) ([]storage.Answere, error)
	DeleteAnswere(id int32) error
	GetAnswereByID(id int32) (*storage.Answere, error)
	AnswereUpdate(u storage.Answere) (*storage.Answere, error)

	CorrectAnswere(u storage.Answere) (*storage.Answere, error)
}

type CoreAnswere struct {
	store AnswereStore
}

func NewCoreAnswere(as AnswereStore) *CoreAnswere {
	return &CoreAnswere{
		store: as,
	}
}


func (ca CoreAnswere) CreateAnswere(u storage.Answere) (*storage.Answere, error) {
	cc, err := ca.store.CreateAnswere(u)
	if err != nil {
		return nil, err
	}
	if cc == nil {
		return nil, fmt.Errorf("unable to Answere")
	}
	return cc, nil
}


func (ca CoreAnswere) ListAnswere(id int32,uf storage.UserFilter) ([]storage.Answere, error) {
	list, err := ca.store.ListAnswere(id,uf)
	if err != nil {
		return nil, err
	}
	return list, nil
}


func (ca CoreAnswere) DeleteAnswere(id int32) error {
	err := ca.store.DeleteAnswere(id)

	if err != nil {
		return err
	}

	return nil
}


func (ca CoreAnswere) GetAnswereByID(id int32) (*storage.Answere, error){
	answere, err := ca.store.GetAnswereByID(id)
	if err != nil {
		return answere, err
	}

	return answere, err
}


func (ca CoreAnswere) AnswereUpdate(u storage.Answere) (*storage.Answere, error) {
	answere, err := ca.store.AnswereUpdate(u)
	if err != nil {
		return nil, err
	}

	if answere == nil {
		return nil, fmt.Errorf("unable to update")
	}

	return answere, err
}



func (ca CoreAnswere) CorrectAnswere(u storage.Answere) (*storage.Answere, error) {


	answere, err := ca.store.CorrectAnswere(u)
	if err != nil {
		return nil, err
	}

	if answere == nil {
		return nil, fmt.Errorf("unable to update")
	}

	return answere, err


}