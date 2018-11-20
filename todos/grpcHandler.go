package todos

import (
	"../protos"
	"encoding/json"
	"github.com/oklog/ulid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"strconv"
	"time"
)

// Gibt den grpc ServiceServer zurück
func GetServiceServer() todo.TodoServiceServer {
	var s todoServiceServer
	return &s
}

// TodoServiceServer is used to implement todo.TodoServiceServer.
type todoServiceServer struct {
}

type Hateoas struct {
	Links []*todo.Link
}

func (h *Hateoas) AddLink(rel, contenttype, href string, method todo.Link_Method) {
	self := todo.Link{Rel: rel, Href: href, Type: contenttype, Method: method}
	h.Links = append(h.Links, &self)
}

func (s *todoServiceServer) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {
	id := GenerateULID().String()
	var item todo.Todo
	item.Id = id
	item.Title = req.Item.Title
	item.Description = req.Item.Description
	if req.Item.Completed != 0 {
		item.Completed = req.Item.Completed
	} else {
		item.Completed = 1
	}
	_, err := dbCollectionTodo.Insert(&item)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not create entity: %s", err)
	}
	e := makeTodoEntity(item)
	return &todo.CreateTodoResponse{Links: e.Links}, nil
}

func (s *todoServiceServer) DeleteTodo(ctx context.Context, req *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {
	err := deleteTodoItem(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve entity from the database: %s", err)
	}
	return nil, nil
}

func (s *todoServiceServer) UpdateTodo(ctx context.Context, req *todo.UpdateTodoRequest) (*todo.UpdateTodoResponse, error) {
	item, err := updateTodoItem(req.Id, req.Item)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not update entity: %s", err)
	}
	entity := makeTodoEntity(item)
	return &todo.UpdateTodoResponse{Data: entity.Data, Links: entity.Links}, nil
}

func (s *todoServiceServer) GetTodo(ctx context.Context, req *todo.GetTodoRequest) (*todo.TodoEntity, error) {
	item, err := getTodoItem(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Todo not Found: %s", err)
	}
	entity := makeTodoEntity(item)
	return &entity, nil
}

func GetListOptionsFromRequest(options interface{}) QueryOptions {
	tmp, _ := json.Marshal(options)
	var opts QueryOptions
	json.Unmarshal(tmp, &opts)
	return opts
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

func GenerateCollectionHATEOAS(dbMeta DBMeta) Hateoas {
	var h Hateoas
	h.AddLink("self", "application/json", "http://localhost:8080/todos?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage), 10), todo.Link_GET)
	if dbMeta.PrevPage != 0 {
		h.AddLink("prev", "application/json", "http://localhost:8080/todos?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage-1), 10), todo.Link_GET)
	}
	if dbMeta.NextPage != 0 {
		h.AddLink("next", "application/json", "http://localhost:8080/todos?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage+1), 10), todo.Link_GET)
	}
	h.AddLink("first", "application/json", "http://localhost:8080/todos?page="+strconv.FormatUint(uint64(dbMeta.FirstPage+1), 10), todo.Link_GET)
	h.AddLink("last", "application/json", "http://localhost:8080/todos?page="+strconv.FormatUint(uint64(dbMeta.LastPage), 10), todo.Link_GET)
	h.AddLink("create", "application/json", "http://localhost:8080/todos", todo.Link_POST)
	return h
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

// Erzeuge eine ULID
func GenerateULID() ulid.ULID {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	newID, _ := ulid.New(ulid.Timestamp(t), entropy)
	return newID
}
