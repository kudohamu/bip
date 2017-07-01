package bip

// TemplateSet is group of template.
type TemplateSet struct {
	asset func(path string) ([]byte, error)
}

// NewSet create TemplateSet using bindata.
func NewSet(asset func(path string) ([]byte, error)) *TemplateSet {
	return &TemplateSet{
		asset: asset,
	}
}
