package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"../../../tree" // fixme: will fixed for new repo
)

// Todo: sorting filenames in dirs
// Todo: fix root folder showing
// Todo: passing tests

const (
	outputPadding     = "│	"
	outputPaddingLast = "	"
	outputPrefix      = "├───"
	outputPrefixLast  = "└───"
)

func outputTree(writer io.Writer, file tree.TreeItem, parentIndent string, isLast bool, parentIsLast bool) {

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
		outputTree(writer, item, indentSubstring, itemIsLast, isLast)
	}
}

func initTree() *tree.Tree {
    var t tree.Tree
    flag.BoolVar(&t.IncludeFiles, "f", false, "Set it if files should be included too")

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
    t.Pwd = pwd

    return &t
}

func main() {
	t := initTree()
	t.Resolve()

	outputTree(os.Stdout, t.Root, "", false, false)
}
