package todos

import (
	"context"
	"todo/business/todos"

	"gorm.io/gorm"
)

type mysqlTodosRepository struct {
	Conn *gorm.DB
}

func NewTodosRepository(conn *gorm.DB) todos.Repository {
	return &mysqlTodosRepository{
		Conn: conn,
	}
}

func (nr *mysqlTodosRepository) Store(ctx context.Context, todosDomain *todos.Domain) (todos.Domain, error) {
	rec := fromDomain(todosDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return todos.Domain{}, result.Error
	}

	record, err := nr.GetByID(ctx, rec.ID)
	if err != nil {
		return 	todos.Domain{}, err
	}

	return record, nil

	// record, err := nr.Conn.Where("todo.id = ?", todo)
}

func (nr *mysqlTodosRepository) GetTodos(ctx context.Context) ([]todos.Domain, error) {
	var recordTodo []Todo

	result := nr.Conn.Unscoped().Find(&recordTodo)
	if result.Error != nil {
		return []todos.Domain{}, result.Error
	}

	return toDomainArray(recordTodo), nil
}

func (nr *mysqlTodosRepository) GetByID(ctx context.Context, todoId int) (todos.Domain, error) {
	rec := Todo{}
	err := nr.Conn.Where("todo.id = ?", todoId).First(&rec).Error
	if err != nil {
		return todos.Domain{}, err
	}

	return rec.toDomain(), nil
}

func (nr *mysqlTodosRepository) Update(ctx context.Context, todosDomain *todos.Domain) (todos.Domain, error) {
	rec := fromDomain(todosDomain)

	result := nr.Conn.Save(&rec)
	if result.Error != nil {
		return todos.Domain{}, result.Error
	}

	record, err := nr.GetByID(ctx, rec.ID)
	if err != nil {
		return 	todos.Domain{}, err
	}	

	return record, nil
}

func (nr *mysqlTodosRepository) Delete(ctx context.Context, todosDomain *todos.Domain) (todos.Domain, error) {
	rec := fromDomain(todosDomain)

	result := nr.Conn.Unscoped().Delete(&rec)
	if result.Error != nil {
		return todos.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}