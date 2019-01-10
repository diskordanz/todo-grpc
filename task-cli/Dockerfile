FROM debian:latest

RUN mkdir -p /app
WORKDIR /app

ADD task-cli /app
ADD task.json /app

ENTRYPOINT ["./task-cli", "task.json"]
#EXPOSE 9092

CMD ["./task-cli"]
