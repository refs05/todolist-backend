package todos

import (
	"context"
	"time"
)

type todosUsecase struct {
	todosRepository Repository
	contextTimeout time.Duration
}

func NewTodosUsecase(tr Repository, timeout time.Duration) Usecase {
	return &todosUsecase {
		todosRepository: tr,
		contextTimeout: timeout,
	}
}

func (tu *todosUsecase) Store(ctx context.Context, todosDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()
	
	result, err := tu.todosRepository.Store(ctx, todosDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (tu *todosUsecase) GetTodos(ctx context.Context) ([]Domain, error) {
	result, err := tu.todosRepository.GetTodos(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (tu *todosUsecase) GetByID(ctx context.Context, todoId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()

	res, err := tu.todosRepository.GetByID(ctx, todoId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (tu *todosUsecase) Update(ctx context.Context, todosDomain *Domain) (*Domain, error) {
	result, err := tu.todosRepository.Update(ctx, todosDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func(tu *todosUsecase) Delete(ctx context.Context, todosDomain *Domain) (*Domain, error) {
	result, err := tu.todosRepository.Delete(ctx, todosDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}