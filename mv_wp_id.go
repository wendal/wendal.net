package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(os.Args)
		usage()
		os.Exit(1)
	}
	_, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = os.MkdirAll(os.Args[2], os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
	fmt.Println("Found " + path)
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	br := bufio.NewReaderSize(f, 1024*1024)
	wp_id := int64(-1)
	for {
		data, _, err := br.ReadLine()
		if err != nil {
			break
		}
		if len(data) == 1 {
			continue
		}
		str := string(data)
		if !strings.HasPrefix(str, "wordpress_id: ") {
			continue
		}
		wp_id, err = strconv.ParseInt(str[len("wordpress_id: "):], 10, 64)
		if err != nil {
			fmt.Println("Error --> " + str)
		}
	}
	if wp_id < 0 {
		fmt.Println("Can't found wordpress_id")
		return
	}
	data, err := ioutil.ReadFile(path)
	ioutil.WriteFile(fmt.Sprintf("%s/%d.markdown", os.Args[2], wp_id), data, os.ModePerm)
}

func usage() {
	fmt.Println("usage: mv_wp_id <src_dir> <dst_dir>")
}
