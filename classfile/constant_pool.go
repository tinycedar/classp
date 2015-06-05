package classfile

import (
	"fmt"
)

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantPoolInfo interface {
	ReadInfo(reader *ClassReader)
}

type ConstantClassInfo struct {
	nameIndex uint16
}

func (this *ConstantClassInfo) ReadInfo(reader *ClassReader) {
	this.nameIndex = reader.ReadUint16()

	fmt.Printf("Class\t\t#%d\n", this.nameIndex)
}

type ConstantFieldrefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantFieldrefInfo) ReadInfo(reader *ClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()

	fmt.Printf("Fieldref\t\t#%d.#%d\n", this.classIndex, this.nameAndTypeIndex)
}

type ConstantMethodrefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantMethodrefInfo) ReadInfo(reader *ClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()

	fmt.Printf("Methodref\t#%d.#%d\n", this.classIndex, this.nameAndTypeIndex)
}

type ConstantInterfaceMethodrefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantInterfaceMethodrefInfo) ReadInfo(reader *ClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()

	fmt.Printf("InterfaceMethodref\t\t#%d.#%d\n", this.classIndex, this.nameAndTypeIndex)
}

type ConstantStringInfo struct {
	stringIndex uint16
}

func (this *ConstantStringInfo) ReadInfo(reader *ClassReader) {
	this.stringIndex = reader.ReadUint16()

	fmt.Printf("String\t\t#%d\n", this.stringIndex)
}

type ConstantIntegerInfo struct {
	bytes uint32
}

func (this *ConstantIntegerInfo) ReadInfo(reader *ClassReader) {
	this.bytes = reader.ReadUint32()

	fmt.Printf("Integer\t\t%s\n", this.bytes)
}

type ConstantFloatInfo struct {
	bytes uint32
}

func (this *ConstantFloatInfo) ReadInfo(reader *ClassReader) {
	this.bytes = reader.ReadUint32()

	fmt.Printf("Float\t\t%s\n", this.bytes)
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
	fmt.Printf("Double\t\t%s%s\n", this.highBytes, this.lowBytes)
}

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (this *ConstantNameAndTypeInfo) ReadInfo(reader *ClassReader) {
	this.nameIndex = reader.ReadUint16()
	this.descriptorIndex = reader.ReadUint16()
	fmt.Printf("NameAndType\t#%d:#%d\n", this.nameIndex, this.descriptorIndex)
}

type ConstantUtf8Info struct {
	bytes []byte //u2 length
}

func (this *ConstantUtf8Info) ReadInfo(reader *ClassReader) {
	this.bytes = reader.ReadBytes(int(reader.ReadUint16()))
	fmt.Printf("Utf8\t\t%s\n", this.bytes)
}

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (this *ConstantMethodHandleInfo) ReadInfo(reader *ClassReader) {
	this.referenceKind = reader.ReadBytes(1)[0]
	this.referenceIndex = reader.ReadUint16()
	fmt.Printf("MethodHandle\t\t%s%s\n", this.referenceKind, this.referenceIndex)
}

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (this *ConstantMethodTypeInfo) ReadInfo(reader *ClassReader) {
	this.descriptorIndex = reader.ReadUint16()
	fmt.Printf("MethodType\t%s\n", this.descriptorIndex)
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (this *ConstantInvokeDynamicInfo) ReadInfo(reader *ClassReader) {
	this.bootstrapMethodAttrIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()
	fmt.Printf("InvokeDynamic\t\t%s%s\n", this.bootstrapMethodAttrIndex, this.nameAndTypeIndex)
}
