package jsontemplate_test

import (
	"fmt"

	"github.com/wolfeidau/jsontemplate"
)

func ExampleTemplate_ExecuteToString() {

	tpl, _ := jsontemplate.NewTemplate(`{"name": "${msg.name}","age": "${msg.age}","cyclist": "${msg.cyclist}"}`)

	res, _ := tpl.ExecuteToString([]byte(`{"msg":{"name":"markw","age":23,"cyclist":true}}`))
	fmt.Println(res)
	// Output:
	// {"name": "markw","age": 23,"cyclist": true}
}

func ExampleTemplate_ExecuteToString_encoded() {

	tpl, _ := jsontemplate.NewTemplate(`{"msg": "${msg;escape}"}`)

	res, _ := tpl.ExecuteToString([]byte(`{"msg":{"name":"markw","age":23,"cyclist":true}}`))
	fmt.Println(res)
	// Output:
	// {"msg": "{\"name\":\"markw\",\"age\":23,\"cyclist\":true}"}
}
