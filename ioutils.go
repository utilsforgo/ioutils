package ioutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ListAllDir(dir string, callback int) {
	info, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range info {
		fmt.Println(strings.Repeat("|  ", callback) + "|--" + item.Name())
		if item.IsDir() {
			ListAllDir(dir+string(os.PathSeparator)+item.Name(), callback+1)
		}
	}
}
