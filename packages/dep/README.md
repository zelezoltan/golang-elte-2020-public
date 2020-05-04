# DEP example

1, Install dep. https://github.com/golang/dep

2, Init in in the project root

```
$ dep init
```

3, Add packaged you want to use (or just add the import then run the `dep ensure`)

```
$ dep ensure -add github.com/google/uuid
Feching sources...

"github.com/google/uuid" is not imported by your project, and has been temporarily added to Gopkg.lock and vendor/.
If you run "dep ensure" again before actually importing it, it will disappear from Gopkg.lock and vendor/.
```
