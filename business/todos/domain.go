package todos

import (
	"context"
	"time"
)

type Domain struct {
	ID int
	Activity string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Usecase interface {
	Store(ctx context.Context, todosDomain *Domain) (Domain, error)
	GetTodos(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, todosDomain *Domain) ([]Domain, error)
	Delete(ctx context.Context, todosDomain *Domain) ([]Domain, error)
}

type Repository interface {
	Store(ctx context.Context, todosDomain *Domain) (Domain, error)
	GetTodos(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, todosDomain *Domain) ([]Domain, error)
	Delete(ctx context.Context, todosDomain *Domain) ([]Domain, error)
}