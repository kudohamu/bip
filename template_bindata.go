// +build bip

package bip

import (
	"io"

	"github.com/flosch/pongo2"
)

type BindatTemplate struct {
	original *pongo2.Template
}

// Execute the template and returns the rendered template as a string.
func (t *BindatTemplate) Execute(context Context) (string, error) {
	return t.original.Execute(pongo2.Context(context))
}

// ExecuteBytes the template and returns the rendered template as a []byte.
func (t *BindatTemplate) ExecuteBytes(context Context) ([]byte, error) {
	return t.original.ExecuteBytes(pongo2.Context(context))
}

// ExecuteWriter writes result that the template with the given context to writer.
func (t *BindatTemplate) ExecuteWriter(context Context, w io.Writer) error {
	return t.original.ExecuteWriter(pongo2.Context(context), w)
}
