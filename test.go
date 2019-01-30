package main

import (
    "fmt"
    "os"
    "io/ioutil"
     "log"
     "path/filepath"
)

// https://github.com/KyleBanks/depth/

func getFiles(dir string) []os.FileInfo {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

    return files
}

func investigateFiles(dir string, files []os.FileInfo) {
    for _, f := range files {
        if f.IsDir() {
            log.Println(filepath.Join(dir, f.Name()))
            dirTree(filepath.Join(dir, f.Name()))
        } else {
            fmt.Println(f.Name())
        }
    }
}

func dirTree(dir string) {
    files := getFiles(dir)
    investigateFiles(dir, files)
}

func main() {
    dirTree("./")
}
