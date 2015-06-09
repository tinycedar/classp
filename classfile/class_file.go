package classfile

import (
	"fmt"
)

/*
ClassFile {
	u4				magic;
	u2 				minor_version;
	u2 				major_version;
	u2 				constant_pool_count;
	cp_info 		constant_pool[constant_pool_count-1];
	u2 				access_flags;
	u2 				this_class;
	u2 				super_class;
	u2 				interfaces_count;
	u2 				interfaces[interfaces_count];
	u2 				fields_count;
	field_info 		fields[fields_count];
	u2 				methods_count;
	method_info 	methods[methods_count];
	u2 				attributes_count;
	attribute_info 	attributes[attributes_count];
}
*/
type ClassFile struct {
	size 	     int
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool []ConstantPoolInfo
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []fieldInfo
	methods      []methodInfo
	attributes   []AttributeInfo
}

func NewClassFile() *ClassFile {
	return &ClassFile{}
}

func (this *ClassFile) Read(reader *ClassReader) {
	this.size = reader.Length()
	this.readMagic(reader)
	this.readMinorVersion(reader)
	this.readMajorVersion(reader)

	fmt.Printf("Size: %d bytes\n", this.size)
	fmt.Printf("magic: %x\n", this.magic)
	fmt.Printf("minor version: %d\n", this.minorVersion)
	fmt.Printf("major version: %d\n", this.majorVersion)

	this.readConstantPool(reader)
	this.readAccessFlags(reader)
	this.readThisClass(reader)
	this.readSuperClass(reader)
	fmt.Printf("accessFlags: %d\n", this.accessFlags)
	fmt.Printf("thisClass: #%d\n", this.thisClass)
	fmt.Printf("superClass: #%d\n", this.superClass)
	this.readInterfaces(reader)
	this.readFieldInfo(reader)
	this.readMethodInfo(reader)
	this.readAttributes(reader)
}

func (this *ClassFile) readMagic(reader *ClassReader) {
	this.magic = reader.ReadUint32()
}

func (this *ClassFile) readMinorVersion(reader *ClassReader) {
	this.minorVersion = reader.ReadUint16()
}

func (this *ClassFile) readMajorVersion(reader *ClassReader) {
	this.majorVersion = reader.ReadUint16()
}

