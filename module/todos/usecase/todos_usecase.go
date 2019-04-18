package usecase

import (
	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/module/todos"
	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/module/todos/model"
)

type todosUsecase struct {
	todosRepo todos.Repository
}

func NewTodosUsecase(tr todos.Repository) todos.Usecase {
	return &todosUsecase{
		todosRepo: tr,
	}
}

func (u *todosUsecase) Save(mt *model.Todo) (err error) {

	err = u.todosRepo.Save(mt)

	return err
}

func (u *todosUsecase) FindByID(id string) (mt *model.Todo, err error) {

	mt, err = u.todosRepo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return mt, nil
}

func (u *todosUsecase) FindAll(limit, offset, order string) (mtl model.Todos, count int64, err error) {

	mtl, err = u.todosRepo.FindAll(limit, offset, order)

	if err != nil {
		return nil, -1, err
	}

	count, err = u.todosRepo.Count()

	if err != nil {
		return nil, -1, err
	}

	return mtl, count, nil
}

func (u *todosUsecase) Update(id string, mt *model.Todo) (rowAffected *string, err error) {

	v, err := u.todosRepo.FindByID(id)

	if err != nil {
		return nil, err
	}

	if mt.TodoName != "" {
		v.TodoName = mt.TodoName
	}

	if mt.IsDone != v.IsDone {
		v.IsDone = mt.IsDone
	}

	rowAffected, err = u.todosRepo.Update(id, v)

	if err != nil {
		return nil, err
	}

	return rowAffected, err
}

func (u *todosUsecase) Delete(idUser string) (err error) {

	err = u.todosRepo.Delete(idUser)

	return err
}

func (u *todosUsecase) IsExistsByID(todoID string) (isExist bool, err error) {

	isExist, err = u.todosRepo.IsExistsByID(todoID)

	if err != nil {
		return false, err
	}

	return isExist, nil
}

func (u *todosUsecase) Count() (count int64, err error) {

	count, err = u.todosRepo.Count()

	if err != nil {
		return -1, err
	}

	return count, nil
}
