package todos

import "todo/business/todos"

type TodosController struct {
	todosUsecase todos.Usecase
}

