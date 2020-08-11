package controllers

import (
	"net/http"
	"strconv"

	"github.com/s-moteki/go-todo/app/domain"
	"github.com/s-moteki/go-todo/app/interfaces/database"
	"github.com/s-moteki/go-todo/app/usecase"
)

// TodoController Controller
type TodoController struct {
	Interactor usecase.TodoInteractor
}

// Index 検索
func (controller *TodoController) Index(c Context) {

	todos, err := controller.Interactor.Todos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, todos)
}

// Create 登録
func (controller *TodoController) Create(c Context) {
	t := domain.Todo{}
	_ = c.Bind(&t)
	todo, err := controller.Interactor.Add(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, todo)
}

// Show 1件
func (controller *TodoController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := controller.Interactor.TodoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, todo)
}

// Edit 編集
func (controller *TodoController) Edit(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	t := domain.Todo{}
	_ = c.Bind(&t)
	todo, err := controller.Interactor.Update(id, t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, todo)
}

// Search 検索
func (controller *TodoController) Search(c Context) {
	title := c.Query("title")
	todos, err := controller.Interactor.Search(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, todos)
	return
}

// NewTodoController  controller definition
func NewTodoController(sqlHandler database.SQLHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}
