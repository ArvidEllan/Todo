package api_todos

import (
	"errors"
	"net/http"

	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/core"
)

type TodosApiService struct{}

// NewTodosApiService creates a default api service
func NewTodosApiService() openapi.TodosApiServicer {
	return &TodosApiService{}
}

// TodosPost - Add new todo note
func (s *TodosApiService) TodosPost(ctx *core.MifyRequestContext, todoNoteCreateRequest openapi.TodoNoteCreateRequest) (openapi.ServiceResponse, error) {
	// TODO: add handler logic

	//TODO: Uncomment the next line to return response Response(200, TodoNote{}) or use other options such as http.Ok
	//return openapi.Response(200, TodoNote{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok
	//return openapi.Response(500, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("TodosPost method not implemented")
}
