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

func (s *todoServiceServer) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (*todo.TodoEntity, error) {
	item, err := createTodoItem(req.Item)
	entity := makeTodoEntity(item)
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
	entity := makeTodoEntity(item)
	return &entity, nil
}

func (s *todoServiceServer) GetTodo(ctx context.Context, req *todo.GetTodoRequest) (*todo.TodoEntity, error) {
	item, err := getTodoItem(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Todo not Found: %s", err)
	}
	entity := makeTodoEntity(item)
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
		entity := makeTodoEntity(item)
		collection = append(collection, &entity)
	}
	return &todo.TodoCollection{Data: collection, Links: GenerateCollectionHATEOAS(dbMeta).Links}, nil
}

// erzeuge aus einem db todo item eine todoEntity
func makeTodoEntity(item todo.Todo) todo.TodoEntity {
	var h Hateoas
	h.AddLink("self", "application/json", "http://localhost:8080/todos/"+item.Id, todo.Link_GET)
	h.AddLink("delete", "application/json", "http://localhost:8080/todos/"+item.Id, todo.Link_DELETE)
	h.AddLink("update", "application/json", "http://localhost:8080/todos/"+item.Id, todo.Link_PATCH)
	h.AddLink("parent", "application/json", "http://localhost:8080/todos", todo.Link_GET)
	entity := todo.TodoEntity{Data: &item, Links: h.Links}
	return entity
}
