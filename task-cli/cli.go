package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/diskordanz/todo-grpc/task-service/proto"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/micro/cmd"
	"golang.org/x/net/context"
)

const (
	defaultFilename = "task.json"
)

func parseFile(file string) (*pb.Task, error) {
	var task *pb.Task
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &task)
	return task, err
}

func main() {
	cmd.Init()

	client := pb.NewTaskService("go.micro.srv.task", microclient.DefaultClient)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	task, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateTask(context.TODO(), &pb.CreateTaskRequest{Task: task})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r)

	tasks, err := client.ListTasks(context.Background(), &pb.ListTasksRequest{})
	if err != nil {
		log.Fatalf("Could not list tasks: %v", err)
	}
	for _, v := range tasks.Tasks {
		log.Println(v)
	}
}
