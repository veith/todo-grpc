package task

import (
	"../protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Gibt den grpc ServiceServer zurück
func GetServiceServer() task.TodoServiceServer {
	var s taskServiceServer
	return &s
}

// TodoServiceServer is used to implement task.TodoServiceServer.
type taskServiceServer struct {
}

func (s *taskServiceServer) CompleteTodo(ctx context.Context, req *task.GetTodoRequest) (*task.TodoEntity, error) {
	item, err := completeTodoItem(req.Id)
	entity := task.TodoEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, err
}

func (s *taskServiceServer) CreateTodo(ctx context.Context, req *task.CreateTodoRequest) (*task.TodoEntity, error) {
	item, err := createTodoItem(req.Item)
	entity := task.TodoEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, err
}

func (s *taskServiceServer) DeleteTodo(ctx context.Context, req *task.DeleteTodoRequest) (*task.DeleteTodoResponse, error) {
	err := deleteTodoItem(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve entity from the database: %s", err)
	}
	return nil, nil
}

func (s *taskServiceServer) UpdateTodo(ctx context.Context, req *task.UpdateTodoRequest) (*task.TodoEntity, error) {
	item, err := updateTodoItem(req.Id, req.Item)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not update entity: %s", err)
	}
	entity := task.TodoEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, nil
}

func (s *taskServiceServer) GetTodo(ctx context.Context, req *task.GetTodoRequest) (*task.TodoEntity, error) {
	item, err := getTodoItem(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Todo not Found: %s", err)
	}
	entity := task.TodoEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, nil
}

func (s *taskServiceServer) ListTodo(ctx context.Context, req *task.ListTodoRequest) (*task.TodoCollection, error) {

	opts := GetListOptionsFromRequest(req)
	items, dbMeta, err := listTodoItems(opts)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Data Error: %s", err)
	}
	var collection []*task.TodoEntity
	for _, item := range items {
		entity := task.TodoEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
		collection = append(collection, &entity)
	}
	return &task.TodoCollection{Data: collection, Links: GenerateCollectionHATEOAS(dbMeta).Links}, nil
}
