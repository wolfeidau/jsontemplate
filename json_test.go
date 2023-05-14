package jsontemplate

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecute(t *testing.T) {
	assert := require.New(t)

	tpl, err := NewTemplate(`{"data":${msg.name},count:${msg.age},"flag":${msg.cyclist}}`)
	assert.NoError(err)

	res, err := tpl.ExecuteToString([]byte(`{"msg":{"name":"markw","age":23,"cyclist":true}}`))
	assert.NoError(err)
	assert.Equal(`{"data":"markw",count:23,"flag":true}`, res)
}

func TestTemplate_Execute(t *testing.T) {
	assert := require.New(t)

	content := `{"greeting":${name}}`
	tmpl, err := NewTemplate(content)
	assert.NoError(err)

	evt := []byte(`{"name": "John"}`)
	buf := new(bytes.Buffer)
	n, err := tmpl.Execute(buf, evt)
	assert.NoError(err)

	expectedDoc := `{"greeting":"John"}`

	assert.Equal(int64(len(expectedDoc)), n)
	assert.Equal(expectedDoc, buf.String())
}
