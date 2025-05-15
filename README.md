# rabbitmq-go

## Installing RabbitMQ
```
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4-management
```

Go to `localhost:15672` for management page

## Running hello example
Start the receiving
```
go run cmd/hello/receiving/main.go
```

Start the sending
```
go run cmd/hello/send/main.go
```

## Running work queue example

Run new task
```
go run cmd/workqueue/newtask/main.go message content one .....
go run cmd/workqueue/newtask/main.go message content two .......
go run cmd/workqueue/newtask/main.go message content three ......
```

Run worker
```
go run cmd/workqueue/worker/main.go
```

## Running pub/sub example

```
go run cmd/pubsub/receive_logs/main.go &> logs_from_rabbit.log
```

```
go run cmd/pubsub/receive_logs/main.go
```

```
go run cmd/pubsub/emit_log/main.go
```
