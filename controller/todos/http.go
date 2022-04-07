package todos

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"todo/business/todos"
	"todo/controller"
	"todo/controller/todos/request"
	"todo/controller/todos/response"

	"github.com/labstack/echo/v4"
)

type TodosController struct {
	todosUsecase todos.Usecase
}

func NewTodosController(todosC todos.Usecase) *TodosController {
	return &TodosController{
		todosUsecase: todosC,
	}
}

func (ctrl *TodosController) Create(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Todo{}
	if err := c.Bind(&req);
	err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.todosUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *TodosController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	req := request.Todo{}
	if err := c.Bind(&req);
	err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = id

	resp, err := ctrl.todosUsecase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}

func (ctrl *TodosController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	req := request.Todo{}
	if err := c.Bind(&req);
	err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = id
	resp, err := ctrl.todosUsecase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewDeleteResponse(c, response.FromDomain(*resp))
}

func (ctrl *TodosController) GetTodos(c echo.Context) error {
	ctx := c.Request().Context()

	todos, err := ctrl.todosUsecase.GetTodos(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.NewResponseArray(todos))
}