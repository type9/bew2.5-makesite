package main

import (
	"html/template"
	"io/ioutil"
	"fmt"
	"os"
)

type Post struct {
	Content string
}

func main() {
	//Files are a slice of strings
	dat, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	postNew := Post{string(dat)}

	t := template.Must(template.New("html-tmpl").Parse("template.tmpl"))
		err = t.Execute(os.Stdout, postNew)
		if err != nil {
			panic(err)
		}
}