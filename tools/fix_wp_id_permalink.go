package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Args)
		usage()
		os.Exit(1)
	}
	_, err := os.Stat(os.Args[1])
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
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(data)
	if strings.Contains(str, "permalink:") {
		fmt.Println("Had permalink, skip")
		return
	}
	if !strings.Contains(str, "wordpress_id: ") {
		fmt.Println("Not wordpress_id found ,skip")
		return
	}

	br := bufio.NewReader(bytes.NewReader(data))

	wp_id := int64(-1)
	for {
		data, _, err := br.ReadLine()
		if err != nil {
			break
		}
		if len(data) < 1 {
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
	str = strings.Replace(str, "wordpress_id", fmt.Sprintf("permalink: '/%d.html'\nwordpress_id", wp_id), -1)
	ioutil.WriteFile(path, []byte(str), os.ModePerm)
	fmt.Println("Done --> " + path)
}

func usage() {
	fmt.Println("usage: mv_wp_id <src_dir>")
}
