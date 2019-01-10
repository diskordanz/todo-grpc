run:
		docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns user-service

build:
		GOOS=linux GOARCH=amd64 go build -o ./task-service/task-service -i ./task-service/*.go
		GOOS=linux GOARCH=amd64 go build -o ./task-cli/task-cli -i ./task-cli/*.go
		GOOS=linux GOARCH=amd64 go build -o ./user-cli/user-cli -i ./user-cli/*.go
		GOOS=linux GOARCH=amd64 go build -o ./user-service/user-service -i ./user-service/*.go

		docker build -t task-service ./task-service
		docker build -t task-cli ./task-cli
		docker build -t user-cli ./user-cli
		docker build -t user-service ./user-service

		rm ./task-service/task-service
		rm ./task-cli/task-cli
		rm ./user-cli/user-cli
		rm ./user-service/user-service

run:
		docker-compose up

down:
		docker-compose down

clean:
		docker rm task-service user-service task-cli user-cli 

protobuf:
		protoc --proto_path=GOPATH/src/github.com/diskordanz/todo-grpc/task-service:. --micro_out=. --go_out=. task-service/proto/task.proto
		protoc --proto_path=GOPATH/src/github.com/diskordanz/todo-grpc/user-service:. --micro_out=. --go_out=. user-service/proto/user.proto
		#protoc -I. --go_out=plugins=micro:. task-service/proto/task.proto
		#protoc -I. --go_out=plugins=micro:. user-service/proto/user.proto
