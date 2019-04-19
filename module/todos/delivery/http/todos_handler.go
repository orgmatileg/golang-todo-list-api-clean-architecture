package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/module/todos/model"

	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/helper"
	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/module/todos"

	"github.com/gorilla/mux"
)

type HttpTodosHandler struct {
	TUsecase todos.Usecase
}

func NewTodosHttpHandler(r *mux.Router, tu todos.Usecase) {

	handler := HttpTodosHandler{
		TUsecase: tu,
	}

	r.HandleFunc("/todos", handler.TodosSaveHttpHandler).Methods("POST")
	r.HandleFunc("/todos", handler.TodosFindAllHttpHandler).Methods("GET")
	r.HandleFunc("/todos/count", handler.TodosCountHttpHandler).Methods("GET")
	r.HandleFunc("/todos/{id}", handler.TodosFindByIDHttpHandler).Methods("GET")
	r.HandleFunc("/todos/{id}", handler.TodosUpdateHttpHandler).Methods("PUT")
	r.HandleFunc("/todos/{id}", handler.TodosDeleteHttpHandler).Methods("DELETE")
	r.HandleFunc("/todos/{id}/exists", handler.TodosExistsByIDHttpHandler).Methods("GET")
}

func (u *HttpTodosHandler) TodosSaveHttpHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	mt := model.NewTodo()

	err := decoder.Decode(mt)

	res := helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	err = u.TUsecase.Save(mt)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mt

}

func (u *HttpTodosHandler) TodosFindAllHttpHandler(w http.ResponseWriter, r *http.Request) {

	queryParam := r.URL.Query()

	// Set default query
	limit := "10"
	offset := "0"
	order := "desc"

	if v := queryParam.Get("limit"); v != "" {
		limit = queryParam.Get("limit")
	}

	if v := queryParam.Get("offset"); v != "" {
		offset = queryParam.Get("offset")
	}

	if v := queryParam.Get("order"); v != "" {
		order = queryParam.Get("order")
	}

	res := helper.Response{}

	mtl, count, err := u.TUsecase.FindAll(limit, offset, order)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload, res.Body.Count = mtl, count

}

func (u *HttpTodosHandler) TodosFindByIDHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	mt, err := u.TUsecase.FindByID(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = mt

}

func (u *HttpTodosHandler) TodosUpdateHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	decoder := json.NewDecoder(r.Body)

	var mt model.Todo

	err := decoder.Decode(&mt)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	rowsAffected, err := u.TUsecase.Update(idP, &mt)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = fmt.Sprintf("Total rows affected: %s", *rowsAffected)

}

func (u *HttpTodosHandler) TodosDeleteHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	idP := vars["id"]

	res := helper.Response{}

	err := u.TUsecase.Delete(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = "OK"

}

func (u *HttpTodosHandler) TodosExistsByIDHttpHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := helper.Response{}

	idP := vars["id"]

	isExists, err := u.TUsecase.IsExistsByID(idP)

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	payload := struct {
		IsExists bool `json:"is_exists"`
	}{
		IsExists: isExists,
	}

	res.Body.Payload = payload

}

func (u *HttpTodosHandler) TodosCountHttpHandler(w http.ResponseWriter, r *http.Request) {

	res := helper.Response{}

	count, err := u.TUsecase.Count()

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	payload := struct {
		Count int64 `json:"count"`
	}{
		Count: count,
	}

	res.Body.Payload = payload

}
