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


