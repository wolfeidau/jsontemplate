package jsontemplate_test

import (
	"context"
	"fmt"

	"github.com/wolfeidau/jsontemplate"
)

func Example() {
	template := `{
  "name": ${msg.name},
  "age": ${msg.age},
  "cyclist": ${msg.cyclist}
}`

	tpl, _ := jsontemplate.NewTemplate(template)

	res, _ := tpl.ExecuteToString(context.TODO(), []byte(`{"msg":{"name":"markw","age":23,"cyclist":true}}`))
	fmt.Println(res)
	// Output:
	// {
	//   "name": "markw",
	//   "age": 23,
	//   "cyclist": true
	// }
}
