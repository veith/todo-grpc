package todos

import (
	"../protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Gibt den grpc ServiceServer zurück
func GetServiceServer() todo.TodoServiceServer {
	var s todoServiceServer
	return &s
}

// TodoServiceServer is used to implement todo.TodoServiceServer.
type todoServiceServer struct {
}

func (s *todoServiceServer) CompleteTodo(ctx context.Context, req *todo.GetTodoRequest) (*todo.TodoEntity, error) {
	item, err := completeTodoItem(req.Id)
	entity := todo.TodoEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, err
}

func (s *todoServiceServer) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (*todo.TodoEntity, error) {
	item, err := createTodoItem(req.Item)
	entity := todo.TodoEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, err
}

func (s *todoServiceServer) DeleteTodo(ctx context.Context, req *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {
	err := deleteTodoItem(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve entity from the database: %s", err)
	}
	return nil, nil
}

func (s *todoServiceServer) UpdateTodo(ctx context.Context, req *todo.UpdateTodoRequest) (*todo.TodoEntity, error) {
	item, err := updateTodoItem(req.Id, req.Item)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not update entity: %s", err)
	}
	entity := todo.TodoEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, nil
}

func (s *todoServiceServer) GetTodo(ctx context.Context, req *todo.GetTodoRequest) (*todo.TodoEntity, error) {
	item, err := getTodoItem(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Todo not Found: %s", err)
	}
	entity := todo.TodoEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, nil
}

func (s *todoServiceServer) ListTodo(ctx context.Context, req *todo.ListTodoRequest) (*todo.TodoCollection, error) {

	opts := GetListOptionsFromRequest(req)
	items, dbMeta, err := listTodoItems(opts)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Data Error: %s", err)
	}
	var collection []*todo.TodoEntity
	for _, item := range items {
		entity := todo.TodoEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
		collection = append(collection, &entity)
	}
	return &todo.TodoCollection{Data: collection, Links: GenerateCollectionHATEOAS(dbMeta).Links}, nil
}
