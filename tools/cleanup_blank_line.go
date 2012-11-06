package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".markdown") {
			fmt.Println("Skip " + path)
			return nil
		}
		one(path)
		return nil
	})
}

func one(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	strs := ""
	br := bufio.NewReader(bytes.NewReader(data))
	blank_prev := false
	for {
		d, _, err := br.ReadLine()
		if err != nil {
			break
		}
		str := string(d)
		if len(d) == 0 || len(strings.Trim(str, " \t")) == 0 {
			if blank_prev {
				continue
			}
			blank_prev = true
		} else {
			blank_prev = false
		}
		strs += strings.TrimRight(str, "") + "\n"
	}
	ioutil.WriteFile(path, []byte(strs), os.ModePerm)
}
