package main

import (
	"bytes"
	"testing"
)

const fullTreeResult = `├───project
│	├───file.txt (19b)
│	└───gopher.png (70372b)
├───static
│	├───a_lorem
│	│	├───dolor.txt (empty)
│	│	├───gopher.png (70372b)
│	│	└───ipsum
│	│		└───gopher.png (70372b)
│	├───css
│	│	└───body.css (28b)
│	├───empty.txt (empty)
│	├───html
│	│	└───index.html (57b)
│	├───js
│	│	└───site.js (10b)
│	└───z_lorem
│		├───dolor.txt (empty)
│		├───gopher.png (70372b)
│		└───ipsum
│			└───gopher.png (70372b)
├───zline
│	├───empty.txt (empty)
│	└───lorem
│		├───dolor.txt (empty)
│		├───gopher.png (70372b)
│		└───ipsum
│			└───gopher.png (70372b)
└───zzfile.txt (empty)
`

func TestTreeShouldBeFull(t *testing.T) {
	out := new(bytes.Buffer)

	dirTree(out, []string{"./tree", "-f", "../../fixtures"})

	result := out.String()
	if result != fullTreeResult {
		t.Errorf("Test Failed - results not match\nGot:\n%v\nExpected:\n%v", result, fullTreeResult)
	}
}

const dirTreeResult = `├───project
├───static
│	├───a_lorem
│	│	└───ipsum
│	├───css
│	├───html
│	├───js
│	└───z_lorem
│		└───ipsum
└───zline
	└───lorem
		└───ipsum
`

func TestTreeDir(t *testing.T) {

	out := new(bytes.Buffer)

	dirTree(out, []string{"./tree", "../../fixtures"})

	result := out.String()
	if result != dirTreeResult {
		t.Errorf("Test Failed - results not match\nGot:\n%v\nExpected:\n%v", result, dirTreeResult)
	}
}
