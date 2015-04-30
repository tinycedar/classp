package main

import (
	"io/ioutil"
	"log"

	"github.com/tinycedar/classParser/classfile"
)

func main() {
	bytes, err := ioutil.ReadFile("test/Sample.class")
	if err != nil {
		panic(err)
	}
	reader := classfile.NewClassReader(bytes)
	log.Printf("Class file size: %d bytes\n", len(bytes))
	log.Printf("Magic: %x\n", reader.ReadUint32())
	log.Printf("Minor version: %d\n", reader.ReadUint16())
	log.Printf("Major version: %d\n", reader.ReadUint16())
	log.Printf("Constant pool count: %d\n", reader.ReadUint16())
}
