
```shell
curl -v http://localhost:8080/todos
```

```shell
curl -v http://localhost:8080/todos/5d146d915992aa1941069611
```

```shell
curl -v -X POST -d '{"content": "gogo"}' http://localhost:8080/todos
```

```shell
curl -v -X PUT -d '{"id": "5d146d915992aa1941069611", "content": "wuwulalaaaaa"}' http://localhost:8080/todos
```

```shell
curl -v -X DELETE http://localhost:8080/todos/5d146d915992aa1941069611
```