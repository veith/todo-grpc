package task

import (
	"../task"
	"github.com/oklog/ulid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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

// Override Funktion um nicht über die default auth-middleware
func (s *taskServiceServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (s *taskServiceServer) CompleteTask(ctx context.Context, req *GetTaskRequest) (*TaskEntity, error) {
	taskID, _ := ulid.Parse(req.Id)
	item, err := task.CompleteTaskItem(taskID)
	entity := TaskEntity{Data: MapTaskToProtoTask(&item), Links: GenerateEntityHateoas(item.Id.String()).Links}

	return &entity, err
}

func (s *taskServiceServer) CreateTask(ctx context.Context, req *CreateTaskRequest) (*TaskEntity, error) {
	item, err := task.CreateTaskItem(MapProtoTaskToTask(req.Item))
	entity := TaskEntity{Data: MapTaskToProtoTask(&item), Links: GenerateEntityHateoas(item.Id.String()).Links}
	return &entity, err
}

func (s *taskServiceServer) DeleteTask(ctx context.Context, req *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	taskID, _ := ulid.Parse(req.Id)
	err := task.DeleteTaskItem(taskID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve entity from the database: %s", err)
	}
	return nil, nil
}

func (s *taskServiceServer) UpdateTask(ctx context.Context, req *UpdateTaskRequest) (*TaskEntity, error) {
	taskID, _ := ulid.Parse(req.Id)

	item, err := task.UpdateTaskItem(taskID, MapProtoTaskToTask(req.Item))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not update entity: %s", err)
	}
	entity := TaskEntity{Data: MapTaskToProtoTask(&item), Links: GenerateEntityHateoas(item.Id.String()).Links}
	return &entity, nil
}

func (s *taskServiceServer) GetTask(ctx context.Context, req *GetTaskRequest) (*TaskEntity, error) {
	taskID, _ := ulid.Parse(req.Id)
	item, err := task.GetTaskItem(taskID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Task not Found: %s", err)
	}
	entity := TaskEntity{Data: MapTaskToProtoTask(&item), Links: GenerateEntityHateoas(item.Id.String()).Links}
	return &entity, nil
}

func (s *taskServiceServer) ListTask(ctx context.Context, req *ListTaskRequest) (*TaskCollection, error) {
	//token := ctx.Value("tokenInfo")

	opts := GetListOptionsFromRequest(req)
	items, dbMeta, err := task.ListTaskItems(opts)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Data Error: %s", err)
	}
	var collection []*TaskEntity
	for _, item := range items {
		entity := TaskEntity{Data: MapTaskToProtoTask(&item), Links: GenerateEntityHateoas(item.Id.String()).Links}
		collection = append(collection, &entity)
	}

	// create and send header
	header := metadata.Pairs("Set-Cookie", "Authorization=Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlcm5hbWUiOiJ2ZWl0aHpAZ21haWwuY29tIiwiaWF0IjoxNTE2MjM5MDIyfQ.qL-xs3KVWWpe6lCMPCDPE4ZW2EoAo0KI5g36Dm1ouKU;HttpOnly")
	grpc.SendHeader(ctx, header)

	return &TaskCollection{Data: collection, Links: GenerateCollectionHATEOAS(dbMeta).Links}, nil
}
