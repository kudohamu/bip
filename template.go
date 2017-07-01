package bip

import "io"

// Template ...
type Template interface {
	Execute(context Context) (string, error)
	ExecuteBytes(context Context) ([]byte, error)
	ExecuteWriter(context Context, w io.Writer) error
}

// Must raises panic, if a Template couldn't successfully parsed.
func Must(tpl Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return tpl
}
