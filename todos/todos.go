package todos

import (
	"../protos"
	"upper.io/db.v3"
)

// Interface zur DB
var dbCollectionTodo db.Collection
var paginationDefault uint

func ConnectDatabase(database db.Database) {
	dbCollectionTodo = database.Collection("todo")
	paginationDefault = 23
}

func listTodoItems(options QueryOptions) ([]todo.Todo, DBMeta, error) {
	res := dbCollectionTodo.Find()
	var meta DBMeta
	r, meta := ApplyRequestOptionsToQuery(res, options)
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
