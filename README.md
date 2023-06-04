# jsontemplate [![Go Report Card](https://goreportcard.com/badge/github.com/wolfeidau/jsontemplate)](https://goreportcard.com/report/github.com/wolfeidau/jsontemplate) [![Documentation](https://godoc.org/github.com/wolfeidau/jsontemplate?status.svg)](https://godoc.org/github.com/wolfeidau/jsontemplate)

This library provides a way to template JSON using paths extracted from another json document. This uses [github.com/json-iterator/go](https://github.com/json-iterator/go) to parse the document with marshalling it, only the fields needed are extracted. In addition to extracting fields from the current document, it will also attempt to parse and extract fields from a nested JSON document if the path requires.

# Example

So we have an event come through, we extract a few fields and insert them into another JSON document.

```go
import (
	"context"
	"fmt"

	"github.com/wolfeidau/jsontemplate"
)

func ExampleTemplate_ExecuteToString() {
	template := `{
  "name": ${msg.name},
  "age": ${msg.age},
  "cyclist": ${msg.cyclist}
}`

	tpl, _ := jsontemplate.NewTemplate(template)

	res, _ := tpl.ExecuteToString([]byte(`{"msg":{"name":"markw","age":23,"cyclist":true}}`))
	fmt.Println(res)
	// Output:
	// {
	//   "name": "markw",
	//   "age": 23,
	//   "cyclist": true
	// }
}
```

# Performance

Given I expect to run this for large numbers of events I have attempted to keep the code very simple and capitalise on the work done in the underlying libraries.

```
go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/wolfeidau/jsontemplate
BenchmarkTemplate_Execute-10    	  527289	      2027 ns/op	    4655 B/op	     129 allocs/op
PASS
ok  	github.com/wolfeidau/jsontemplate	1.264s
```

# License

This project is released under Apache 2.0 license and is copyright [Mark Wolfe](https://www.wolfe.id.au).