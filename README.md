# Todo App
**Note**
This project was created on the basis of video tutorials by [@zhashkevych](https://github.com/zhashkevych)

## Migrate
```sh
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
```

## Setup
Setup PostgreSQL environment with Docker
```sh
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres:alpine
```

## Testing
You can quickly test this REST API with ready-made Postman collections. [Run in Postman](https://app.getpostman.com/run-collection/ac6d5a80226c1540debb?action=collection%2Fimport)



