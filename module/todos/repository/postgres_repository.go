package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/module/todos"
	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/module/todos/model"
)

// postgresTodosRepository struct
type postgresTodosRepository struct {
	db *sql.DB
}

// NewTodosRepositoryPostgres func
func NewTodosRepositoryPostgres(db *sql.DB) todos.Repository {
	return &postgresTodosRepository{db}
}

// Save Todo
func (r *postgresTodosRepository) Save(mt *model.Todo) error {

	query := `
	INSERT INTO tbl_todos 
	(
		todo_name,
		is_done,
		created_at,
		updated_at
	)
	VALUES ($1, $2, $3, $4)
	RETURNING todo_id`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	err = statement.QueryRow(mt.TodoName, mt.IsDone, mt.CreatedAt, mt.UpdatedAt).Scan(&mt.TodoID)

	if err != nil {
		return err
	}

	return nil
}

// FindByID Todo
func (r *postgresTodosRepository) FindByID(id string) (*model.Todo, error) {

	query := `
	SELECT *
	FROM tbl_todos WHERE todo_id = $1
	`

	var mt model.Todo

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&mt.TodoID, &mt.TodoName, &mt.IsDone, &mt.CreatedAt, &mt.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &mt, nil
}

// FindAll Example
func (r *postgresTodosRepository) FindAll(limit, offset, order string) (mtl model.Todos, err error) {

	query := fmt.Sprintf(`
	SELECT *
	FROM tbl_todos
	ORDER BY created_at %s
	LIMIT %s
	OFFSET %s`, order, limit, offset)

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var mt model.Todo

		err = rows.Scan(&mt.TodoID, &mt.TodoName, &mt.IsDone, &mt.CreatedAt, &mt.UpdatedAt)

		if err != nil {
			return nil, err
		}
		mtl = append(mtl, mt)
	}

	return mtl, nil
}

// Update Todos
func (r *postgresTodosRepository) Update(id string, mt *model.Todo) (rowAffected *string, err error) {

	query := `
	UPDATE tbl_todos
	SET
		todo_name	= $1,
		is_done 	= $2,
		updated_at	= $3
	WHERE todo_id = $4`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.Exec(mt.TodoName, mt.IsDone, time.Now(), id)

	if err != nil {
		return nil, err
	}

	rowsAffectedInt64, err := result.RowsAffected()

	if err != nil {
		return nil, err
	}

	rowsAffectedStr := strconv.FormatInt(rowsAffectedInt64, 10)

	rowAffected = &rowsAffectedStr

	return rowAffected, nil

}

// Delete Todos
func (r *postgresTodosRepository) Delete(id string) error {

	query := `
	DELETE FROM tbl_todos
	WHERE todo_id = $1`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

// IsExistsByID Todos
func (r *postgresTodosRepository) IsExistsByID(id string) (isExist bool, err error) {

	query := "SELECT EXISTS(SELECT true FROM tbl_todos WHERE todo_id = $1)"

	statement, err := r.db.Prepare(query)

	if err != nil {
		return false, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&isExist)

	if err != nil {
		return false, err
	}

	return isExist, nil
}

// Count Todos
func (r *postgresTodosRepository) Count() (count int64, err error) {

	query := `
	SELECT COUNT(*)
	FROM tbl_todos`

	err = r.db.QueryRow(query).Scan(&count)

	return count, err
}
