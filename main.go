package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	var text string
	var err error
	var size int
	var prefix string
	var out string

	flag.IntVar(&size, "size", 128, "size of qr code")
	flag.StringVar(&prefix, "prefix", "qr_", "file name prefix")
	flag.StringVar(&out, "out", "./qr", "out put path")
	flag.Parse()
	
	if f, err := os.Stat(out); os.IsNotExist(err) || !f.IsDir() {
		if err := os.Mkdir(out, 0777); err != nil {
			panic(err)
		}
	}

	if len(flag.Args()) == 0 {
		text = readPipe()
	} else if flag.Arg(0) == "-" {
		text = readStdin()
	} else {
		text= readFileByArg(flag.Arg(0))
	}
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of qrray \n")
		fmt.Fprintf(os.Stderr, "  qrray [options] text_file_path \n")
		fmt.Fprintf(os.Stderr, "Options: \n")
		flag.PrintDefaults()
	}
	if(text == ""){
		flag.Usage()
		return
	}

	// create qr code
	Creates(text, prefix, size, filepath.Base(out))

	// create index.html
	f, err := os.Create(filepath.Join(out,"index.html"))
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(f)
	ToHtml(text, prefix, size,writer )
	writer.Flush()
	defer f.Close()
}

func readPipe() string {
	stats, _ := os.Stdin.Stat()
	if stats != nil && (stats.Mode()&os.ModeCharDevice) == 0 {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		return string(bytes)
	} else {
		return ""
	}
}
func readStdin() string {
	var text string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		text += s.Text() + "\n"
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return text
}

func readFileByArg(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}