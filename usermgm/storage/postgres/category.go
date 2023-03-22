package postgres

import (
	"fmt"
	"log"
	"stackoverflow/usermgm/storage"
)


const createCategory = `INSERT INTO category (
	name
) VALUES(
	:name
) RETURNING *`

func (s PostgresStorage) CreateCategory(u storage.Category) (*storage.Category, error) {
	stmt, err := s.DB.PrepareNamed(createCategory)
	if err != nil {
		return nil, err
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("unable to insert category into db")
	}
	return &u, nil
}




// const listcatQuery = `SELECT * FROM category WHERE deleted_at IS NULL`
const listcatQuery = `SELECT * FROM category WHERE deleted_at IS NULL
					AND (name ILIKE '%%' || $1 || '%%')
					ORDER BY id ASC
						OFFSET $2
						LIMIT $3`

func (s PostgresStorage) ListCategory(uf storage.UserFilter) ([]storage.Category, error) {
	var category []storage.Category
	if uf.Limit == 0 {
		uf.Limit = 15
	}
	if err := s.DB.Select(&category, listcatQuery, uf.SearchTerm, uf.Offset, uf.Limit); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return category, nil
}




const deleteCategoryByIdQuery = `UPDATE category SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) DeleteCategory(id int32) error {
	res, err := s.DB.Exec(deleteCategoryByIdQuery, id)
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




const getCategoryByID = `SELECT * FROM category WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetCategoryByID(id int32) (*storage.Category, error) {
	var u storage.Category
	if err := s.DB.Get(&u, getCategoryByID, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}




const CategoryUpdate = `
UPDATE
category
SET
	name = :name
	WHERE id = :id AND deleted_at IS NULL RETURNING *;`

func (s PostgresStorage) UpdateCategory(u storage.Category) (*storage.Category, error) {
	stmt, err := s.DB.PrepareNamed(CategoryUpdate)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &u, nil
}