func (this *ClassFile) readConstantPool(reader *ClassReader) {
	cpInfoCount := reader.ReadUint16()
	fmt.Printf("cp_info count: %d\n", cpInfoCount)
	this.constantPool = make([]ConstantPoolInfo, cpInfoCount)
	for i := uint16(1); i < cpInfoCount; i++ {
		fmt.Printf(" #%2d = ", i)
		switch reader.ReadUint8() {
		case CONSTANT_Class:
			cpInfo := &ConstantClassInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_Fieldref:
			cpInfo := &ConstantFieldrefInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_Methodref:
			cpInfo := &ConstantMethodrefInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_InterfaceMethodref:
			cpInfo := &ConstantInterfaceMethodrefInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_String:
			cpInfo := &ConstantStringInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_Integer:
			cpInfo := &ConstantIntegerInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_Float:
			cpInfo := &ConstantFloatInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_Long:
			cpInfo := &ConstantLongInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_Double:
			cpInfo := &ConstantDoubleInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_NameAndType:
			cpInfo := &ConstantNameAndTypeInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_Utf8:
			cpInfo := &ConstantUtf8Info{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_MethodHandle:
			cpInfo := &ConstantMethodHandleInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_MethodType:
			cpInfo := &ConstantMethodTypeInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		case CONSTANT_InvokeDynamic:
			cpInfo := &ConstantInvokeDynamicInfo{}
			cpInfo.ReadInfo(reader)
			this.constantPool[i] = cpInfo
		default:
			fmt.Println()
		}
	}
}

func (this *ClassFile) readAccessFlags(reader *ClassReader) {
	this.accessFlags = reader.ReadUint16()
}

func (this *ClassFile) readThisClass(reader *ClassReader) {
	this.thisClass = reader.ReadUint16()
}

func (this *ClassFile) readSuperClass(reader *ClassReader) {
	this.superClass = reader.ReadUint16()
}

func (this *ClassFile) readInterfaces(reader *ClassReader) {
	var interfacesCount = reader.ReadUint16()
	this.interfaces = make([]uint16, interfacesCount)
	for i := uint16(0); i < interfacesCount; i++ {
		this.interfaces[i] = reader.ReadUint16()

		fmt.Printf("interface: #%d\n", this.interfaces[i])
	}
}

func (this *ClassFile) readFieldInfo(reader *ClassReader) {
	var fieldsCount = reader.ReadUint16()
	this.fields = make([]fieldInfo, fieldsCount)
	for i := uint16(0); i < fieldsCount; i++ {
		fieldInfo := fieldInfo{}
		fieldInfo.ReadInfo(reader)
		this.fields[i] = fieldInfo
	}
}

func (this *ClassFile) readMethodInfo(reader *ClassReader) {
	methodsCount := reader.ReadUint16()
	this.methods = make([]methodInfo, methodsCount)
	for i := uint16(0); i < methodsCount; i++ {
		methodInfo := methodInfo{classFile: this}
		methodInfo.ReadInfo(reader)
		this.methods[i] = methodInfo
	}
}

func (this *ClassFile) readAttributes(reader *ClassReader) {
	attributesCount := reader.ReadUint16()
	this.attributes = make([]AttributeInfo, attributesCount)
	for i := uint16(0); i < attributesCount; i++ {
		readAttributeInfo(reader)
	}
}

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type fieldInfo struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func (this *fieldInfo) ReadInfo(reader *ClassReader) {
	this.accessFlags = reader.ReadUint16()
	this.nameIndex = reader.ReadUint16()
	this.descriptorIndex = reader.ReadUint16()
	fmt.Printf("================= FieldInfo start ==========================\t\t#%d\n", this.nameIndex)
	var attributesCount = reader.ReadUint16()
	this.attributes = make([]AttributeInfo, attributesCount)
	for i := uint16(0); i < attributesCount; i++ {
		readAttributeInfo(reader)
	}
	fmt.Printf("================= FieldInfo end ==========================\t\t#%d\n", this.nameIndex)
}

/*
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type methodInfo struct {
	classFile       *ClassFile
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func (this *methodInfo) ReadInfo(reader *ClassReader) {
	this.accessFlags = reader.ReadUint16()
	this.nameIndex = reader.ReadUint16()
	this.descriptorIndex = reader.ReadUint16()
	var attributesCount = reader.ReadUint16()
	this.attributes = make([]AttributeInfo, attributesCount)
	fmt.Printf("================= MethodInfo start ==========================\t\t#%d\n", this.nameIndex)
	for i := uint16(0); i < attributesCount; i++ {
		readAttributeInfo(reader)
	}
	fmt.Printf("================= MethodInfo end   ==========================\t\t#%d\n", this.nameIndex)
}

func readAttributeInfo(reader *ClassReader) {
	attributeNameIndex := reader.ReadUint16()
	attributeLength := reader.ReadUint32()
	fmt.Printf("Code attributeLength\t\t%d\n", attributeLength)
	//TODO attributeNameIndex
	if attributeNameIndex == uint16(22) {
		code := CodeAttribute{}
		code.ReadInfo(reader)
		fmt.Printf("Code ending....\t\t#%d\n", attributeNameIndex)
	} else if attributeNameIndex == uint16(23) {
		attribute := LineNumberTableAttribute{}
		attribute.ReadInfo(reader)
		fmt.Printf("LineNumberTable ending....\t\t#%d\n", attributeNameIndex)
	} else if attributeNameIndex == uint16(31) {
		//TODO add SourceFile attribute
		fmt.Printf("SourceFile attribute\t\t#%d\n", attributeNameIndex)
	} else {
		fmt.Printf("Not a Code attribute\t\t#%d\n", attributeNameIndex)
	}
}
