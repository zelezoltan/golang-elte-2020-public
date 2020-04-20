# Simple HTTP server


```
$ go run httpserver/ex1/server.go 
```

Now just open up a web browser and check this URL: [http://127.0.0.1:8080/?name=Bob](http://127.0.0.1:8080/?name=Bob)

Check the logs in the console.
Now stop the server (Ctrl+C) then uncomment the line 22 (where you could see the `r.ParseForm()` part) and start again.
Check what's the difference.


**Sidenote:** As you see there are 2 HTTP calls, because the browser tries to load the `favicon.ico` too but it's unavailable.


Now try the following URL's and check which handler will catch your call:
- [http://127.0.0.1:8080/somethin?name=Bob2](http://127.0.0.1:8080/somethin?name=Bob2)
- [http://127.0.0.1:8080/other](http://127.0.0.1:8080/other)
- [http://127.0.0.1:8080/other2](http://127.0.0.1:8080/other2)
- [http://127.0.0.1:8080/another](http://127.0.0.1:8080/another)
- [http://127.0.0.1:8080/another/hello](http://127.0.0.1:8080/another/hello)

As you see the `/` at the end of the handler pattern has a bit special meaning.