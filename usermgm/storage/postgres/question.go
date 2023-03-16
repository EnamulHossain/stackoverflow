package postgres

import (
	"fmt"
	"log"
	"stackoverflow/usermgm/storage"
)



const CreateQuestion = `INSERT INTO question (
	user_id,
	category_id,
	title,
	description
) VALUES(
	:user_id,
	:category_id,
	:title,
	:description
) RETURNING *`

func (s PostgresStorage) CreateQuestion(u storage.Question) (*storage.Question, error) {
	stmt, err := s.DB.PrepareNamed(CreateQuestion)
	if err != nil {
		fmt.Println("prepared err", err.Error())
		return nil, err
	}

	if err := stmt.Get(&u, u); err != nil {
		fmt.Println("stmr err", err.Error())
		return nil, err
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("unable to insert category into db")
	}
	fmt.Println("create question response", u)
	return &u, nil
}


const listcqueQuery = `SELECT * FROM question WHERE deleted_at IS NULL`

func (s PostgresStorage) ListQuestion() ([]storage.Question, error) {

	var question []storage.Question
	if err := s.DB.Select(&question, listcqueQuery); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return question, nil
}




const deleteQuestionByIdQuery = `UPDATE question SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) DeleteQuestion(id int32) error {
	res, err := s.DB.Exec(deleteQuestionByIdQuery, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println("storage err", err)
		return err
	}

	if rowCount <= 0 {
		return fmt.Errorf("unable to delete category")
	}

	return nil
}



const getQuestionByID = `SELECT * FROM question WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetQuestionByID(id int32) (*storage.Question, error) {
	var u storage.Question
	if err := s.DB.Get(&u, getQuestionByID, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}




const QuestionUpdate = `
UPDATE
question
SET
	title = :title,
	description = :description

	WHERE id = :id AND deleted_at IS NULL RETURNING *;`

func (s PostgresStorage) QuestionUpdate(u storage.Question) (*storage.Question, error) {
	stmt, err := s.DB.PrepareNamed(QuestionUpdate)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &u, nil
}
