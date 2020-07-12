package main

import (
	"html/template"
	"path/filepath"
	"io/ioutil"
	"flag"
	"strings"
	"os"
)

type Post struct {
	Content string
}

func main() {
	//Flags
	// filePtr := flag.String("file", "first-post.txt", "string")
	dirPtr := flag.String("dir", ".", "string")

	flag.Parse()

	//Files are a slice of strings
	// dat, err := ioutil.ReadFile(*filePtr)
	// if err != nil {
	// 	panic(err)
	// }
	// postNew := Post{string(dat)}

	matches, _ := filepath.Glob(*dirPtr + "/*.txt")

	for _, match := range matches {
		dat, err := ioutil.ReadFile(match)
		if err != nil {
			panic(err)
		}
		postNew := Post{string(dat)}

		newFileName := strings.TrimSuffix(match, filepath.Ext(match)) + ".html"
		outfile, err := os.Create(newFileName)
			if err != nil {
				panic(err)
			}
	
		tmpl, err := template.ParseFiles("template.tmpl")
		if err != nil {
			panic(err)
		}
	
		err = tmpl.Execute(outfile, postNew)
		if err != nil {
			panic(err)
		}
	}
}