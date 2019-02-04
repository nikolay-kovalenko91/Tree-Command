package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"flag"
)

// Todo: fix root folder showing
// Todo: passing tests
// Todo: sorting filenames in dirs
// Todo: refactor it all to packages

// HELP HERE: https://github.com/KyleBanks/depth

const (
	outputPadding     = "│	"
	outputPaddingLast = "	"
	outputPrefix      = "├───"
	outputPrefixLast  = "└───"
)

type TreeItem interface {
	Resolve(bool)
	ToString() string
	GetChildren() []TreeItem
}

type Properties struct {
	Name string
	Path string

	Parent *Dir
}

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

type Dir struct {
	Properties

	ContentItems []TreeItem
}

func (dir *Dir) AddContentItems(files []os.FileInfo, includeFiles bool) {
	for _, file := range files {
		var item TreeItem
		path := filepath.Join(dir.Path, file.Name())
		name := file.Name()
		if file.IsDir() {
			item = &Dir{
				Properties: Properties{
					Name:   name,
					Path:   path,
					Parent: dir,
				},
			}
		} else {
		    if !includeFiles {
		        continue
		    }

            item = &File{
                Size: file.Size(),
                Properties: Properties{
                    Name: name,
                    Path: path,
                },
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
	// sort.Sort(byInternalAndName(p.Deps))
}

func (dir *Dir) ToString() string {
	return dir.Properties.Name
}

func (dir *Dir) GetChildren() []TreeItem {
	return dir.ContentItems
}

type Tree struct {
	Root *Dir
	Pwd string

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

func OutputTree(writer io.Writer, file TreeItem, parentIndent string, isLast bool, parentIsLast bool) {

	indentSubstring := fmt.Sprintf("%s%s", parentIndent, outputPadding)
	if parentIsLast {
		indentSubstring = fmt.Sprintf("%s%s", parentIndent, outputPaddingLast)
	}

	prefixSubstring := outputPrefix
	if isLast {
		prefixSubstring = outputPrefixLast
	}

	_, err := fmt.Fprintf(writer, "%s%s%s\n", indentSubstring, prefixSubstring, file.ToString())
	if err != nil {
		log.Printf("Can not output the data: %s", err)
	}

	fileChildren := file.GetChildren()
	for index, item := range fileChildren {
		itemIsLast := index == len(fileChildren) - 1
		OutputTree(writer, item, indentSubstring, itemIsLast, isLast)
	}
}

func initTree() *Tree {
    var tree Tree
    flag.BoolVar(&tree.IncludeFiles, "f", false, "Set it if files should be included too")

    flag.Parse()

    tail := flag.Args()
    var (
        pwd string
        err error
    )
    if len(tail) > 0 {
        pwd, err = filepath.Abs(tail[0])
        if err != nil {
            log.Fatal(err)
        }
    } else {
    	pwd, err = os.Getwd()
        if err != nil {
            log.Fatal(err)
        }
    }
    tree.Pwd = pwd

    return &tree
}

func main() {
	t := initTree()
	t.Resolve()

	OutputTree(os.Stdout, t.Root, "", false, false)
}
