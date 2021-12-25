
```shell
docker run -d -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=<username> -e MONGO_INITDB_ROOT_PASSWORD=<password> mongo:4.4.10
export MONGODB_URL=mongodb://<username>:<password>@localhost:27017/todo_data?authSource=admin
```

```shell
curl -v http://localhost:8080/todos
[{"id":"5d146d915992aa1941069611","content":"gogo"}]
```

```shell
curl -v http://localhost:8080/todos/5d146d915992aa1941069611
{"id":"5d146d915992aa1941069611","content":"gogo"}
```

```shell
curl -v -X POST -d '{"content": "gogo"}' http://localhost:8080/todos
{"id":"5d146d915992aa1941069611","content":"gogo"}
```

```shell
curl -v -X PUT -d '{"id": "5d146d915992aa1941069611", "content": "wuwulalaaaaa"}' http://localhost:8080/todos
{"result":"success"}
```

```shell
curl -v -X DELETE http://localhost:8080/todos/5d146d915992aa1941069611
{"result":"success"}
```