package jsontemplate

import (
	"bytes"
	"context"
	"fmt"
	"io"

	jsoniterator "github.com/json-iterator/go"

	"github.com/valyala/fasttemplate"
)

type Template struct {
	ft *fasttemplate.Template
}

func NewTemplate(content string) (*Template, error) {
	ft, err := fasttemplate.NewTemplate(content, "${", "}")
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	return &Template{ft: ft}, nil
}

func (t *Template) ExecuteToString(ctx context.Context, evt []byte) (string, error) {

	doc := NewDocument(evt)

	return t.ft.ExecuteFuncStringWithErr(func(w io.Writer, tag string) (int, error) {
		result, err := doc.Read(tag)
		if err != nil {
			return 0, fmt.Errorf("failed to read field: %w", err)
		}

		buf := new(bytes.Buffer)

		err = jsoniterator.NewEncoder(buf).Encode(result)
		if err != nil {
			return 0, fmt.Errorf("failed to encode result: %w", err)
		}

		return w.Write(bytes.TrimSuffix(buf.Bytes(), []byte("\n")))
	})
}
