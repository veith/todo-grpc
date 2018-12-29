package task

import (
	"../proto"
	"github.com/oklog/ulid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var RegisterTaskServiceServer = proto.RegisterTaskServiceServer

// Gibt den grpc ServiceServer zurück
func GetServiceServer() proto.TaskServiceServer {
	var s taskServiceServer
	return &s
}

// taskServiceServer is used to implement taskServiceServer.
type taskServiceServer struct {
}

func (s *taskServiceServer) CompleteTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.TaskEntity, error) {
	taskID, _ := ulid.Parse(req.Id)
	item, err := CompleteTaskItem(taskID)
	entity := proto.TaskEntity{Data: MapTaskToProtoTask(&item), Links: GenerateEntityHateoas(item.Id.String()).Links}

	return &entity, err
}

func (s *taskServiceServer) CreateTask(ctx context.Context, req *proto.CreateTaskRequest) (*proto.TaskEntity, error) {
	item, err := CreateTaskItem(MapProtoTaskToTask(req.Item))
	entity := proto.TaskEntity{Data: MapTaskToProtoTask(&item), Links: GenerateEntityHateoas(item.Id.String()).Links}
	return &entity, err
}

func (s *taskServiceServer) DeleteTask(ctx context.Context, req *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error) {
	taskID, _ := ulid.Parse(req.Id)
	err := DeleteTaskItem(taskID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve entity from the database: %s", err)
	}
	return nil, nil
}

func (s *taskServiceServer) UpdateTask(ctx context.Context, req *proto.UpdateTaskRequest) (*proto.TaskEntity, error) {
	taskID, _ := ulid.Parse(req.Id)

	item, err := UpdateTaskItem(taskID, MapProtoTaskToTask(req.Item))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not update entity: %s", err)
	}
	entity := proto.TaskEntity{Data: MapTaskToProtoTask(&item), Links: GenerateEntityHateoas(item.Id.String()).Links}
	return &entity, nil
}

func (s *taskServiceServer) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.TaskEntity, error) {
	taskID, _ := ulid.Parse(req.Id)
	item, err := GetTaskItem(taskID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Task not Found: %s", err)
	}
	entity := proto.TaskEntity{Data: MapTaskToProtoTask(&item), Links: GenerateEntityHateoas(item.Id.String()).Links}
	return &entity, nil
}

func (s *taskServiceServer) ListTask(ctx context.Context, req *proto.ListTaskRequest) (*proto.TaskCollection, error) {
	//token := ctx.Value("tokenInfo")

	opts := GetListOptionsFromRequest(req)
	items, dbMeta, err := ListTaskItems(opts)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Data Error: %s", err)
	}

	var collection []*proto.TaskEntity
	for _, item := range items {
		entity := proto.TaskEntity{Data: MapTaskToProtoTask(&item), Links: GenerateEntityHateoas(item.Id.String()).Links}
		collection = append(collection, &entity)
	}

	return &proto.TaskCollection{Data: collection, Links: GenerateCollectionHATEOAS(dbMeta).Links}, nil
}
