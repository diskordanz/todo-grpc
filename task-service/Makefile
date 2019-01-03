build:	
		protoc --proto_path=GOPATH/src/github.com/diskordanz/todo-grpc/task-service:. --micro_out=. --go_out=. proto/task.proto
		GOOS=linux GOARCH=amd64 go build
		docker build -t task-service .
run:
		docker run -p 50051:50051 \
			-e MICRO_SERVER_ADDRESS=:50051 \
			-e MICRO_REGISTRY=mdns task-service