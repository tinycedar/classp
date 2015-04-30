package main

import (
	"fmt"
	"io/ioutil"

	"github.com/tinycedar/classParser/classfile"
)

func main() {
	bytes, err := ioutil.ReadFile("test/Sample.class")
	if err != nil {
		panic(err)
	}
	reader := classfile.NewClassReader(bytes)
	fmt.Printf("Class file size: %d bytes\n", len(bytes))
	fmt.Printf("Magic: %x\n", reader.ReadUint32())
	fmt.Printf("Minor version: %d\n", reader.ReadUint16())
	fmt.Printf("Major version: %d\n", reader.ReadUint16())
	constantPoolCount := int(reader.ReadUint16())
	fmt.Printf("Constant pool count: %d\n", constantPoolCount)
	for i := 1; i < constantPoolCount; i++ {
		tag := reader.ReadUint8()
		fmt.Printf("tag: %d\t", tag)
		switch tag {
		case classfile.CONSTANT_Class:
			fmt.Printf("name_index: %d\n", reader.ReadUint16())
		case classfile.CONSTANT_Fieldref:
			fmt.Printf("class_index: %d\t", reader.ReadUint16())
			fmt.Printf("name_and_type_index: %d\n", reader.ReadUint16())
		case classfile.CONSTANT_Methodref:
			fmt.Printf("class_index: %d\t", reader.ReadUint16())
			fmt.Printf("name_and_type_index: %d\n", reader.ReadUint16())
		case classfile.CONSTANT_InterfaceMethodref:
			fmt.Printf("class_index: %d\t", reader.ReadUint16())
			fmt.Printf("name_and_type_index: %d\n", reader.ReadUint16())
		case classfile.CONSTANT_String:
			fmt.Printf("string_index: %d\n", reader.ReadUint16())

		case classfile.CONSTANT_Integer:
			fmt.Printf("bytes: %d\n", reader.ReadUint32())
		case classfile.CONSTANT_Float:
			fmt.Printf("bytes: %d\n", reader.ReadUint32())
		case classfile.CONSTANT_Long:
			fmt.Printf("high_bytes: %d\n", reader.ReadUint32())
			fmt.Printf("low_bytes: %d\n", reader.ReadUint32())
		case classfile.CONSTANT_Double:
			fmt.Printf("high_bytes: %d\n", reader.ReadUint32())
			fmt.Printf("low_bytes: %d\n", reader.ReadUint32())
		case classfile.CONSTANT_NameAndType:
			fmt.Printf("name_index: %d\n", reader.ReadUint16())
			fmt.Printf("descriptor_index: %d\n", reader.ReadUint16())
		case classfile.CONSTANT_Utf8:
			fmt.Printf("length: %d\t", reader.ReadUint16())
			fmt.Printf("bytes: %s\n", reader.ReadUint8())

		case classfile.CONSTANT_MethodHandle:
			fmt.Printf("reference_kind: %d\n", reader.ReadUint8())
			fmt.Printf("reference_index: %d\n", reader.ReadUint16())
		case classfile.CONSTANT_MethodType:
			fmt.Printf("descriptor_index: %d\n", reader.ReadUint16())
		case classfile.CONSTANT_InvokeDynamic:
			fmt.Printf("bootstrap_method_attr_index: %d\n", reader.ReadUint16())
			fmt.Printf("name_and_type_index: %d\n", reader.ReadUint16())
		default:
			fmt.Println("error")
		}
	}
}
