package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/module/todos"
	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/module/todos/model"
)

// sqlite3TodosRepository struct
type sqlite3TodosRepository struct {
	db *sql.DB
}

// NewTodosRepositorySqlite3 func
func NewTodosRepositorySqlite3(db *sql.DB) todos.Repository {
	return &sqlite3TodosRepository{db}
}

// Save Example
func (r *sqlite3TodosRepository) Save(mt *model.Todo) error {

	query := `
	INSERT INTO tbl_todos 
	(
		todo_name,
		is_done,
		created_at,
		updated_at
	)
	VALUES (?,?,?,?)`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(mt.TodoName, mt.IsDone, mt.CreatedAt, mt.UpdatedAt)

	if err != nil {
		return err
	}

	lastInsertIdInt64, err := result.LastInsertId()

	if err != nil {
		return err
	}

	lastInsertIdStr := strconv.FormatInt(lastInsertIdInt64, 10)

	mt.TodoID = lastInsertIdStr

	return nil
}

// FindByID Example
func (r *sqlite3TodosRepository) FindByID(id string) (*model.Todo, error) {

	query := `
	SELECT *
	FROM tbl_todos WHERE todo_id = ?
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
func (r *sqlite3TodosRepository) FindAll(limit, offset, order string) (mtl model.Todos, err error) {

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
func (r *sqlite3TodosRepository) Update(id string, mt *model.Todo) (rowAffected *string, err error) {

	query := `
	UPDATE tbl_todos
	SET
		todo_name	= ?,
		is_done 	= ?,
		updated_at	= ?
	WHERE todo_id=?`

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
func (r *sqlite3TodosRepository) Delete(id string) error {

	query := `
	DELETE FROM tbl_todos
	WHERE todo_id = ?`

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
func (r *sqlite3TodosRepository) IsExistsByID(id string) (isExist bool, err error) {

	query := "SELECT EXISTS(SELECT TRUE from tbl_todos WHERE todo_id = ?)"

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
func (r *sqlite3TodosRepository) Count() (count int64, err error) {

	query := `
	SELECT COUNT(*)
	FROM tbl_todos`

	err = r.db.QueryRow(query).Scan(&count)

	return count, err
}
