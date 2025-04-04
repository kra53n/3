package main

import (
	"os"
	"path"
)

type Figures struct {
	_0 string
	_1 string
	_2 string
	_3 string
}

var f Figures = Figures{
	_0: "─",
	_1: "│",
	_2: "└",
	_3: "├",
}
var f2rounded string = "╰"

func main() {
	var dir string
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir = "."
	}
	println(dir)
	_3(dir, "")
}

func _3(base string, prefix string) {
	entries, _ := os.ReadDir(base)
	for i, entry := range entries {
		if entry.IsDir() {
			if !isDirOk(entry) {
				continue
			}
		} else {
			if !isFileOk(entry) {
				continue
			}
		}
		subpath := path.Join(base, entry.Name())
		if i == len(entries)-1 {
			println(prefix+f2rounded+f._0+f._0, entry.Name())
			_3(subpath, prefix+"    ")
		} else {
			println(prefix+f._3+f._0+f._0, entry.Name())
			_3(subpath, prefix+f._1+"   ")
		}
	}
}

func isDirOk(entry os.DirEntry) bool {
	switch entry.Name() {
	case ".git", "__pycache__", ".idea", ".pytest_cache", ".venv":
		return false
	}
	return true
}

func isFileOk(entry os.DirEntry) bool {
	return true
}
