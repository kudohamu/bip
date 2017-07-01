// +build !bip

package bip

import "github.com/flosch/pongo2"

// FromFile loads a template from a filename and returns a Template instance.
func (ts *TemplateSet) FromFile(filename string) (Template, error) {
	if _, err := pongo2.FromFile(filename); err != nil {
		return nil, err
	}

	return &FileTemplate{
		fileName: filename,
	}, nil
}
