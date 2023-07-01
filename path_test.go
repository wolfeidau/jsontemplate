package jsontemplate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDocument_Read(t *testing.T) {

	type args struct {
		path string
	}
	tests := []struct {
		name       string
		content    []byte
		args       args
		wantResult any
		wantErr    bool
	}{
		{
			name:       "should extract name from inner document",
			args:       args{path: "msg.inner.name"},
			content:    []byte(`{"msg":{"inner":"{\"name\": \"mark\"}"}}`),
			wantResult: "mark",
		},
		{
			name:       "should extract inner document",
			args:       args{path: "msg.inner"},
			content:    []byte(`{"msg":{"inner":"{\"name\": \"mark\"}"}}`),
			wantResult: `{"name": "mark"}`,
		},
		{
			name:    "should error on missing field in document",
			args:    args{path: "msg.nothere"},
			content: []byte(`{"msg":{"inner":"{\"name\": \"mark\"}"}}`),
			wantErr: true,
		},
		{
			name:       "should extract numeric field in document",
			args:       args{path: "data.count"},
			content:    []byte(`{"data":{"count":23}}`),
			wantResult: float64(23),
		},
		{
			name:       "should extract boolean field in document",
			args:       args{path: "data.set"},
			content:    []byte(`{"data":{"set":true}}`),
			wantResult: true,
		},
		{
			name:       "should extract null field in document",
			args:       args{path: "data.set"},
			content:    []byte(`{"data":{"set":null}}`),
			wantResult: nil,
		},
		{
			name:       "should extract numeric field from array in document",
			args:       args{path: "data.counts.0"},
			content:    []byte(`{"data":{"counts":[23]}}`),
			wantResult: float64(23),
		},
		{
			name:       "should extract object field from document",
			args:       args{path: "data"},
			content:    []byte(`{"data":{"counts":[23]}}`),
			wantResult: map[string]interface{}(map[string]interface{}{"counts": []interface{}{float64(23)}}),
		},
		{
			name:       "should extract string field from array in document",
			args:       args{path: "data.counts.0"},
			content:    []byte(`{"data":{"counts":[23]}}`),
			wantResult: float64(23),
		},
		{
			name:       "should extract string field from document",
			args:       args{path: "data"},
			content:    []byte(`{"data":{"names":["a","b","c"]}}`),
			wantResult: map[string]interface{}(map[string]interface{}{"names": []interface{}{"a", "b", "c"}}),
		},
		{
			name:       "should extract string field from document",
			args:       args{path: "data.names.0"},
			content:    []byte(`{"data":{"names":["a","b","c"]}}`),
			wantResult: "a",
		},
		{
			name:       "should extract string field from document",
			args:       args{path: "data;escape"},
			content:    []byte(`{"data":{"names":["a","b","c"]}}`),
			wantResult: `{"names":["a","b","c"]}`,
		},
		{
			name:    "should extract string field from document",
			args:    args{path: "data;nope"},
			content: []byte(`{"data":{"names":["a","b","c"]}}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := require.New(t)
			d := NewDocument(tt.content)
			gotResult, err := d.Read(tt.args.path)
			if tt.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}

			assert.Equal(tt.wantResult, gotResult)
		})
	}
}
