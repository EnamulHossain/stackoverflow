package question

import (
	"context"
	questionpb "stackoverflow/gunk/v1/question"
	"stackoverflow/usermgm/storage"
)

type CoreQuestion interface {
	CreateQuestion(u storage.Question) (*storage.Question, error)
	ListQuestion() ([]storage.Question, error)
	DeleteQuestion(id int32) error
	GetQuestionByID(id int32) (*storage.Question, error)
	QuestionUpdate(u storage.Question) (*storage.Question, error)
}



type QuestionSvc struct {
	questionpb.UnimplementedQuestionServiceServer
	core CoreQuestion
}


func NewQuestionSvc(cq CoreQuestion) *QuestionSvc {
	return &QuestionSvc{
		core: cq,
	}
}




func (qs QuestionSvc) CreateQuestion(ctx context.Context,r *questionpb.CreateQuestionRequest) (*questionpb.CreateQuestionResponse, error) {
	question := storage.Question{
		CategoryId:  int(r.GetCategoryId()),
		Title:       r.Title,
		Description: r.Description,
	}

	if err := question.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := qs.core.CreateQuestion(question)
	if err != nil {
		return nil, err
	}

	return &questionpb.CreateQuestionResponse{
		Question: &questionpb.Question{
			ID:          int32(u.ID),
			CategoryId:  int32(u.CategoryId),
			Title:       u.Title,
			Description: u.Description,
		},
	}, nil
}

func (qs QuestionSvc) ListQuestion(ctx context.Context,r *questionpb.ListQuestionRequest) (*questionpb.ListQuestionResponse, error) {

	question, err := qs.core.ListQuestion()
	if err != nil {
		return nil, err
	}
	list := make([]*questionpb.Question, len(question))
	for i, q := range question {
		list[i] = &questionpb.Question{
			ID:          int32(q.ID),
			CategoryId:  int32(q.CategoryId),
			Title:       q.Title,
			Description: q.Description,
		}
	}

	return &questionpb.ListQuestionResponse{
		Questions:  list,
	}, nil
}

func (qs QuestionSvc) DeleteQuestion(ctx context.Context, r *questionpb.DeleteQuestionRequest) (*questionpb.DeleteQuestionResponse, error) {
	err := qs.core.DeleteQuestion(r.GetID())

	if err != nil {
		return nil, err
	}
	return &questionpb.DeleteQuestionResponse{}, nil
}

func (qs QuestionSvc) EditQuestion(ctx context.Context, r *questionpb.EditQuestionRequest) (*questionpb.EditQuestionResponse, error) {
	res, err := qs.core.GetQuestionByID(r.GetID())
	if err != nil {
		return nil, err
	}

	return &questionpb.EditQuestionResponse{
		Question: &questionpb.Question{
			ID:          int32(res.ID),
			CategoryId:  int32(res.CategoryId),
			Title:       res.Title,
			Description: res.Description,
		},
	},err

}

func (qs QuestionSvc) UpdateQuestion(ctx context.Context, r *questionpb.UpdateQuestionRequest) (*questionpb.UpdateQuestionResponse, error) {
	question := storage.Question{
		ID:          int(r.GetID()),
		Title:       r.Title,
		Description: r.Description,
	}
	if err := question.Validate(); err != nil {
		return nil, err
	}
	_, err := qs.core.QuestionUpdate(question)
	if err != nil {
		return nil, err
	}

	return &questionpb.UpdateQuestionResponse{},nil
}