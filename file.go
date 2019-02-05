package tree

import "fmt"

type File struct {
	Size int64

	Properties
}

func (f *File) Resolve(_ bool) {}

func (f *File) ToString() string {
	fSize := f.Size
	sizeSubstring := fmt.Sprintf("%db", fSize)
	if fSize == 0 {
		sizeSubstring = "empty"
	}

	return fmt.Sprintf("%s (%s)", f.Properties.Name, sizeSubstring)
}

func (f *File) GetChildren() []TreeItem {
	return []TreeItem{}
}
