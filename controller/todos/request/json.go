package request

import "todo/business/todos"

type Todo struct {
	Activity string `json:"activity"`
}

func (req *Todo) ToDomain() *todos.Domain {
	return &todos.Domain{
		Activity: req.Activity,
	}
}