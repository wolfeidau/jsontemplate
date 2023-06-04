package jsontemplate_test

import (
	"fmt"

	"github.com/wolfeidau/jsontemplate"
)

var template = `{
  "name": ${msg.name},
  "age": ${msg.age},
  "cyclist": ${msg.cyclist}
}`

func ExampleTemplate_ExecuteToString() {

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
