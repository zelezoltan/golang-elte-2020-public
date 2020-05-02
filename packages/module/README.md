# Go module example

1, Init in the root folder (no need to install, it's part of the go language tooling).

```
$ go mod init
```

2, Add the imports to your sources and run tidy to match your module files.
(Or run any go commands, they will also update/pull the missing deps.)

```
$ go mod tidy
```

3, Optionally you could pull all your deps into the vendor folder.

```
$ go mod vendor
```

Details: https://blog.golang.org/using-go-modules