package main

import (
	"github.com/tinycedar/classp/classfile"
	"io/ioutil"
	"log"
)

func main() {
	bytes, err := ioutil.ReadFile("test/Sample.class")
	if err != nil {
		log.Fatal("Error reading class file")
	}

	cf := classfile.Parse(bytes)
	cf.Print()
}
