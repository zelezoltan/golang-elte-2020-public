# Middleware magic

Here I'll use the same router library, because it's a bit easier to organize the endpoints.
Middlewares are custom http handler which you could chain together and do different checks and modification on the incoming (`http.Request`) data.

```
$ go run httpserver/ex6/server.go
```


## First example
The `/endpoint1` and `/endpoint2` are "protected" by a middleware, which checks the incoming Authorization header value. Check these call's on the endpoint:

```
$ curl 127.0.0.1:8080/endpoint1
$ curl -H "Authorization: test" 127.0.0.1:8080/endpoint1
$ curl -H "Authorization: someToken" 127.0.0.1:8080/endpoint1
$ curl -H "Authorization: wrong" 127.0.0.1:8080/endpoint2
$ curl -H "Authorization: someToken" 127.0.0.1:8080/endpoint2
```

## Second example

This time time the every endpoint under the `/api` endpoint is protected via the middleware.

Check these calls:

```
$ curl -H "Authorization: token2" 127.0.0.1:8080/api
$ curl -H "Authorization: token2" 127.0.0.1:8080/api/
$ curl -H "Authorization: token2" 127.0.0.1:8080/api/something
$ curl -H "Authorization: token2" 127.0.0.1:8080/api/something2
$ curl -H "Authorization: token2" "http://127.0.0.1:8080/api/something2?name=hello"
$ curl -H "Authorization: token2" "http://127.0.0.1:8080/api/something?name=hello"
```

## Third example

Here you can see how could you add some logger with middleware.

```
$ curl http://127.0.0.1:8080/longtime
$ for i in {1..5}; do curl http://127.0.0.1:8080/longtime &; done
```

Try this with a bit modification on the logger:
```
$ curl http://127.0.0.1:8080/longtime2
$ for i in {1..5}; do curl http://127.0.0.1:8080/longtime2 &; done
```