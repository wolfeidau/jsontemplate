package jsontemplate

import "testing"

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
			wantResult: 23,
		},
		{
			name:       "should extract boolean field in document",
			args:       args{path: "data.set"},
			content:    []byte(`{"data":{"set":true}}`),
			wantResult: true,
		},
		{
			name:       "should extract numeric field from array in document",
			args:       args{path: "data.counts.0"},
			content:    []byte(`{"data":{"counts":[23]}}`),
			wantResult: 23,
		},
		{
			name:       "should extract object field from document",
			args:       args{path: "data"},
			content:    []byte(`{"data":{"counts":[23]}}`),
			wantResult: `{"counts":[23]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDocument(tt.content)
			gotResult, err := d.Read(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Document.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("Document.Read() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
