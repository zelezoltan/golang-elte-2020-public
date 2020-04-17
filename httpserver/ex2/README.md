# Simple server with form handling


```
$ go run httpserver/ex2/server.go
```

Use curl and do some calls:

```
$ curl 127.0.0.1:8080/post
$ curl -X post 127.0.0.1:8080/post #TODO: fix this
$ curl -X POST 127.0.0.1:8080/post
$ curl -X POST 127.0.0.1:8080/post -F 'name=test' -F 'password=test'
$ curl -X POST 127.0.0.1:8080/post -F 'name=test' -F 'password=testPass'
```

Now try the json endpoint:

```
$ curl -X POST 127.0.0.1:8080/postjson
$ curl -X POST 127.0.0.1:8080/postjson -d '{}'
$ curl -X POST 127.0.0.1:8080/postjson -d '{"name":"test","password":"testPass"}'
```
