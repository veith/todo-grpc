package todos

import (
	"../protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"upper.io/db.v3"
)

// Gib mir eine DB Session und du bekommst die impl
func Register(database db.Database) todo.TodoServiceServer {
	todoCollection = database.Collection("todo")
	var s server
	return &s
}

var todoCollection db.Collection

// server is used to implement todo.TodoServiceServer.
type server struct {
}

func (s *server) CreateTodo(context.Context, *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {
	return &todo.CreateTodoResponse{Id: "33"}, nil
}

func (s *server) GetTodo(ctx context.Context, req *todo.GetTodoRequest) (*todo.TodoEntity, error) {
	var entity *todo.Todo
	res := todoCollection.Find(db.Cond{"id": req.Id})
	if err := res.One(&entity); err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve entity from the database: %s", err)
	}
	// todo: Apply Hateoas to items
	var links []*todo.Link
	self := todo.Link{Rel: "self", Href: "http:localhost:8080/todos/" + entity.Id, Type: "application/json", Method: todo.Link_GET}
	links = append(links, &self)
	return &todo.TodoEntity{Data: entity, Links: links}, nil
}

func (s *server) ListTodo(ctx context.Context, req *todo.ListTodoRequest) (*todo.TodoCollection, error) {
	var items []*todo.Todo
	res := todoCollection.Find().Paginate(uint(req.Limit))
	if err := res.Page(uint(req.Page)).All(&items); err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
	}
	// todo: Apply Hateoas to items

	var collection []*todo.TodoEntity
	for _, item := range items {
		var links []*todo.Link
		self := todo.Link{Rel: "self", Href: "http:localhost:8080/todos/" + item.Id, Type: "application/json", Method: todo.Link_GET}
		links = append(links, &self)
		var entity todo.TodoEntity
		entity.Data = item
		entity.Links = links
		collection = append(collection, &entity)
	}

	var collectionlinks []*todo.Link
	self := todo.Link{Rel: "self", Href: "http:localhost:8080/todos", Type: "application/json", Method: todo.Link_GET}
	collectionlinks = append(collectionlinks, &self)

	return &todo.TodoCollection{Data: collection, Links: collectionlinks}, nil
}

func (s *server) DeleteTodo(context.Context, *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {
	panic("implement me")
}

func (s *server) UpdateTodo(context.Context, *todo.UpdateTodoRequest) (*todo.UpdateTodoResponse, error) {
	panic("implement me")
}
