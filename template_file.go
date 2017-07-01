// +build !bip

package bip

import (
	"io"

	"github.com/flosch/pongo2"
)

// FileTemplate ...
type FileTemplate struct {
	fileName string
}

// Execute the template and returns the rendered template as a string.
func (t *FileTemplate) Execute(context Context) (string, error) {
	tpl, err := pongo2.FromFile(t.fileName)
	if err != nil {
		return "", err
	}
	return tpl.Execute(pongo2.Context(context))
}

// ExecuteBytes the template and returns the rendered template as a []byte.
func (t *FileTemplate) ExecuteBytes(context Context) ([]byte, error) {
	tpl, err := pongo2.FromFile(t.fileName)
	if err != nil {
		return nil, err
	}
	return tpl.ExecuteBytes(pongo2.Context(context))
}

// ExecuteWriter writes result that the template with the given context to writer.
func (t *FileTemplate) ExecuteWriter(context Context, w io.Writer) error {
	tpl, err := pongo2.FromFile(t.fileName)
	if err != nil {
		return err
	}
	return tpl.ExecuteWriter(pongo2.Context(context), w)
}
