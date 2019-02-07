package main

import (
	"testing"
)

// Todo: implement me
func Test_outputs(t *testing.T) {
    dirTree(os.Stdout, []string{"./tree", "-f", "./"})
}
