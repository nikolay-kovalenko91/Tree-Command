package tree

type Tree struct {
	Root *Dir
	Pwd  string

	IncludeFiles bool
}

func (tree *Tree) Resolve() {
	tree.Root = &Dir{
		Properties: Properties{
			Name: ".",
			Path: tree.Pwd,
		},
	}
	tree.Root.Resolve(tree.IncludeFiles)
}

type TreeItem interface {
	Resolve(bool)
	ToString() string
	GetChildren() []TreeItem
}

type Properties struct {
	Name string
	Path string
}
