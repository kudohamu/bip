// +build bip

package bip

import "github.com/flosch/pongo2"

// FromFile loads a template from a filename and returns a Template instance.
func (ts *TemplateSet) FromFile(filename string) (Template, error) {
	data, err := ts.asset(filename)
	if err != nil {
		return nil, err
	}
	original, err := pongo2.FromBytes(data)
	if err != nil {
		return nil, err
	}

	return &BindatTemplate{
		original: original,
	}, nil
}
