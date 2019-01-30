package main

import (
    "os"
    "io/ioutil"
    "log"
    "path/filepath"
    "fmt"
)


type TreeItem interface {
    IsDir() bool
    Resolve()
}


type Properties struct {
    Name string
    Path string

    Parent *Dir
}


type File struct {
    Properties
}

func (f *File) IsDir() bool {
    return false
}

func (f *File) Resolve() {}


type Dir struct {
    Properties

    ContentItems []TreeItem
}

func (f *Dir) IsDir() bool {
    return true
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
    // TODO: sort.Sort(byInternalAndName(p.Deps))
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



func main() {
    t := Tree{}
    t.Resolve()
    fmt.Printf("%v", t)
}
