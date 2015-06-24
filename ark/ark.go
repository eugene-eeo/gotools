package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func compileAll(patterns []string) []*regexp.Regexp {
	pats := []*regexp.Regexp{}
	for _, pattern := range patterns {
		r, err := regexp.Compile(pattern)
		if err != nil {
			panic(err)
		}
		pats = append(pats, r)
	}
	return pats
}

func listDir(path string) []os.FileInfo {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}
	return files
}

func allMatch(pats []*regexp.Regexp, str string) bool {
	for _, pat := range pats {
		if pat.FindString(str) != str {
			return false
		}
	}
	return true
}

func main() {
	pats := compileAll(os.Args[1:])
	for _, f := range listDir("./") {
		name := f.Name()
		if allMatch(pats, name) {
			fmt.Println(name)
		}
	}
}
