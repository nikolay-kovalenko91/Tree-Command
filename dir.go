package tree

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Dir struct {
	Properties

	ContentItems []TreeItem
}

func (dir *Dir) AddContentItems(files []os.FileInfo, includeFiles bool) {
	for _, file := range files {
		var item TreeItem
		properties := Properties{
			Name: file.Name(),
			Path: filepath.Join(dir.Path, file.Name()),
			Parent: dir,
		}

		if file.IsDir() {
			item = &Dir{Properties: properties}
		} else {
			if !includeFiles {
				continue
			}

			item = &File{
				Size: file.Size(),
				Properties: properties,
			}
		}

		item.Resolve(includeFiles)

		dir.ContentItems = append(dir.ContentItems, item)
	}
}

func (dir *Dir) Resolve(includeFiles bool) {
	path := dir.Path
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("Error occured reading %s: %s", path, err)
	}

	dir.AddContentItems(files, includeFiles)
}

func (dir *Dir) ToString() string {
	return dir.Properties.Name
}

func (dir *Dir) GetChildren() []TreeItem {
	return dir.ContentItems
}
