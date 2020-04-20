<html>
<body>
<h1>{{ .SomeTitle }}</h1>
<ul>
{{range $val := .SomeValues}}
     <li>{{ $val.Key }} -> {{$val.Val}}</li>
{{end}}
</ul>
</body>
</html>