golang version
```go
go 1.22.4
```
Start
```go
cd time_tracker
go mod download
```
Run API
```go
go run time_tracker/cmd/main.go
```

```
.
├── cmd
│   └── main.go
├── config
│   └── config.go
├── controllers
│   ├── task_controller.go
│   └── user_controller.go
├── models
│   ├── task.go
│   └── user.go
├── repositories
│   ├── task_repository.go
│   └── user_repository.go
├── routers
│   └── router.go
├── services
│   ├── task_service.go
│   └── user_service.go
├── utils
│   ├── logger.go
│   └── external_api.go
├── migrations
│   └── init.sql
├── .env
├── go.mod
├── go.sum
└── swagger.yaml

```
