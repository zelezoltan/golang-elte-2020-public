# Binary packing

Sometimes it's useful to pack some files (for example templates, static assets) into the binary.
There are multiple tools for that:
https://github.com/gobuffalo/packr/tree/master/v2
https://github.com/GeertJohan/go.rice

The trick here to add a `generate` step before the final build. These libs will generate (one or more) custom `.go` file which contains
your resource in a binary format and add some helpers to extract these on the run.

For example server CSS static files, embedded with `go.rice`:
```
box := rice.MustFindBox("cssfiles")
cssFileServer := http.StripPrefix("/css/", http.FileServer(box.HTTPBox()))
http.Handle("/css/", cssFileServer)
http.ListenAndServe(":8080", nil)
```
