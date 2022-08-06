package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestQrray(t *testing.T) {
	dir, _ :=  os.MkdirTemp("", "aaa")
	prefix:="qrray_test_"

	input := `hoge
huga
piyo`
	Creates(input, prefix, 256, dir)

	for i := 0; i < 3; i++ {
		path := filepath.Join(dir, prefix+fmt.Sprint(i)+".png")
		_, err := os.Stat(path)
		if err != nil {
			t.Errorf("Not Found \n%s\n", path)
		}
	}
}
