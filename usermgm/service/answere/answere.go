package answere

import (
	// "context"
	"context"
	answerepb "stackoverflow/gunk/v1/answere"
	"stackoverflow/usermgm/storage"
)

type CoreAnswere interface {
	CreateAnswere(u storage.Answere) (*storage.Answere, error)
	ListAnswere(id int32) ([]storage.Answere, error)
	DeleteAnswere(id int32) error
	GetAnswereByID(id int32) (*storage.Answere, error)
	AnswereUpdate(u storage.Answere) (*storage.Answere, error)

	CorrectAnswere(u storage.Answere) (*storage.Answere, error)
}

type AnswereSvc struct {
	answerepb.UnimplementedAnswereServiceServer
	core CoreAnswere
}

func NewAnswereSvc(ca CoreAnswere) *AnswereSvc {
	return &AnswereSvc{
		core: ca,
	}
}

func (as AnswereSvc) AnswereCreate(ctx context.Context, r *answerepb.AnswereCreateRequest) (*answerepb.AnswereCreateResponse, error) {
	answere := storage.Answere{
		UserId:     int(r.GetUserId()),
		QuestionId: int(r.GetQuestionId()),
		Answere:    r.Answere,
	}

	if err := answere.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := as.core.CreateAnswere(answere)
	if err != nil {
		return nil, err
	}

	return &answerepb.AnswereCreateResponse{
		Answere: &answerepb.Answere{
			ID:         int32(u.ID),
			UserId:     r.UserId,
			QuestionId: int32(u.QuestionId),
			Answere:    u.Answere,
			IsCorrect:  false,
		},
	}, err
}

func (as AnswereSvc) AnswereList(ctx context.Context, r *answerepb.AnswereListRequest) (*answerepb.AnswereListResponse, error) {


	answere, err := as.core.ListAnswere(r.GetID())
	if err != nil {
		return nil, err
	}

	list := make([]*answerepb.Answere, len(answere))
	for i, q := range answere {
		list[i] = &answerepb.Answere{
			ID:         int32(q.ID),
			UserId:     int32(q.UserId),
			QuestionId: int32(q.QuestionId),
			Answere:    q.Answere,
			IsCorrect:  q.IsCorrect,
		}
	}

	return &answerepb.AnswereListResponse{
		Answere: list,
	}, nil
}

func (as AnswereSvc) AnswereDelete(ctx context.Context, r *answerepb.AnswereDeleteRequest) (*answerepb.AnswereDeleteResponse, error) {
	err := as.core.DeleteAnswere(r.GetID())

	if err != nil {
		return nil, err
	}
	return &answerepb.AnswereDeleteResponse{}, nil
}

func (as AnswereSvc) AnswereEdit(ctx context.Context, r *answerepb.AnswereEditRequest) (*answerepb.AnswereEditResponse, error) {
	res, err := as.core.GetAnswereByID(r.GetID())
	if err != nil {
		return nil, err
	}

	return &answerepb.AnswereEditResponse{
		Answere: &answerepb.Answere{
			ID:         int32(res.ID),
			UserId:     int32(res.UserId),
			QuestionId: int32(res.QuestionId),
			Answere:    res.Answere,
		},
	}, err
}

func (as AnswereSvc) AnswereUpdate(ctx context.Context, r *answerepb.AnswereUpdateRequest) (*answerepb.AnswereUpdateResponse, error) {
	answere := storage.Answere{
		ID:      int(r.GetID()),
		Answere: r.Answere,
	}
	// if err := answere.Validate(); err != nil {
	// 	return nil, err
	// }
	_, err := as.core.AnswereUpdate(answere)
	if err != nil {
		return nil, err
	}

	return &answerepb.AnswereUpdateResponse{}, nil
}

func (as AnswereSvc) CorrectAnswere(ctx context.Context, r *answerepb.CorrectAnswereRequest) (*answerepb.CorrectAnswereResponse, error) {

	cans := storage.Answere{
		ID:        int(r.GetID()),
		IsCorrect: false,
	}

	corretans := storage.Answere{
		ID:        int(r.GetID()),
		IsCorrect: true,
	}

	if r.IsCorrect == true {
		_, err := as.core.CorrectAnswere(cans)
		if err != nil {
			return nil, err
		}
	}

	if r.IsCorrect == false {
		_, err := as.core.CorrectAnswere(corretans)
		if err != nil {
			return nil, err
		}
	}

	return &answerepb.CorrectAnswereResponse{}, nil
}
