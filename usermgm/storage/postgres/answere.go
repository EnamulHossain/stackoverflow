package postgres

import (
	"fmt"
	"log"
	"stackoverflow/usermgm/storage"
)


const CreateAnswere = `INSERT INTO answere (
	user_id,
	question_id,
	answere
) VALUES(
	:user_id,
	:question_id,
	:answere
) RETURNING *`

func (s PostgresStorage) CreateAnswere(u storage.Answere) (*storage.Answere, error) {
	stmt, err := s.DB.PrepareNamed(CreateAnswere)
	if err != nil {
		fmt.Println("prepared err", err.Error())
		return nil, err
	}

	if err := stmt.Get(&u, u); err != nil {
		fmt.Println("stmr err", err.Error())
		return nil, err
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("unable to insert answere into db")
	}
	fmt.Println("create answere response", u)
	return &u, nil
}




const listanswereQuery = `SELECT * FROM answere WHERE question_id = $1 AND deleted_at IS NULL`

func (s PostgresStorage) ListAnswere(id int32) ([]storage.Answere, error) {

	var answere []storage.Answere
	if err := s.DB.Select(&answere, listanswereQuery,id); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return answere, nil
}



const deleteanswereByIdQuery = `UPDATE answere SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) DeleteAnswere(id int32) error {
	res, err := s.DB.Exec(deleteanswereByIdQuery, id)
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
		return fmt.Errorf("unable to delete answere")
	}

	return nil
}



const getAnswereByID = `SELECT * FROM answere WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetAnswereByID(id int32) (*storage.Answere, error) {
	var u storage.Answere
	if err := s.DB.Get(&u, getAnswereByID, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}





const answereUpdate = `
UPDATE
answere
SET
	answere = :answere

	WHERE id = :id AND deleted_at IS NULL RETURNING *;`

func (s PostgresStorage) AnswereUpdate(u storage.Answere) (*storage.Answere, error) {
	stmt, err := s.DB.PrepareNamed(answereUpdate)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &u, nil
}





const correctanswere = `
UPDATE
answere
SET
	is_correct = :is_correct

	WHERE id = :id AND deleted_at IS NULL RETURNING *;`

func (s PostgresStorage) CorrectAnswere(u storage.Answere) (*storage.Answere, error) {
	stmt, err := s.DB.PrepareNamed(correctanswere)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &u, nil
}
