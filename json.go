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
	ft      *fasttemplate.Template
	encoder Encoder
}

func NewTemplate(content string) (*Template, error) {
	ft, err := fasttemplate.NewTemplate(content, "${", "}")
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	return &Template{ft: ft, encoder: JSONEncoder}, nil
}

// ExecuteToString executes the template against the given JSON event
// and returns the result as a string.
func (t *Template) ExecuteToString(evt []byte) (string, error) {
	return t.ft.ExecuteFuncStringWithErr(eventJSONTagFunc(evt, t.encoder))
}

// Execute executes the template against the given JSON event
// and writes the result to the given writer. It returns the number of
// bytes written and any error.
func (t *Template) Execute(wr io.Writer, evt []byte) (int64, error) {
	return t.ft.ExecuteFunc(wr, eventJSONTagFunc(evt, t.encoder))
}

func eventJSONTagFunc(evt []byte, encode Encoder) fasttemplate.TagFunc {
	doc := NewDocument(evt)

	return func(wr io.Writer, tag string) (int, error) {
		result, err := doc.Read(tag)
		if err != nil {
			return 0, fmt.Errorf("failed to read field: %w", err)
		}

		buf := new(bytes.Buffer)

		err = encode(buf, result)
		if err != nil {
			return 0, fmt.Errorf("failed to encode result: %w", err)
		}

		return wr.Write(bytes.TrimSuffix(buf.Bytes(), []byte("\n")))
	}
}

// Encoder encodes the given value v to the given writer wr.
//
// It returns any error encountered during encoding.
type Encoder func(wr io.Writer, v any) error

// JSONEncoder encodes the given value v to the given writer wr using the JSON encoding format.
//
// It returns any error encountered during encoding.
func JSONEncoder(wr io.Writer, v any) error {
	err := jsoniterator.NewEncoder(wr).Encode(v)
	if err != nil {
		return fmt.Errorf("failed to JSON encode result: %w", err)
	}

	return nil
}
