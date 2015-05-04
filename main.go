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
	fmt.Printf("\tClass file size: %d bytes\n", len(bytes))
	fmt.Printf("\tMagic: %x\n", reader.ReadUint32())
	fmt.Printf("\tMinor version: %d\n", reader.ReadUint16())
	fmt.Printf("\tMajor version: %d\n", reader.ReadUint16())
	constantPoolCount := int(reader.ReadUint16())
	fmt.Printf("Constant pool:\n")
	for i := 1; i < constantPoolCount; i++ {
		tag := reader.ReadUint8()
		fmt.Printf("\t#%2d = %2d ", i, tag)
		switch tag {
		case classfile.CONSTANT_Class:
			fmt.Printf("Class\t\t#%d\n", reader.ReadUint16())
		case classfile.CONSTANT_Fieldref:
			fmt.Printf("Fieldref\t#%d", reader.ReadUint16())
			fmt.Printf("#%d\n", reader.ReadUint16())
		case classfile.CONSTANT_Methodref:
			fmt.Printf("Methodref\t#%d", reader.ReadUint16())
			fmt.Printf("#%d\n", reader.ReadUint16())
		case classfile.CONSTANT_InterfaceMethodref:
			fmt.Printf("InterfaceMethodref\t\t#%d\t", reader.ReadUint16())
			fmt.Printf("#%dn", reader.ReadUint16())
		case classfile.CONSTANT_String:
			fmt.Printf("String\t\t#%d\n", reader.ReadUint16())
		case classfile.CONSTANT_Integer:
			fmt.Printf("Integer\t\tbytes: %d\n", reader.ReadUint32())
		case classfile.CONSTANT_Float:
			fmt.Printf("Float\t\tbytes: %d\n", reader.ReadUint32())
		case classfile.CONSTANT_Long:
			fmt.Printf("Long\t\thigh_bytes: %d\n", reader.ReadUint32())
			fmt.Printf("low_bytes: %d\n", reader.ReadUint32())
		case classfile.CONSTANT_Double:
			fmt.Printf("Double\t\thigh_bytes: %d\n", reader.ReadUint32())
			fmt.Printf("low_bytes: %d\n", reader.ReadUint32())
		case classfile.CONSTANT_NameAndType:
			fmt.Printf("NameAndType\t#%d", reader.ReadUint16())
			fmt.Printf("#%d\n", reader.ReadUint16())
		case classfile.CONSTANT_Utf8:
			fmt.Printf("Utf8\t\t%s\n", reader.ReadString())
		case classfile.CONSTANT_MethodHandle:
			fmt.Printf("MethodHandle\t\treference_kind: %d\n", reader.ReadUint8())
			fmt.Printf("#%d\n", reader.ReadUint16())
		case classfile.CONSTANT_MethodType:
			fmt.Printf("MethodType\t\t#%d\n", reader.ReadUint16())
		case classfile.CONSTANT_InvokeDynamic:
			fmt.Printf("InvokeDynamic\t\t#%d\n", reader.ReadUint16())
			fmt.Printf("#%d\n", reader.ReadUint16())
		default:
			fmt.Println("error")
		}
	}
	fmt.Printf("accessFlags: %x\n", reader.ReadUint16())
	fmt.Printf("this class: #%d\n", reader.ReadUint16())
	fmt.Printf("super class: #%d\n", reader.ReadUint16())

	interfaceCount := int(reader.ReadUint16())
	for i := 0; i < interfaceCount; i++ {
		fmt.Printf("interface: #%d\n", reader.ReadUint16())
	}

	fields_count := int(reader.ReadUint16())
	for i := 0; i < fields_count; i++ {
		fmt.Printf("field accessFlags: %d\t\t", reader.ReadUint16())
		fmt.Printf("field #%d:#%d\t", reader.ReadUint16(), reader.ReadUint16())
		fmt.Printf("field attribute_count: %d\n", reader.ReadUint16())
	}

	method_count := int(reader.ReadUint16())
	for i := 0; i < method_count; i++ {
		fmt.Printf("method accessFlags: %d\t\t", reader.ReadUint16())
		fmt.Printf("method #%d.#%d\t", reader.ReadUint16(), reader.ReadUint16())
		attrbutes_count := int(reader.ReadUint16())
		for j := 0; j< attrbutes_count;j++{
			fmt.Printf("method: %d\t\t", reader.ReadUint16())
			fmt.Printf("method attr_info: %x\n", string(reader.ReadBytes(reader.ReadUint32())))
		}
	}

	// todo attribute_info
}
