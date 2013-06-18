package main

import "os"
import "log"

func main() {
	attr := &os.ProcAttr{Files:[]*os.File{nil,nil,nil}}
	_,err := os.StartProcess("gor.exe", []string{}, attr)
	log.Println(err)
}