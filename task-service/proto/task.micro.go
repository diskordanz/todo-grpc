// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/task.proto

/*
Package go_micro_srv_task is a generated protocol buffer package.

It is generated from these files:
	proto/task.proto

It has these top-level messages:
	Task
	CreateTaskRequest
	UpdateTaskRequest
	GetTaskRequest
	DeleteTaskRequest
	ListTasksRequest
	ListTasksResponse
	TaskResponse
*/
package go_micro_srv_task

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = google_protobuf1.Empty{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TaskService service

type TaskService interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...client.CallOption) (*TaskResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...client.CallOption) (*TaskResponse, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...client.CallOption) (*google_protobuf1.Empty, error)
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error)
	ListTasks(ctx context.Context, in *ListTasksRequest, opts ...client.CallOption) (*ListTasksResponse, error)
}

type taskService struct {
	c    client.Client
	name string
}

func NewTaskService(name string, c client.Client) TaskService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.task"
	}
	return &taskService{
		c:    c,
		name: name,
	}
}

func (c *taskService) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.CreateTask", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.UpdateTask", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...client.CallOption) (*google_protobuf1.Empty, error) {
	req := c.c.NewRequest(c.name, "TaskService.DeleteTask", in)
	out := new(google_protobuf1.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetTask(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetTask", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) ListTasks(ctx context.Context, in *ListTasksRequest, opts ...client.CallOption) (*ListTasksResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.ListTasks", in)
	out := new(ListTasksResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskService service

type TaskServiceHandler interface {
	CreateTask(context.Context, *CreateTaskRequest, *TaskResponse) error
	UpdateTask(context.Context, *UpdateTaskRequest, *TaskResponse) error
	DeleteTask(context.Context, *DeleteTaskRequest, *google_protobuf1.Empty) error
	GetTask(context.Context, *GetTaskRequest, *TaskResponse) error
	ListTasks(context.Context, *ListTasksRequest, *ListTasksResponse) error
}

func RegisterTaskServiceHandler(s server.Server, hdlr TaskServiceHandler, opts ...server.HandlerOption) error {
	type taskService interface {
		CreateTask(ctx context.Context, in *CreateTaskRequest, out *TaskResponse) error
		UpdateTask(ctx context.Context, in *UpdateTaskRequest, out *TaskResponse) error
		DeleteTask(ctx context.Context, in *DeleteTaskRequest, out *google_protobuf1.Empty) error
		GetTask(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error
		ListTasks(ctx context.Context, in *ListTasksRequest, out *ListTasksResponse) error
	}
	type TaskService struct {
		taskService
	}
	h := &taskServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TaskService{h}, opts...))
}

type taskServiceHandler struct {
	TaskServiceHandler
}

func (h *taskServiceHandler) CreateTask(ctx context.Context, in *CreateTaskRequest, out *TaskResponse) error {
	return h.TaskServiceHandler.CreateTask(ctx, in, out)
}

func (h *taskServiceHandler) UpdateTask(ctx context.Context, in *UpdateTaskRequest, out *TaskResponse) error {
	return h.TaskServiceHandler.UpdateTask(ctx, in, out)
}

func (h *taskServiceHandler) DeleteTask(ctx context.Context, in *DeleteTaskRequest, out *google_protobuf1.Empty) error {
	return h.TaskServiceHandler.DeleteTask(ctx, in, out)
}

func (h *taskServiceHandler) GetTask(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error {
	return h.TaskServiceHandler.GetTask(ctx, in, out)
}

func (h *taskServiceHandler) ListTasks(ctx context.Context, in *ListTasksRequest, out *ListTasksResponse) error {
	return h.TaskServiceHandler.ListTasks(ctx, in, out)
}
