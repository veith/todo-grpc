package task

import (
	"../protos"
	"upper.io/db.v3"
)

// Interface zur DB
var dbCollectionTodo db.Collection
var paginationDefault uint

func ConnectDatabase(database db.Database) {
	dbCollectionTodo = database.Collection("tasks")
	paginationDefault = 23
}
func createTodoItem(data *task.Todo) (task.Todo, error) {
	var item task.Todo
	//task ulid typ in protobuf bauen

	item.Id = GenerateULID().String()
	item.Title = data.Title
	item.Description = data.Description
	if data.Completed != 0 {
		item.Completed = data.Completed
	} else {
		item.Completed = 1
	}
	// id interface not needed, we create the ids ourself
	_, err := dbCollectionTodo.Insert(&item)
	//fire("item.generated",&item)
	return item, err
}

func listTodoItems(options QueryOptions) ([]task.Todo, DBMeta, error) {
	res := dbCollectionTodo.Find()
	var meta DBMeta
	res, meta = ApplyRequestOptionsToQuery(res, options)
	var items []task.Todo
	err := res.All(&items)

	return items, meta, err
}

func completeTodoItem(id string) (task.Todo, error) {
	var item task.Todo
	item.Completed = 2
	return updateTodoItem(id, &item)
}

func deleteTodoItem(id string) error {
	var item task.Todo
	res := dbCollectionTodo.Find(db.Cond{"id": id})
	if err := res.One(&item); err != nil {
		return err
	}
	err := res.Delete()
	return err
}
func getTodoItem(id string) (task.Todo, error) {
	var item task.Todo
	res := dbCollectionTodo.Find(db.Cond{"id": id})
	err := res.One(&item)
	return item, err
}

func updateTodoItem(id string, data *task.Todo) (task.Todo, error) {
	var item task.Todo
	res := dbCollectionTodo.Find(db.Cond{"id": id})
	// fields to update
	item.Title = data.Title
	item.Description = data.Description
	item.Completed = data.Completed

	if err := res.Update(&item); err != nil {
		return task.Todo{}, err
	}
	// read your write
	err := res.One(&item)
	return item, err
}
