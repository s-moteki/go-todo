package usecase

import (
	"time"

	"github.com/s-moteki/go-todo/app/domain"
)

// TodoInteractor Interactor
type TodoInteractor struct {
	TodoRepository TodoRepository
}

// Add todo登録
func (interactor *TodoInteractor) Add(t domain.Todo) (todo domain.Todo, err error) {
	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now
	identifier, err := interactor.TodoRepository.Store(t)
	if err != nil {
		return
	}
	todo, err = interactor.TodoRepository.FindByID(identifier)
	return
}

// Update todo更新
func (interactor *TodoInteractor) Update(identifier int, t domain.Todo) (todo domain.Todo, err error) {
	t.ID = identifier
	_, err = interactor.TodoRepository.Edit(t)
	if err != nil {
		return
	}
	todo, err = interactor.TodoRepository.FindByID(identifier)
	return
}

// Todos todo全件検索
func (interactor *TodoInteractor) Todos() (todos domain.Todos, err error) {
	todos, err = interactor.TodoRepository.FindAll()
	return
}

// TodoByID idから取得
func (interactor *TodoInteractor) TodoByID(identifier int) (todo domain.Todo, err error) {
	todo, err = interactor.TodoRepository.FindByID(identifier)
	return
}

// Search 検索
func (interactor *TodoInteractor) Search(q string) (todos domain.Todos, err error) {
	if q == "" {
		todos, err = interactor.Todos()
	} else {
		todos, err = interactor.TodoRepository.Search(q)
	}
	return
}
