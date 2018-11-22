package task

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Gibt den grpc ServiceServer zurück
func GetServiceServer() TaskServiceServer {
	var s taskServiceServer
	return &s
}

// taskServiceServer is used to implement taskServiceServer.
type taskServiceServer struct {
}

func (s *taskServiceServer) CompleteTask(ctx context.Context, req *GetTaskRequest) (*TaskEntity, error) {
	item, err := completeTaskItem(req.Id)
	entity := TaskEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, err
}

func (s *taskServiceServer) CreateTask(ctx context.Context, req *CreateTaskRequest) (*TaskEntity, error) {
	item, err := createTaskItem(req.Item)
	entity := TaskEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, err
}

func (s *taskServiceServer) DeleteTask(ctx context.Context, req *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	err := deleteTaskItem(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve entity from the database: %s", err)
	}
	return nil, nil
}

func (s *taskServiceServer) UpdateTask(ctx context.Context, req *UpdateTaskRequest) (*TaskEntity, error) {
	item, err := updateTaskItem(req.Id, req.Item)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not update entity: %s", err)
	}
	entity := TaskEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, nil
}

func (s *taskServiceServer) GetTask(ctx context.Context, req *GetTaskRequest) (*TaskEntity, error) {
	item, err := getTaskItem(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Task not Found: %s", err)
	}
	entity := TaskEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
	return &entity, nil
}

func (s *taskServiceServer) ListTask(ctx context.Context, req *ListTaskRequest) (*TaskCollection, error) {

	opts := GetListOptionsFromRequest(req)
	items, dbMeta, err := listTaskItems(opts)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Data Error: %s", err)
	}
	var collection []*TaskEntity
	for _, item := range items {
		entity := TaskEntity{Data: &item, Links: GenerateEntityHateoas(item.Id).Links}
		collection = append(collection, &entity)
	}
	return &TaskCollection{Data: collection, Links: GenerateCollectionHATEOAS(dbMeta).Links}, nil
}
