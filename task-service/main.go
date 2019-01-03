package main

import (
	"errors"
	"fmt"

	pb "github.com/diskordanz/todo-grpc/task-service/proto"
	"github.com/golang/protobuf/ptypes/empty"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type IRepository interface {
	CreateTask(*pb.CreateTaskRequest) (*pb.TaskResponse, error)
	UpdateTask(*pb.UpdateTaskRequest) (*pb.TaskResponse, error)
	GetTask(*pb.GetTaskRequest) (*pb.TaskResponse, error)
	DeleteTask(*pb.DeleteTaskRequest) error
	ListTasks(*pb.ListTasksRequest) (*pb.ListTasksResponse, error)
}

type Repository struct {
	tasks []*pb.Task
}

func (repo *Repository) CreateTask(req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	updated := append(repo.tasks, req.Task)
	repo.tasks = updated
	return &pb.TaskResponse{Task: req.Task}, nil
}

func (repo *Repository) UpdateTask(req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {

	return &pb.TaskResponse{Task: req.Task}, nil
}

func (repo *Repository) GetTask(req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	find := repo.tasks[req.Id]
	return &pb.TaskResponse{Task: find}, nil
}

func (repo *Repository) DeleteTask(req *pb.DeleteTaskRequest) error {
	updated := append(repo.tasks[:req.Id], repo.tasks[1+req.Id:]...)
	repo.tasks = updated
	return nil
}

func (repo *Repository) ListTasks(req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	return &pb.ListTasksResponse{Tasks: repo.tasks}, nil
}

type Service struct {
	Repo IRepository
}

func (s *Service) CreateTask(ctx context.Context, req *pb.CreateTaskRequest, res *pb.TaskResponse) error {
	if req == nil {
		return errors.New("got nil request")
	}
	task, err := s.Repo.CreateTask(req)
	if err != nil {
		return err
	}
	res = task
	return nil
}

func (s *Service) ListTasks(ctx context.Context, req *pb.ListTasksRequest, res *pb.ListTasksResponse) error {
	if req == nil {
		return errors.New("got nil request")
	}
	task, err := s.Repo.ListTasks(req)
	if err != nil {
		return err
	}
	res = task
	return nil
}

func (s *Service) GetTask(ctx context.Context, req *pb.GetTaskRequest, res *pb.TaskResponse) error {
	if req == nil {
		return errors.New("got nil request")
	}
	task, err := s.Repo.GetTask(req)
	if err != nil {
		return err
	}
	res = task

	return nil
}

func (s *Service) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest, res *pb.TaskResponse) error {
	if req == nil {
		return errors.New("got nil request")
	}
	task, err := s.Repo.UpdateTask(req)
	if err != nil {
		return err
	}
	res = task

	return nil
}

func (s *Service) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest, res *empty.Empty) error {
	if req == nil {
		return nil
	}
	err := s.Repo.DeleteTask(req)
	if err != nil {
		return err
	}
	res = &empty.Empty{}
	return nil
}

func main() {

	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.task"),
		micro.Version("latest"),
	)
	srv.Init()

	pb.RegisterTaskServiceHandler(srv.Server(), &Service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
	/*
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()

		pb.RegisterTaskServiceServer(s, &Service{Repo: repo})

		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}*/
}
