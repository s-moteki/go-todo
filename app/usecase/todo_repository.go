package usecase

import (
	"github.com/s-moteki/go-todo/app/domain"
)

// TodoRepository interface
type TodoRepository interface {
	Store(domain.Todo) (int, error)
	FindByID(int) (domain.Todo, error)
	FindAll() (domain.Todos, error)
	Edit(domain.Todo) (int, error)
	Search(string) (domain.Todos, error)
}
