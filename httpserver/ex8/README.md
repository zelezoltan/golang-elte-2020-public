# HTTP testing example

This is the same app which we used in `ex7` with some modifications.

I'll use some 3rd party libraries here and a db connection:
- [Chi](https://github.com/go-chi/chi) - Easier HTTP routing
- [SQLx](https://github.com/jmoiron/sqlx) - Small additional lib for easier SQL queries (result auto mapping to structs)
- [Zap](https://github.com/uber-go/zap) - Better, structured logger
- [SQLite driver](https://github.com/mattn/go-sqlite3) - SQLite driver for the built in `database/sql` interfaces


To test our code we have multiple ways to do it in multiple parts of our code:
- Use the power of interfaces and the [testing](https://golang.org/pkg/testing/) package to create unit tests
- Use [httptest](https://golang.org/pkg/net/http/httptest/) package to create "fake" server or client code (useful for integrationt tests)
- Mock the sql driver itself with [https://github.com/DATA-DOG/go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) (There are multiple libs for this)
- Mock everything we need (use with proper interfaces) [https://github.com/golang/mock](https://github.com/golang/mock)
- Add a bit smarter assertion with [https://github.com/stretchr/testify](https://github.com/stretchr/testify)

I won't go into every details in every package, but it's a good starting point to check these and see the examples in their docs.


For mock generation I'll use the  [mockery](https://github.com/vektra/mockery), but you could also use [gomock](https://github.com/golang/mock).
(Mockery works perfectly with the testify packages, because it'll gen with the help of the [testify's mock package](https://github.com/stretchr/testify#mock-package).)


## Gomock usage example
Installation:
```
$ go get github.com/golang/mock/mockgen@latest
```

If you already added your Go bin folder to your path, you could call it simply like this (in the `httpserver/ex8` folder):
```
$ mockgen -source=app/app.go -destination=app/app_mock_test.go -package=app database
```

It'll load the `app/app.go` and generate the `database` interface's mock into the `app/app_mock_test.go` file. 
Without the `-destination` flag it'll print out the results.
(You could also add helper for this mock generation with Go's built in code generation tool: [https://blog.golang.org/generate](https://blog.golang.org/generate))
Without the `-package` flag it'll generate the package name `mock_app` which causes an error at the next compile.

Running the `mockgen` multiple times if the output file exists will cause an error: `2020/04/25 11:51:42 Loading input failed: loading package failed`.
So before you run the command make sure you've deleted the `app/app_mock_test.go` file first.

## Mockery usage example

Install:
```
$ go get github.com/vektra/mockery/.../
```

If you already added your Go bin folder to your path, you could call it simply like this (in the `httpserver/ex8` folder):
```
$ mockery -name=database -dir=app -output=app -inpkg
```

The `-name` will tell which interface to generate, the `-dir` flag helps where to find the interface definition,
the `-output` tells where to generate the result and the `-inpkg` will force the generate code to have the same package name as the interface.

## Run the tests

Run the tests:
```
$ cd httpserver/ex8
$ go test -v ./...
```

Or run the server:
```
$ go run -v server.go
```

The `-v` will give you a detailed result.

Or you could create a coverage report:

```
$ go test -v -cover ./...
```