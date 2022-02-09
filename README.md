# restful-go

## Start Mongo DB

 - By Docker

```shell
docker run -d -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=<username> -e MONGO_INITDB_ROOT_PASSWORD=<password> mongo:4.4.10
export MONGODB_URL=mongodb://<username>:<password>@localhost:27017/todo_data?authSource=admin
```
- By MongoDB atlas

```shell
export MONGODB_URL=mongodb+srv://<username>:<password>@todo-data.7hhkw.mongodb.net/todo_data?retryWrites=true\&w=majority
```

## Endpoints

- Select all

```shell
curl -v http://localhost:8080/todos
[{"id":"5d146d915992aa1941069611","content":"gogo"}]
```

- Select one

```shell
curl -v http://localhost:8080/todos/5d146d915992aa1941069611
{"id":"5d146d915992aa1941069611","content":"gogo"}
```

- Add

```shell
curl -v -X POST -d '{"content": "gogo"}' http://localhost:8080/todos
{"id":"5d146d915992aa1941069611","content":"gogo"}
```

- Update

```shell
curl -v -X PUT -d '{"content": "wuwulalaaaaa"}' http://localhost:8080/todos/5d146d915992aa1941069611
{"result":"success"}
```

- Delete

```shell
curl -v -X DELETE http://localhost:8080/todos/5d146d915992aa1941069611
{"result":"success"}
```