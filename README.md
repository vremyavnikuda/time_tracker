### Time Tracker API

[golang download](https://go.dev/)

```
go version
```

```go
go 1.22
```

```go
cd time_tracker
go mod download
go mod tidy
```

Start

```go
go run main.go
```

```
├── controllers\
│   └── user_controller.go
│   └── time_controller.go
├── models\
│   └── user.go
│   └── time_entry.go
├── routes\
│   └── routes.go
├── docs\             
│   ├── docs.go
│   └── swagger.json
├── main.go
├── database\
│   └── database.go
├── go.mod
├── go.sum

```
