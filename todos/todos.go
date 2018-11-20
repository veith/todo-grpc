package todos

import (
	"../protos"
	"strings"
	"upper.io/db.v3"
)

// Interface zur DB
var dbCollectionTodo db.Collection
var paginationDefault uint

func ConnectDatabase(database db.Database) {
	dbCollectionTodo = database.Collection("todo")
	paginationDefault = 23
}

type QueryOptions struct {
	Fields  string `json:"fields,omitempty"`
	Sort    string `json:"sort,omitempty"`
	Filter  string `json:"filter,omitempty"`
	Count   bool   `json:"count,omitempty"`
	Sum     string `json:"sum,omitempty"`
	Context string `json:"context,omitempty"`
	Limit   uint   `json:"limit,omitempty"`
	Page    uint   `json:"page,omitempty"`
}

type DBMeta struct {
	Count       uint64
	CurrentPage uint
	NextPage    uint
	PrevPage    uint
	FirstPage   uint
	LastPage    uint
}

func ApplyRequestOptionsToResult(res db.Result, options QueryOptions) (db.Result, DBMeta) {
	var meta DBMeta
	if options.Limit > 0 {
		res = res.Paginate(options.Limit)
	} else {
		res.Paginate(paginationDefault)
	}

	if options.Fields != "" {
		fields := strings.Split(options.Fields, ",")
		s := make([]interface{}, len(fields))
		for i, field := range fields {
			s[i] = field
		}
		res = res.Select(s...)
	}

	if options.Sort != "" {
		res = res.OrderBy(options.Sort)
	}

	meta.CurrentPage = 1
	if options.Page > 0 {
		meta.CurrentPage = options.Page
		res = res.Page(options.Page)
	}
	pages, _ := res.TotalPages()
	meta.LastPage = pages
	meta.FirstPage = 1
	if meta.CurrentPage < meta.LastPage {
		meta.NextPage = meta.CurrentPage + 1
	}
	if meta.CurrentPage > 1 {
		meta.PrevPage = meta.CurrentPage - 1
	}

	if options.Count {
		meta.Count, _ = res.Count()
	}

	return res, meta
}

func listTodoItems(options QueryOptions) ([]todo.Todo, DBMeta, error) {
	res := dbCollectionTodo.Find()
	var meta DBMeta
	r, meta := ApplyRequestOptionsToResult(res, options)
	var items []todo.Todo
	err := r.All(&items)

	return items, meta, err
}
func deleteTodoItem(id string) error {
	var item todo.Todo
	res := dbCollectionTodo.Find(db.Cond{"id": id})
	if err := res.One(&item); err != nil {
		return err
	}
	err := res.Delete()
	return err
}
func getTodoItem(id string) (todo.Todo, error) {
	var item todo.Todo
	res := dbCollectionTodo.Find(db.Cond{"id": id})
	err := res.One(&item)
	return item, err
}

func updateTodoItem(id string, data *todo.Todo) (todo.Todo, error) {
	var item todo.Todo
	res := dbCollectionTodo.Find(db.Cond{"id": id})

	if err := res.One(&item); err != nil {
		return item, err
	}
	// fields update
	item.Title = data.Title
	item.Description = data.Description
	item.Completed = data.Completed

	err := res.Update(&item)
	return item, err
}
