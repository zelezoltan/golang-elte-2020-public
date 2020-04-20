# Custom mux

I'll use (chi)[https://github.com/go-chi/chi] here.
Right now the package management is not important, I "vendored" everything, but I add the commands what I did, but don't have to repeate it.

This pulls every 3rd party package into a folder called `vendor`.
The Go runtime checks that specific folder if it could not find the package in the standard library. 
(Also `go mod` could pull everything on the fly, but more about this later.)

```
$ cd httpserver/ex5
$ go mod init
$ go mod vendor
```

You could just run the server:
```
$ go run server.go
```


Check these URLs:
- http://127.0.0.1:8080/
- http://127.0.0.1:8080/aaaaaa
- http://127.0.0.1:8080/aaaBaa