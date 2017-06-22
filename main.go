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

	classfile.Parse(bytes).Print()
}
