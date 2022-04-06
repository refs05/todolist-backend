package response

import (
	"time"
	"todo/business/todos"
)

type Todo struct {
	ID int `json:"id"`
	Activity string `json:"activity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain todos.Domain) Todo {
	return Todo{
		ID: domain.ID,
		Activity: domain.Activity,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func NewResponseArray(domainTodos []todos.Domain) []Todo {
	var resp []Todo

	for _, value := range domainTodos {
		resp = append(resp, FromDomain(value))
	}

	return resp
}