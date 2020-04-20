# Complex example

Here you see a bit more complex example with global context handling like db connection and logger.

I'll use some 3rd party libraries here and a db connection:
- [Chi](https://github.com/go-chi/chi) - Easier HTTP routing
- [MySQL driver](https://github.com/go-sql-driver/mysql) - MySQL driver for the built in `database/sql` interfaces
- [SQLx](https://github.com/jmoiron/sqlx) - Small additional lib for easier SQL queries (result auto mapping to structs)
- [Zap](https://github.com/uber-go/zap) - Better, structured logger 

DB setup I use here:
```
Server: 127.0.0.1
User: root
Pass: root
Database: testdb
```

DB Structure:

```sql
CREATE TABLE `message` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `message` text NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

Run the server:
```
$ go run httpserver/ex7/server.go httpserver/ex7/app.go
```


Do some HTTP calls:

```
$ curl http://127.0.0.1:8080/
$ curl http://127.0.0.1:8080/api
$ curl http://127.0.0.1:8080/api/list
$ curl -X POST http://127.0.0.1:8080/api/add -d '{"name":"test","message":""}'
$ curl -X POST http://127.0.0.1:8080/api/add -d '{"name":"test","message":"hello1"}'
$ curl -X POST http://127.0.0.1:8080/api/add -d '{"name":"test","message":"hello2"}'
$ curl http://127.0.0.1:8080/api/list
```
