package jsontemplate

import (
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
