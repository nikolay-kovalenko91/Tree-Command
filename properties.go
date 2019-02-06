package tree

type Properties struct {
	Name string
	Path string

	Parent *Dir
}

func (p *Properties) HasRootParent() bool {
    return p.Parent.Parent == nil
}
