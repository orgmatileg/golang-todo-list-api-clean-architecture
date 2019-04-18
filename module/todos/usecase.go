package todos

import "github.com/orgmatileg/golang-todo-list-api-clean-architecture/module/todos/model"

type Usecase interface {
	Save(*model.Todo) error
	FindByID(id string) (*model.Todo, error)
	FindAll(limit, offset, order string) (mt model.Todos, count int64, err error)
	Update(id string, todoModel *model.Todo) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
	Count() (int64, error)
}
