// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/services.proto

package proto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	// Einloggen
	Login(ctx context.Context, in *CredentialsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *CredentialsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/task.v1.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	// Einloggen
	Login(context.Context, *CredentialsRequest) (*empty.Empty, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*CredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "task.v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/services.proto",
}

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskServiceClient interface {
	// Erstellen eines Tasks
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TaskEntity, error)
	// Laden eines Tasks
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*TaskEntity, error)
	// Laden aller Tasks. Es werden per default 23 Tasks pro Page gezeigt.
	ListTask(ctx context.Context, in *ListTaskRequest, opts ...grpc.CallOption) (*TaskCollection, error)
	// Einen Task löschen.
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error)
	// Inhalt eines Tasks aktualisieren. Es werden nur gelieferte Felder aktualisiert. Ist eigentlich ein PATCH
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*TaskEntity, error)
	// Benutzerdefinierte Methode um einen Task als abgeschlossen zu setzen.
	CompleteTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*TaskEntity, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TaskEntity, error) {
	out := new(TaskEntity)
	err := c.cc.Invoke(ctx, "/task.v1.TaskService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*TaskEntity, error) {
	out := new(TaskEntity)
	err := c.cc.Invoke(ctx, "/task.v1.TaskService/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ListTask(ctx context.Context, in *ListTaskRequest, opts ...grpc.CallOption) (*TaskCollection, error) {
	out := new(TaskCollection)
	err := c.cc.Invoke(ctx, "/task.v1.TaskService/ListTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error) {
	out := new(DeleteTaskResponse)
	err := c.cc.Invoke(ctx, "/task.v1.TaskService/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*TaskEntity, error) {
	out := new(TaskEntity)
	err := c.cc.Invoke(ctx, "/task.v1.TaskService/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) CompleteTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*TaskEntity, error) {
	out := new(TaskEntity)
	err := c.cc.Invoke(ctx, "/task.v1.TaskService/CompleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
type TaskServiceServer interface {
	// Erstellen eines Tasks
	CreateTask(context.Context, *CreateTaskRequest) (*TaskEntity, error)
	// Laden eines Tasks
	GetTask(context.Context, *GetTaskRequest) (*TaskEntity, error)
	// Laden aller Tasks. Es werden per default 23 Tasks pro Page gezeigt.
	ListTask(context.Context, *ListTaskRequest) (*TaskCollection, error)
	// Einen Task löschen.
	DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error)
	// Inhalt eines Tasks aktualisieren. Es werden nur gelieferte Felder aktualisiert. Ist eigentlich ein PATCH
	UpdateTask(context.Context, *UpdateTaskRequest) (*TaskEntity, error)
	// Benutzerdefinierte Methode um einen Task als abgeschlossen zu setzen.
	CompleteTask(context.Context, *GetTaskRequest) (*TaskEntity, error)
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.TaskService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.TaskService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ListTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ListTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.TaskService/ListTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ListTask(ctx, req.(*ListTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.TaskService/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.TaskService/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_CompleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CompleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.v1.TaskService/CompleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CompleteTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "task.v1.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskService_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskService_GetTask_Handler,
		},
		{
			MethodName: "ListTask",
			Handler:    _TaskService_ListTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _TaskService_DeleteTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskService_UpdateTask_Handler,
		},
		{
			MethodName: "CompleteTask",
			Handler:    _TaskService_CompleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/services.proto",
}

func init() { proto.RegisterFile("proto/services.proto", fileDescriptor_services_2e8ec3b8c2b6b134) }

var fileDescriptor_services_2e8ec3b8c2b6b134 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0x87, 0x45, 0xc5, 0x3f, 0xad, 0x5b, 0x0e, 0x6b, 0x0b, 0x5a, 0x83, 0x38, 0xf8, 0xd8, 0x83,
	0xad, 0xd2, 0x1b, 0xb7, 0xd6, 0x45, 0xbd, 0xd0, 0x4b, 0x69, 0x25, 0xc4, 0x6d, 0x63, 0x26, 0x66,
	0x85, 0xed, 0x75, 0xbc, 0x6b, 0x24, 0x14, 0xe5, 0x92, 0x37, 0x88, 0xf2, 0x52, 0x39, 0x46, 0xca,
	0x0b, 0x44, 0x28, 0x0f, 0x12, 0x79, 0xbd, 0x6b, 0x40, 0x44, 0x48, 0x39, 0x59, 0xfb, 0x1b, 0xcf,
	0xf7, 0xcd, 0x8e, 0x8d, 0xac, 0x34, 0x63, 0x82, 0x79, 0x1c, 0xb2, 0x0d, 0x0d, 0x80, 0xbb, 0xf2,
	0x88, 0x5b, 0x82, 0xf0, 0xb5, 0xbb, 0xf9, 0x66, 0x0f, 0x42, 0xc6, 0xc2, 0x08, 0x3c, 0x92, 0x52,
	0x8f, 0x24, 0x09, 0x13, 0x44, 0x50, 0x96, 0xa8, 0xd7, 0xec, 0xbe, 0xaa, 0xca, 0xd3, 0x45, 0x7e,
	0xe9, 0x41, 0x9c, 0x8a, 0xad, 0x2a, 0x2a, 0x72, 0x0c, 0x9c, 0x93, 0x50, 0x93, 0x47, 0x21, 0x32,
	0x7e, 0xe4, 0x62, 0x35, 0x2b, 0x7d, 0x78, 0x8e, 0x1a, 0x53, 0x16, 0xd2, 0x04, 0xf7, 0x5d, 0xa5,
	0x74, 0xfd, 0x0c, 0x96, 0x90, 0x08, 0x4a, 0x22, 0xfe, 0x17, 0xae, 0x72, 0xe0, 0xc2, 0xee, 0xba,
	0xa5, 0xc8, 0xd5, 0x22, 0x77, 0x52, 0x88, 0x1c, 0xfb, 0xf6, 0xe9, 0xe5, 0xfe, 0x83, 0xe5, 0x34,
	0x3c, 0x92, 0x8b, 0xd5, 0xd8, 0x08, 0xf6, 0xad, 0xa3, 0xbb, 0x3a, 0x32, 0xfe, 0x11, 0xbe, 0xd6,
	0xa6, 0x19, 0x42, 0x7e, 0x06, 0x44, 0x40, 0x11, 0x62, 0xfb, 0x50, 0xa7, 0x42, 0x6d, 0x33, 0xab,
	0x5a, 0x91, 0x4e, 0x12, 0x41, 0xc5, 0xd6, 0xb1, 0xa4, 0xaa, 0xe3, 0x34, 0xbd, 0xa2, 0xc8, 0xc7,
	0x75, 0x2a, 0x20, 0xc6, 0x53, 0xd4, 0xfa, 0x0d, 0x42, 0x12, 0x7b, 0x55, 0x97, 0x4a, 0xce, 0xe2,
	0x4c, 0x89, 0xfb, 0x84, 0x8d, 0x12, 0xe7, 0x5d, 0xd3, 0xe5, 0x0d, 0xfe, 0x83, 0xda, 0x53, 0xca,
	0x4b, 0xdc, 0xe7, 0xaa, 0x4b, 0x47, 0x9a, 0xd7, 0x3b, 0xe2, 0xf9, 0x2c, 0x8a, 0x20, 0x28, 0x3e,
	0x8a, 0xd3, 0x91, 0xcc, 0x36, 0x56, 0x23, 0xe2, 0x05, 0x42, 0xbf, 0x20, 0x82, 0x93, 0x1b, 0xef,
	0x43, 0x8d, 0xec, 0xbf, 0x59, 0xe3, 0x29, 0x4b, 0x38, 0xe8, 0x51, 0xbf, 0x1e, 0x8d, 0x3a, 0x47,
	0xe8, 0x7f, 0xba, 0x3c, 0xdd, 0xe6, 0x3e, 0x3c, 0x7b, 0xfd, 0x2f, 0x92, 0x69, 0xda, 0x87, 0x4c,
	0xb5, 0xd2, 0x05, 0xfa, 0xe8, 0xb3, 0x38, 0xad, 0xe6, 0x7e, 0xdf, 0x5e, 0x07, 0x12, 0xdc, 0x75,
	0xac, 0x43, 0x70, 0xa0, 0x78, 0x3f, 0xcd, 0x87, 0xdd, 0xb0, 0xf6, 0xb8, 0x1b, 0xd6, 0x9e, 0x77,
	0xc3, 0xda, 0xa2, 0x51, 0xfe, 0x4d, 0x4d, 0xf9, 0xf8, 0xfe, 0x1a, 0x00, 0x00, 0xff, 0xff, 0x66,
	0xae, 0x96, 0xfa, 0x0a, 0x03, 0x00, 0x00,
}
