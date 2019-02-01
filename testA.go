package main

import (
    "os"
    "io/ioutil"
    "log"
    "path/filepath"
    "github.com/davecgh/go-spew/spew"
)


type Properties struct {
    Name string
    Path string

    Parent *Dir
}

type File struct {
    Properties
}

type Dir struct {
    Properties

    ContentItems []TreeItem
}

type TreeItem interface {
    IsDir() bool
    GetChildren() []TreeItem
}


type Tree struct {
    Root *Dir

    IncludeFiles bool
}

func (f *File) IsDir() bool {
    return false
}

func (f *Dir) IsDir() bool {
    return true
}

func (tree *Tree) Resolve() {
    var queue []TreeItem
    pwd, err := os.Getwd()
    if err != nil {
		log.Printf("Error occured reading %s: %s", pwd, err)
	}

    tree.Root = &Dir {
        Properties: Properties {
            Path: pwd,
        },
    }

    queue = append(queue, tree.Root)
    for _, element := range queue {
        children := element.GetChildren()
        queue = append(queue, children...)
    }
}

func (f *File) GetChildren() []TreeItem {
    return []TreeItem{}
}

func (dir *Dir) GetChildren() []TreeItem {
    path := dir.Path
    files, err := ioutil.ReadDir(path)
    if err != nil {
        log.Printf("Error occured reading %s: %s", path, err)
    }

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
                Properties: Properties{
                    Name: name,
                    Path: path,
                },
            }
        }

        dir.ContentItems = append(dir.ContentItems, item)
    }

    return dir.ContentItems
}

func main() {
    t := Tree{}
    t.Resolve()
    spew.Printf("%v", t)
}
