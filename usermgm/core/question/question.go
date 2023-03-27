package question

import (
	"fmt"
	"stackoverflow/usermgm/storage"
)

type QuestionStore interface {
	CreateQuestion(u storage.Question) (*storage.Question, error)
	// ListQuestion() ([]storage.Question, error)
	ListQuestion(uf storage.UserFilter) ([]storage.Question, error)
	DeleteQuestion(id int32) error
	GetQuestionByID(id int32) (*storage.Question, error)
	QuestionUpdate(u storage.Question) (*storage.Question, error)
	ListQuestionByUser(id int32) ([]storage.Question, error)

	QuestionPublished(u storage.Question) (*storage.Question, error)
}

type CoreQuestion struct {
	store QuestionStore
}

func NewCoreQuestion(qs QuestionStore) *CoreQuestion {
	return &CoreQuestion{
		store: qs,
	}
}

func (cq CoreQuestion) CreateQuestion(u storage.Question) (*storage.Question, error) {

	cc, err := cq.store.CreateQuestion(u)
	if err != nil {
		return nil, err
	}
	if cc == nil {
		return nil, fmt.Errorf("unable to register")
	}
	return cc, nil
}

func (cq CoreQuestion) ListQuestion(uf storage.UserFilter) ([]storage.Question, error) {

	list, err := cq.store.ListQuestion(uf)
	if err != nil {
		return nil, err
	}
	// fmt.Println("########## q list",list)
	
	return list, nil
}

func (cq CoreQuestion) ListQuestionByUser(id int32) ([]storage.Question, error){
	list, err := cq.store.ListQuestionByUser(id)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (cq CoreQuestion) DeleteQuestion(id int32) error {
	err := cq.store.DeleteQuestion(id)

	if err != nil {
		return err
	}

	return nil
}

func (cq CoreQuestion) GetQuestionByID(id int32) (*storage.Question, error) {

	question, err := cq.store.GetQuestionByID(id)
	if err != nil {
		return question, err
	}

	return question, err
}

func (cq CoreQuestion) QuestionUpdate(u storage.Question) (*storage.Question, error) {

	question, err := cq.store.QuestionUpdate(u)
	if err != nil {
		return nil, err
	}

	if question == nil {
		return nil, fmt.Errorf("unable to update")
	}

	return question, err
}


func (cq CoreQuestion) QuestionPublished(u storage.Question) (*storage.Question, error) {


	questions, err := cq.store.QuestionPublished(u)
	if err != nil {
		return nil, err
	}

	if questions == nil {
		return nil, fmt.Errorf("unable to update")
	}

	return questions, err


}
