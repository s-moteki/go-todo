package database

import (
	"fmt"
	"log"

	"github.com/s-moteki/go-todo/app/domain"
)

// TodoRepository リポジトリ
type TodoRepository struct {
	SQLHandler
}

// Store 登録
func (repo *TodoRepository) Store(t domain.Todo) (id int, err error) {
	fmt.Printf("日付は %s", t.Deadline)
	result, err := repo.Execute(
		"INSERT INTO todo (title, content, deadline, state, created_at, updated_at) VALUES (?,?,?,?,?,?)", t.Title, t.Content, t.Deadline, t.State, t.CreatedAt, t.UpdatedAt,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

// FindAll 全件検索
func (repo *TodoRepository) FindAll() (todos domain.Todos, err error) {

	rows, err := repo.Query("SELECT  * FROM todo")
	defer rows.Close()

	if err != nil {
		log.Print(err.Error())
		return
	}

	for rows.Next() {
		todo := domain.Todo{}
		if err = rows.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Deadline, &todo.State, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			log.Println(err.Error())
			continue
		}
		todos = append(todos, todo)
	}
	return
}

// Edit 更新
func (repo *TodoRepository) Edit(t domain.Todo) (rowCnt int, err error) {
	result, err := repo.Execute(
		"UPDATE todo SET title=?, content=?, deadline=?, state=? WHERE id=?", t.Title, t.Content, t.Deadline, t.State, t.ID,
	)
	if err != nil {
		return
	}
	cnt, err := result.RowsAffected()
	if err != nil {
		return
	}
	rowCnt = int(cnt)
	return
}

// FindByID idからTodoを検索する
func (repo *TodoRepository) FindByID(identifier int) (todo domain.Todo, err error) {

	row, err := repo.Query("SELECT * FROM todo WHERE id = ?", identifier)

	defer row.Close()
	if err != nil {
		return
	}
	row.Next()
	if err = row.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Deadline, &todo.State, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		return
	}
	return
}

// Search タイトルから検索
func (repo *TodoRepository) Search(q string) (todos domain.Todos, err error) {

	rows, err := repo.Query("SELECT * FROM todo WHERE title like ?", fmt.Sprintf("%%%s%%", q))
	defer rows.Close()
	if err != nil {
		log.Print(err.Error())
		return
	}

	for rows.Next() {
		todo := domain.Todo{}
		if err = rows.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Deadline, &todo.State, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			log.Println(err.Error())
			continue
		}
		todos = append(todos, todo)
	}
	return
}
