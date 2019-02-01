package main

import (
    "os"
    "io/ioutil"
    "log"
    "path/filepath"
    "github.com/davecgh/go-spew/spew"
    "fmt"
    "io"
    "strings"
)

// Todo: -f flag
// Todo: 'empty' size if a file is empty
// Todo: passing tests
// Todo: refactor it all to packages
// Todo: sorting filenames in dirs
// HELP HERE: https://github.com/KyleBanks/depth

const (
	outputPadding    = "│	"
	outputPrefix     = "├───"
	outputPrefixLast = "└───"
)


type TreeItem interface {
    Resolve()
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

func (f *File) Resolve() {}

func (f *File) ToString() string {
    return fmt.Sprintf("%s (%db)", f.Properties.Name, f.Size)
}

func (f *File) GetChildren() []TreeItem {
    return []TreeItem{}
}


type Dir struct {
    Properties

    ContentItems []TreeItem
}


func (dir *Dir) AddContentItems(files []os.FileInfo) {
    for _, file := range files {
        var item TreeItem
        path := filepath.Join(dir.Path, file.Name())
        name := file.Name()
        if file.IsDir() {
            item = &Dir {
                Properties: Properties {
                    Name: name,
                    Path: path,
                    Parent: dir,
                },
            }
        } else {
            item = &File {
                Size: file.Size(),
                Properties: Properties{
                    Name: name,
                    Path: path,
                },
            }
        }

        item.Resolve()

        dir.ContentItems = append(dir.ContentItems, item)
    }
}

func (dir *Dir) Resolve() {
    path := dir.Path
    files, err := ioutil.ReadDir(path)
    if err != nil {
        log.Printf("Error occured reading %s: %s", path, err)
    }

    dir.AddContentItems(files)
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

    IncludeFiles bool
}

func (tree *Tree) Resolve() {
    pwd, err := os.Getwd()
    if err != nil {
		log.Printf("Error occured reading %s: %s", pwd, err)
	}

    tree.Root = &Dir {
        Properties: Properties {
            Path: pwd,
        },
    }
    tree.Root.Resolve()
}

func OutputTree(writer io.Writer, file TreeItem, indentCount int, isLast bool) {

    indentSubstring := strings.Repeat(outputPadding, indentCount)
    prefixSubstring := outputPrefix
    if isLast {
        prefixSubstring = outputPrefixLast
    }

    _, err := fmt.Fprintf(writer, "%s%s%s\n", indentSubstring, prefixSubstring, file.ToString())
    if err != nil {
		log.Printf("Can not output the data: %s", err)
	}

    for index, item := range(file.GetChildren()) {
        itemIsLast := index == len(item.GetChildren()) - 1
	    OutputTree(writer, item, indentCount + 1, itemIsLast)
	}
}


func main() {`
    t := Tree{}
    t.Resolve()
    spew.Printf("%v\n\n\n", t)

    OutputTree(os.Stdout, t.Root, 0, false)
}
