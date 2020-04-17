# HTTPS server

This is the manual solution for just a simple example.
Don't really do this, it'll be only a self-signed cert, but you could add any certificates with the same way.
(But nowadays this whole process could be automatized with the use of `Let's Encrypt` and some libraries like: https://github.com/caddyserver/certmagic or https://github.com/gin-gonic/autotls)


**Sidenote:** I've added a generated cert, so it works and you don't have to do the following step.

How to generate a self-signed TLS cert by hand:

```
$ cd httpserver/ex3/
$ openssl req -x509 -nodes -newkey rsa:2048 -keyout server.key -out server.crt -days 3650
```

Run the server:

```
$ go run httpserver/ex3/server.go
```

Check the following URL: https://127.0.0.1:8080/
You'll see an error, because of the self-signed certs are not trusted by default, but if you still force your browser you should se the handler's response.
