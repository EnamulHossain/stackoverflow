package question

import (
	"context"
	"database/sql"
	questionpb "stackoverflow/gunk/v1/question"
	"stackoverflow/usermgm/storage"
)

type CoreQuestion interface {
	CreateQuestion(u storage.Question) (*storage.Question, error)
	// ListQuestion() ([]storage.Question, error)
	ListQuestion(uf storage.UserFilter) ([]storage.Question, error)
	DeleteQuestion(id int32) error
	GetQuestionByID(id int32) (*storage.Question, error)
	QuestionUpdate(u storage.Question) (*storage.Question, error)
	QuestionPublished(u storage.Question) (*storage.Question, error)

	ListQuestionByUser(id int32) ([]storage.Question, error)
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

func (qs QuestionSvc) CreateQuestion(ctx context.Context, r *questionpb.CreateQuestionRequest) (*questionpb.CreateQuestionResponse, error) {
	question := storage.Question{
		UserId:      int(r.GetUserId()),
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
		ID:          int32(u.ID),
		UserId:      int32(u.UserId),
		CategoryId:  int32(u.CategoryId),
		Title:       u.Title,
		Description: u.Description,
	},nil
}

func (qs QuestionSvc) GetQueByUser(ctx context.Context, r *questionpb.GetQueByUserRequest) (*questionpb.GetQueByUserResponse, error) {

	question, err := qs.core.ListQuestionByUser(r.GetID())
	if err != nil {
		return nil, err
	}
	list := make([]*questionpb.Question, len(question))
	for i, q := range question {
		list[i] = &questionpb.Question{
			ID:          int32(q.ID),
			UserId:      int32(q.UserId),
			CategoryId:  int32(q.CategoryId),
			Title:       q.Title,
			Description: q.Description,
		}
	}
	return &questionpb.GetQueByUserResponse{
		Questions: list,
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
			UserId:      int32(res.UserId),
			CategoryId:  int32(res.CategoryId),
			Title:       res.Title,
			Description: res.Description,
		},
	}, err

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

	return &questionpb.UpdateQuestionResponse{}, nil
}

func (qs QuestionSvc) PublishedQuestion(ctx context.Context, r *questionpb.PublishedQuestionRequest) (*questionpb.PublishedQuestionResponse, error) {

	publishQuestion := storage.Question{
		ID: int(r.GetID()),
		PublishedAt: sql.NullTime{
			Time:  r.PublishedAt.AsTime(),
			Valid: false,
		},
	}

	_, err := qs.core.QuestionPublished(publishQuestion)
	if err != nil {
		return nil, err
	}

	return &questionpb.PublishedQuestionResponse{}, nil
}



func (qs QuestionSvc) ListQuestion(ctx context.Context, r *questionpb.ListQuestionRequest) (*questionpb.ListQuestionResponse, error) {

	
	uf := storage.UserFilter{
		SearchTerm: r.GetSearchTerm(),
		Offset:     int(r.GetOffset()),
		Limit:      int(r.GetLimit()),

	}


	question, err := qs.core.ListQuestion(uf)
	if err != nil {
		return nil, err
	}


	list := make([]*questionpb.Question, len(question))
		for i, c := range question {
			list[i] =&questionpb.Question{
				ID:          int32(c.ID),
				UserId:      int32(c.UserId),
				CategoryId:  int32(c.CategoryId),
				Name:        c.Name.String,
				Title:       c.Title,
				Description: c.Description,
			}
		}
		

		return &questionpb.ListQuestionResponse{
			Questions: list,
		},nil
}