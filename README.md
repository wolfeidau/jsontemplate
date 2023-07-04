# jsontemplate [![Go Report Card](https://goreportcard.com/badge/github.com/wolfeidau/jsontemplate)](https://goreportcard.com/report/github.com/wolfeidau/jsontemplate) [![Documentation](https://godoc.org/github.com/wolfeidau/jsontemplate?status.svg)](https://godoc.org/github.com/wolfeidau/jsontemplate)

This library provides a way to template a [JSON](https://www.json.org/) document using paths extracted from another, typically larger, JSON document. This uses [github.com/json-iterator/go](https://github.com/json-iterator/go) to parse the document without marshalling it, only the fields needed are extracted. In addition to extracting fields from the source document, it will also attempt to parse and extract fields from nested JSON document, stored in a string field, if the path requires.

# What can it be used for?

I use this library to reduce large JSON documents into smaller JSON documents that contain only the data needed for a specific purpose.

Examples of this are the following:

* S3 event notifications contain a large JSON payload with details of the event, this can be templated into a smaller JSON payload containing only the data needed by a Lambda function.
* API Gateway request payloads received by webhooks often contain large JSON payloads, these can be templated into smaller JSON payloads containing only the data needed by a downstream service.

# Example

So we have an event come through, we extract a few fields and insert them into another JSON document.

```go
func ExampleTemplate_ExecuteToString() {

	tpl, _ := jsontemplate.NewTemplate(`{"name": "${msg.name}","age": "${msg.age}","cyclist": "${msg.cyclist}"}`)

	res, _ := tpl.ExecuteToString([]byte(`{"msg":{"name":"markw","age":23,"cyclist":true}}`))
	fmt.Println(res)
	// Output:
	// {"name": "markw","age": 23,"cyclist": true}

}
```

## tags

To customise how the JSON is rendered, you can use tags which are provided after the path and delimited by `;`:

* `escape`, this will escape the JSON output.

```go
func ExampleTemplate_ExecuteToString_encoded() {

	tpl, _ := jsontemplate.NewTemplate(`{"msg": "${msg;escape}"}`)

	res, _ := tpl.ExecuteToString([]byte(`{"msg":{"name":"markw","age":23,"cyclist":true}}`))
	fmt.Println(res)
	// Output:
	// {"msg": "{\"name\":\"markw\",\"age\":23,\"cyclist\":true}"}
}
```

# Performance

Given I expect to run this for large numbers of events I have attempted to keep the code very simple and capitalize on the work done in the underlying libraries.

```
go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/wolfeidau/jsontemplate
BenchmarkTemplate_ExecuteToString-10    	  676122	      2105 ns/op	    4767 B/op	     134 allocs/op
BenchmarkTemplate_Execute-10            	  559828	      2180 ns/op	    4876 B/op	     136 allocs/op
PASS
ok  	github.com/wolfeidau/jsontemplate	3.001s
```

# License

This project is released under Apache 2.0 license and is copyright [Mark Wolfe](https://www.wolfe.id.au).