package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"path/filepath"
	"strings"
	"sync"

	qrcode "github.com/skip2/go-qrcode"
)

//go:embed pages/index.html
var index string

func ToHtml(lines string,prefix string,size int,writer io.Writer) {
    tpl, err := template.New("").Parse(index)
	if err != nil {
		panic(err)
	}
    if err = tpl.Execute(writer, struct {
		Size   int
		Prefix string
		Count    int
	}{
		Size:   size,
		Prefix: prefix,
		Count:    len(strings.Split(lines, "\n")),
	}); err != nil {
        log.Fatal(err)
    }
}

func Creates(lines string,prefix string,size int,dir string) {
	arr := strings.Split(lines, "\n")
	var wg sync.WaitGroup
    for i := 0; i < len(arr); i++ {
		text := arr[i]
		path := filepath.Join(dir ,prefix + fmt.Sprint(i) + ".png")
		wg.Add(1)
		go func (){
			create(text,path,size)
			wg.Done()
		}()
    }
	wg.Wait()
}
func create(text string,path string,size int) {
	err := qrcode.WriteFile(text, qrcode.Medium, size, path)
	if err != nil {
		panic(err)
	}
}
