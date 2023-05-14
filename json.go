package jsontemplate

import (
	"bytes"
	"fmt"
	"io"

	jsoniterator "github.com/json-iterator/go"
	"github.com/valyala/fasttemplate"
)

// NewTemplate parses the given template content into a fasttemplate
// Template and returns it.
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

// ExecuteToString executes the template against the given JSON event
// and returns the result as a string.
func (t *Template) ExecuteToString(evt []byte) (string, error) {
	return t.ft.ExecuteFuncStringWithErr(eventJSONTagFunc(evt))
}

// Execute executes the template against the given JSON event
// and writes the result to the given writer. It returns the number of
// bytes written and any error.
func (t *Template) Execute(wr io.Writer, evt []byte) (int64, error) {
	return t.ft.ExecuteFunc(wr, eventJSONTagFunc(evt))
}

func eventJSONTagFunc(evt []byte) fasttemplate.TagFunc {
	doc := NewDocument(evt)

	return func(w io.Writer, tag string) (int, error) {
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
	}
}
