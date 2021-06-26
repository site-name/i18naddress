package main

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/shurcooL/vfsgen"
)

func main() {
	dir, _ := os.Getwd()
	var fs http.FileSystem = http.Dir(path.Join(dir, "data"))

	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:    "data.go",
		PackageName: "i18naddress",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
