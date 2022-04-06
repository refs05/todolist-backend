package todos

import (
	"time"
	todoUsecase "todo/business/todos"
)

type Todo struct {
	ID int `gorm:"primaryKey"`
	Activity string
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}

func fromDomain(domain *todoUsecase.Domain) *Todo {
	return &Todo{
		ID: domain.ID,
		Activity: domain.Activity,
		// CreatedAt: domain.CreatedAt,
		// UpdatedAt: domain.UpdatedAt,
	}
}

func (rec *Todo) toDomain() todoUsecase.Domain {
	return todoUsecase.Domain{
		ID: rec.ID,
		Activity: rec.Activity,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func toDomainArray(modelTodo []Todo) []todoUsecase.Domain {
	var response []todoUsecase.Domain

	for _, val := range modelTodo{
		response = append(response, val.toDomain())
	}

	return response
}