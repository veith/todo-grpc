package task

import (
	"../protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Gibt den grpc ServiceServer zurück
func GetServiceServer() task.TaskServiceServer {
	var s taskServiceServer
	return &s
}

// taskServiceServer is used to implement task.taskServiceServer.
type taskServiceServer struct {
}

func (s *taskServiceServer) CompleteTask(ctx context.Context, req *task.GetTaskRequest) (*task.TaskEntity, error) {
	item, err := completeTaskItem(req.Id)
	entity := task.TaskEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, err
}

func (s *taskServiceServer) CreateTask(ctx context.Context, req *task.CreateTaskRequest) (*task.TaskEntity, error) {
	item, err := createTaskItem(req.Item)
	entity := task.TaskEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, err
}

func (s *taskServiceServer) DeleteTask(ctx context.Context, req *task.DeleteTaskRequest) (*task.DeleteTaskResponse, error) {
	err := deleteTaskItem(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve entity from the database: %s", err)
	}
	return nil, nil
}

func (s *taskServiceServer) UpdateTask(ctx context.Context, req *task.UpdateTaskRequest) (*task.TaskEntity, error) {
	item, err := updateTaskItem(req.Id, req.Item)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not update entity: %s", err)
	}
	entity := task.TaskEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, nil
}

func (s *taskServiceServer) GetTask(ctx context.Context, req *task.GetTaskRequest) (*task.TaskEntity, error) {
	item, err := getTaskItem(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Task not Found: %s", err)
	}
	entity := task.TaskEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, nil
}

func (s *taskServiceServer) ListTask(ctx context.Context, req *task.ListTaskRequest) (*task.TaskCollection, error) {

	opts := GetListOptionsFromRequest(req)
	items, dbMeta, err := listTaskItems(opts)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Data Error: %s", err)
	}
	var collection []*task.TaskEntity
	for _, item := range items {
		entity := task.TaskEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
		collection = append(collection, &entity)
	}
	return &task.TaskCollection{Data: collection, Links: GenerateCollectionHATEOAS(dbMeta).Links}, nil
}
