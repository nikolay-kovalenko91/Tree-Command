package tree

type Tree struct {
	Root *Dir
	Pwd  string

	IncludeFiles bool
}

func (tree *Tree) Resolve() {
	tree.Root = &Dir{
		Properties: Properties{
			Path: tree.Pwd,
		},
	}
	tree.Root.Resolve(tree.IncludeFiles)
}

type TreeItem interface {
	Resolve(bool)
	ToString() string
	GetChildren() []TreeItem
	HasRootParent() bool
}
