package classfile

import "fmt"

type ConstantIntegerInfo struct {
	bytes uint32
}

func (this *ConstantIntegerInfo) ReadInfo(reader *ClassReader) {
	this.bytes = reader.ReadUint32()
	//fmt.Printf("Integer\t\t%s\n", this.bytes)
}

type ConstantFloatInfo struct {
	bytes uint32
}

func (this *ConstantFloatInfo) ReadInfo(reader *ClassReader) {
	this.bytes = reader.ReadUint32()
	//fmt.Printf("Float\t\t%s\n", this.bytes)
}

type ConstantLongInfo struct {
	highBytes uint32
	lowBytes  uint32
}

func (this *ConstantLongInfo) ReadInfo(reader *ClassReader) {
	this.highBytes = reader.ReadUint32()
	this.lowBytes = reader.ReadUint32()
	fmt.Printf("Long\t\t%s%s\n", this.highBytes, this.lowBytes)
}

type ConstantDoubleInfo struct {
	highBytes uint32
	lowBytes  uint32
}

func (this *ConstantDoubleInfo) ReadInfo(reader *ClassReader) {
	this.highBytes = reader.ReadUint32()
	this.lowBytes = reader.ReadUint32()
	//fmt.Printf("Double\t\t%s%s\n", this.highBytes, this.lowBytes)
}
