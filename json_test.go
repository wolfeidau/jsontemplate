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

	type args struct {
		evt []byte
	}
	tests := []struct {
		name            string
		templateStr     string
		args            args
		wantResult      string
		wantTemplateErr bool
		wantErr         bool
	}{
		{
			name:        "should return number",
			templateStr: `${data.counts.2}`,
			args: args{
				evt: []byte(`{"data": {"counts": [1,2,12]}}`),
			},
			wantResult: `12`,
		},
		{
			name:        "should return array",
			templateStr: `{"names": ${data.names}}`,
			args: args{
				evt: []byte(`{"data": {"names": ["a", "b", "c"]}}`),
			},
			wantResult: `{"names": ["a", "b", "c"]}`,
		},
		{
			name:        "should return error for invalid payload",
			templateStr: `${data.counts.2}`,
			args: args{
				evt: []byte(`{"data": {"counts"": [1,2,12]}}`),
			},
			wantErr: true,
		},
		{
			name:        "should return error for invalid template",
			templateStr: `${`,
			args: args{
				evt: []byte(`{"data": {"counts": [1,2,12]}}`),
			},
			wantTemplateErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)

			tpl, err := NewTemplate(tt.templateStr)
			if tt.wantTemplateErr {
				assert.Error(err)
				return
			}

			assert.NoError(err)

			buf := new(bytes.Buffer)

			n, err := tpl.Execute(buf, tt.args.evt)
			if tt.wantErr {
				assert.Error(err)
				return
			}

			assert.NoError(err)
			assert.True(n > 0, "expected bytes written > 0")
			assert.JSONEq(tt.wantResult, buf.String())
		})
	}

}

func BenchmarkTemplate_ExecuteToString(b *testing.B) {

	content := []byte(`{"msg":{"name":"markw","age":23,"cyclist":true}}`)
	expectedResult := `{"data":"markw",count:23,"flag":true}`

	tpl, err := NewTemplate(`{"data":${msg.name},count:${msg.age},"flag":${msg.cyclist}}`)
	if err != nil {
		b.Fatalf("error in template: %s", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			res, err := tpl.ExecuteToString(content)
			if err != nil {
				b.Fatalf("unexpected error: %s", err)
			}
			if res != expectedResult {
				b.Fatalf("unexpected result\n%q\nExpected\n%q\n", res, expectedResult)
			}
		}
	})
}
