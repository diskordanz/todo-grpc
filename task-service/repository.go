package main

import (
	pb "github.com/diskordanz/todo-grpc/task-service/proto"
	"gopkg.in/mgo.v2"
)

const (
	dbName         = "todolist"
	taskCollection = "tasks"
)

type TaskRepository struct {
	session *mgo.Session
}
type Repository interface {
	CreateTask(*pb.CreateTaskRequest) error
	UpdateTask(*pb.UpdateTaskRequest) (*pb.TaskResponse, error)
	GetTask(*pb.GetTaskRequest) (*pb.TaskResponse, error)
	DeleteTask(*pb.DeleteTaskRequest) error
	ListTasks(*pb.ListTasksRequest) (*pb.ListTasksResponse, error)
	Close()
}

func (repo *TaskRepository) CreateTask(req *pb.CreateTaskRequest) error {
	return repo.collection().Insert(req.Task)
}

func (repo *TaskRepository) UpdateTask(req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {

	return &pb.TaskResponse{Task: req.Task}, nil
}

func (repo *TaskRepository) GetTask(req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	//find := repo.collection().Find(req.Id)
	return &pb.TaskResponse{}, nil
}

func (repo *TaskRepository) DeleteTask(req *pb.DeleteTaskRequest) error {
	//updated := append(repo.tasks[:req.Id], repo.tasks[1+req.Id:]...)
	//repo.tasks = updated
	return nil
}

func (repo *TaskRepository) ListTasks(req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	var tasks *pb.ListTasksResponse
	err := repo.collection().Find(nil).All(&tasks.Tasks)
	return tasks, err
}

func (repo *TaskRepository) Close() {
	repo.session.Close()
}

func (repo *TaskRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(taskCollection)
}
