package main

import (
	"errors"

	pb "github.com/diskordanz/todo-grpc/task-service/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
)

type Service struct {
	session *mgo.Session
}

func (s *Service) GetRepo() Repository {
	return &TaskRepository{s.session.Clone()}
}

func (s *Service) CreateTask(ctx context.Context, req *pb.CreateTaskRequest, res *pb.TaskResponse) error {
	if req == nil {
		return errors.New("got nil request")
	}
	repo := s.GetRepo()
	defer repo.Close()

	err := repo.CreateTask(req)
	if err != nil {
		return err
	}
	res.Task = req.Task
	return nil
}

func (s *Service) ListTasks(ctx context.Context, req *pb.ListTasksRequest, res *pb.ListTasksResponse) error {
	if req == nil {
		return errors.New("got nil request")
	}
	repo := s.GetRepo()
	defer repo.Close()

	tasks, err := repo.ListTasks(req)
	if err != nil {
		return err
	}
	res.Tasks = tasks.Tasks
	return nil
}

func (s *Service) GetTask(ctx context.Context, req *pb.GetTaskRequest, res *pb.TaskResponse) error {
	if req == nil {
		return errors.New("got nil request")
	}

	repo := s.GetRepo()
	defer repo.Close()

	task, err := repo.GetTask(req)
	if err != nil {
		return err
	}
	res.Task = task.Task

	return nil
}

func (s *Service) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest, res *pb.TaskResponse) error {
	if req == nil {
		return errors.New("got nil request")
	}

	repo := s.GetRepo()
	defer repo.Close()

	task, err := repo.UpdateTask(req)
	if err != nil {
		return err
	}
	res.Task = task.Task

	return nil
}

func (s *Service) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest, res *empty.Empty) error {
	if req == nil {
		return errors.New("got nil request")
	}

	repo := s.GetRepo()
	defer repo.Close()

	err := repo.DeleteTask(req)
	if err != nil {
		return err
	}
	res = &empty.Empty{}
	return nil
}
