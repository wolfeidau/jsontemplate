package jsontemplate

import (
	"fmt"
	"strconv"
	"strings"

	jsoniterator "github.com/json-iterator/go"
)

type path struct {
	tokens []any
}

// Parse takes a string path and returns a Path struct
// with the path tokens split on
func parse(t string) *path {

	strs := strings.Split(t, ".")

	tokens := make([]any, len(strs))

	for i := range tokens {
		v := strs[i]

		if intval, err := strconv.ParseInt(v, 10, 64); err == nil {
			tokens[i] = int(intval)
		} else {
			tokens[i] = v
		}
	}

	return &path{tokens: tokens}
}

// Next returns the next path token and updates the Path tokens.
func (p *path) next() (path any) {
	path, p.tokens = p.tokens[0], p.tokens[1:]
	return path
}

// HasNext checks if there are any remaining path tokens.
func (p *path) hasNext() bool {
	return len(p.tokens) > 0
}

// Document holds JSON content.
type Document struct {
	content []byte
}

// NewDocument returns a new Document.
func NewDocument(data []byte) *Document {
	return &Document{
		content: data,
	}
}

// Read reads a JSON value at a given path using jsoniter.Get.
// If an error occurs it returns the error it parses the path
// then iterates over the path otherwise it returns the string
// value at the path.
func (d *Document) Read(path string) (result any, err error) {

	p := parse(path)
	it := jsoniterator.Get(d.content)

	for {
		if it.ValueType() == jsoniterator.StringValue && p.hasNext() {
			it = jsoniterator.Get([]byte(it.ToString()))
		}

		token := p.next()
		it = it.Get(token)

		if err = it.LastError(); err != nil {

			if strings.Contains(err.Error(), "not found") {
				return nil, fmt.Errorf("token not found: [%s] while searching for path: %s", token, path)
			}

			return nil, err
		}

		// if we are done return iterator string
		if !p.hasNext() {
			switch it.ValueType() {
			case jsoniterator.NilValue:
				return nil, nil
			case jsoniterator.BoolValue:
				return it.ToBool(), nil
			case jsoniterator.NumberValue:
				return it.ToInt(), nil
			default:
				return it.ToString(), nil
			}
		}
	}
}
