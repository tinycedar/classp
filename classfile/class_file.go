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

func (cf *ClassFile) Read(reader *ClassReader) {
	cf.size = reader.Length()
	cf.readMagic(reader)
	cf.readMinorVersion(reader)
	cf.readMajorVersion(reader)
	cf.readConstantPool(reader)
	cf.readAccessFlags(reader)
	cf.readThisClass(reader)
	cf.readSuperClass(reader)
	cf.readInterfaces(reader)
	cf.readFieldInfo(reader)
	cf.readMethodInfo(reader)
	cf.readAttributes(reader)
}

func (cf *ClassFile) Print(){
	fmt.Printf("Size: %d bytes\n", cf.size)
	fmt.Printf("magic: %x\n", cf.magic)
	fmt.Printf("minor version: %d\n", cf.minorVersion)
	fmt.Printf("major version: %d\n", cf.majorVersion)

	fmt.Printf("accessFlags: %d\n", cf.accessFlags)
	fmt.Printf("thisClass: #%d\n", cf.thisClass)
	fmt.Printf("superClass: #%d\n", cf.superClass)
}

func (cf *ClassFile) readMagic(reader *ClassReader) {
	cf.magic = reader.ReadUint32()
}

func (cf *ClassFile) readMinorVersion(reader *ClassReader) {
	cf.minorVersion = reader.ReadUint16()
}

func (cf *ClassFile) readMajorVersion(reader *ClassReader) {
	cf.majorVersion = reader.ReadUint16()
}

func (cf *ClassFile) readConstantPool(reader *ClassReader) {
	cpInfoCount := reader.ReadUint16()
	fmt.Printf("cp_info count: %d\n", cpInfoCount)
	cf.constantPool = make([]ConstantPoolInfo, cpInfoCount)
	for i := uint16(1); i < cpInfoCount; i++ {
		fmt.Printf(" #%2d = ", i)
		switch reader.ReadUint8() {
		case CONSTANT_Class:
			cpInfo := &ConstantClassInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_Fieldref:
			cpInfo := &ConstantFieldrefInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_Methodref:
			cpInfo := &ConstantMethodrefInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_InterfaceMethodref:
			cpInfo := &ConstantInterfaceMethodrefInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_String:
			cpInfo := &ConstantStringInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_Integer:
			cpInfo := &ConstantIntegerInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_Float:
			cpInfo := &ConstantFloatInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_Long:
			cpInfo := &ConstantLongInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_Double:
			cpInfo := &ConstantDoubleInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_NameAndType:
			cpInfo := &ConstantNameAndTypeInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_Utf8:
			cpInfo := &ConstantUtf8Info{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_MethodHandle:
			cpInfo := &ConstantMethodHandleInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_MethodType:
			cpInfo := &ConstantMethodTypeInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		case CONSTANT_InvokeDynamic:
			cpInfo := &ConstantInvokeDynamicInfo{}
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
		default:
			fmt.Println()
		}
	}
}

func (cf *ClassFile) readAccessFlags(reader *ClassReader) {
	cf.accessFlags = reader.ReadUint16()
}

func (cf *ClassFile) readThisClass(reader *ClassReader) {
	cf.thisClass = reader.ReadUint16()
}

func (cf *ClassFile) readSuperClass(reader *ClassReader) {
	cf.superClass = reader.ReadUint16()
}

func (cf *ClassFile) readInterfaces(reader *ClassReader) {
	var interfacesCount = reader.ReadUint16()
	cf.interfaces = make([]uint16, interfacesCount)
	for i := uint16(0); i < interfacesCount; i++ {
		cf.interfaces[i] = reader.ReadUint16()

		fmt.Printf("interface: #%d\n", cf.interfaces[i])
	}
}

func (cf *ClassFile) readFieldInfo(reader *ClassReader) {
	var fieldsCount = reader.ReadUint16()
	cf.fields = make([]fieldInfo, fieldsCount)
	for i := uint16(0); i < fieldsCount; i++ {
		fieldInfo := fieldInfo{}
		fieldInfo.ReadInfo(reader)
		cf.fields[i] = fieldInfo
	}
}

func (cf *ClassFile) readMethodInfo(reader *ClassReader) {
	methodsCount := reader.ReadUint16()
	cf.methods = make([]methodInfo, methodsCount)
	for i := uint16(0); i < methodsCount; i++ {
		methodInfo := methodInfo{classFile: cf}
		methodInfo.ReadInfo(reader)
		cf.methods[i] = methodInfo
	}
}

func (cf *ClassFile) readAttributes(reader *ClassReader) {
	attributesCount := reader.ReadUint16()
	cf.attributes = make([]AttributeInfo, attributesCount)
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

func (f *fieldInfo) ReadInfo(reader *ClassReader) {
	f.accessFlags = reader.ReadUint16()
	f.nameIndex = reader.ReadUint16()
	f.descriptorIndex = reader.ReadUint16()
	fmt.Printf("================= FieldInfo start ==========================\t\t#%d\n", f.nameIndex)
	var attributesCount = reader.ReadUint16()
	f.attributes = make([]AttributeInfo, attributesCount)
	for i := uint16(0); i < attributesCount; i++ {
		readAttributeInfo(reader)
	}
	fmt.Printf("================= FieldInfo end ==========================\t\t#%d\n", f.nameIndex)
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

func (m *methodInfo) ReadInfo(reader *ClassReader) {
	m.accessFlags = reader.ReadUint16()
	m.nameIndex = reader.ReadUint16()
	m.descriptorIndex = reader.ReadUint16()
	var attributesCount = reader.ReadUint16()
	m.attributes = make([]AttributeInfo, attributesCount)
	fmt.Printf("================= MethodInfo start ==========================\t\t#%d\n", m.nameIndex)
	for i := uint16(0); i < attributesCount; i++ {
		readAttributeInfo(reader)
	}
	fmt.Printf("================= MethodInfo end   ==========================\t\t#%d\n", m.nameIndex)
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
