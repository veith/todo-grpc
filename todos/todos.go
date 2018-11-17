package todos

import (
	todo "../protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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

func (s *server) GetTodo(ctx context.Context, req *todo.GetTodoRequest) (*todo.GetTodoResponse, error) {
	var entity *todo.Todo
	res := todoCollection.Find(db.Cond{"id": req.Id})
	if err := res.One(&entity); err != nil {
		return nil, grpc.Errorf(codes.NotFound, "Could not retrieve entity from the database: %s", err)
	}
	return &todo.GetTodoResponse{Item: entity}, nil
}

func (s *server) ListTodo(ctx context.Context, req *todo.ListTodoRequest) (*todo.ListTodoResponse, error) {
	var items []*todo.Todo
	res := todoCollection.Find().Paginate(uint(req.Limit))
	if err := res.Page(uint(req.Page)).All(&items); err != nil {
		return nil, grpc.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
	}
	// Apply Hateoas to items
	var hts *todo.Hateoas

	return &todo.ListTodoResponse{Items: items, Hateaoas: hts}, nil
}

func (s *server) DeleteTodo(context.Context, *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {
	panic("implement me")
}

func (s *server) UpdateTodo(context.Context, *todo.UpdateTodoRequest) (*todo.UpdateTodoResponse, error) {
	panic("implement me")
}
