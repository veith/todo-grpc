package todos

import (
	"../protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
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

type Hateoas struct {
	links []*todo.Link
}

func (h *Hateoas) AddHateoas(rel, contenttype, href string, method todo.Link_Method) {
	self := todo.Link{Rel: rel, Href: href, Type: contenttype, Method: method}
	h.links = append(h.links, &self)
}

func (s *server) CreateTodo(context.Context, *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {
	return &todo.CreateTodoResponse{Id: "33"}, nil
}

func (s *server) GetTodo(ctx context.Context, req *todo.GetTodoRequest) (*todo.TodoEntity, error) {
	var item todo.Todo
	res := todoCollection.Find(db.Cond{"id": req.Id})
	if err := res.One(&item); err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve entity from the database: %s", err)
	}
	entity := makeTodoEntity(item)
	return &entity, nil
}

func (s *server) ListTodo(ctx context.Context, req *todo.ListTodoRequest) (*todo.TodoCollection, error) {
	var items []todo.Todo
	var limit uint = 2
	if req.Limit > 0 {
		limit = uint(req.Limit)
	}
	var page uint = 1
	if req.Page > 0 {
		page = uint(req.Page)
	}
	// todo pagination mit hateoas methode
	nextPage := strconv.FormatUint(uint64(page+1), 10)
	currentPage := strconv.FormatUint(uint64(page), 10)
	prevPage := strconv.FormatUint(uint64(page-1), 10)

	res := todoCollection.Find().Paginate(uint(limit))
	if err := res.Page(uint(page)).All(&items); err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
	}
	var collection []*todo.TodoEntity
	for _, item := range items {
		entity := makeTodoEntity(item)
		collection = append(collection, &entity)
	}
	var h Hateoas
	h.AddHateoas("self", "application/json", "http:localhost:8080/todos?page="+currentPage, todo.Link_GET)
	if page > 1 {
		h.AddHateoas("prev", "application/json", "http:localhost:8080/todos?page="+prevPage, todo.Link_GET)
	}
	h.AddHateoas("next", "application/json", "http:localhost:8080/todos?page="+nextPage, todo.Link_GET)
	return &todo.TodoCollection{Data: collection, Links: h.links}, nil
}

func makeTodoEntity(item todo.Todo) todo.TodoEntity {
	var h Hateoas
	h.AddHateoas("self", "application/json", "http:localhost:8080/todos/"+item.Id, todo.Link_GET)
	h.AddHateoas("delete", "application/json", "http:localhost:8080/todos/"+item.Id, todo.Link_DELETE)
	h.AddHateoas("update", "application/json", "http:localhost:8080/todos/"+item.Id, todo.Link_PATCH)
	entity := todo.TodoEntity{Data: &item, Links: h.links}
	return entity
}

func (s *server) DeleteTodo(context.Context, *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {
	panic("implement me")
}

func (s *server) UpdateTodo(context.Context, *todo.UpdateTodoRequest) (*todo.UpdateTodoResponse, error) {
	panic("implement me")
}
