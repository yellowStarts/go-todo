package todo

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type TodoRepository struct {
	database *gorm.DB
}

func (r *TodoRepository) FindAll() []Todo {
	var todos []Todo
	r.database.Find(&todos)
	return todos
}

func (r *TodoRepository) Find(id int) (Todo, error) {
	var todo Todo
	err := r.database.Find(&todo, id).Error
	if todo.Name == "" {
		err = errors.New("Todo not found")
	}
	return todo, err
}

func (r *TodoRepository) Create(todo Todo) (Todo, error) {
	err := r.database.Create(&todo).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *TodoRepository) Save(todo Todo) (Todo, error) {
	err := r.database.Save(todo).Error
	return todo, err
}

func (r *TodoRepository) Delete(id int) int64 {
	count := r.database.Delete(&Todo{}, id).RowsAffected
	return count
}

func NewTodoRepository(database *gorm.DB) *TodoRepository {
	return &TodoRepository{
		database: database,
	}
}
