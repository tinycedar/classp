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
	size         int
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool []ConstantPoolInfo
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []FieldInfo
	methods      []MethodInfo
	attributes   []AttributeInfo
}

func NewClassFile() *ClassFile {
	return &ClassFile{}
}

func (cf *ClassFile) Methods() []MethodInfo {
	return cf.methods
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

func (cf *ClassFile) Print() {
	fmt.Printf("Size: %d bytes\n", cf.size)
	fmt.Printf("magic: %x\n", cf.magic)
	fmt.Printf("minor version: %d\n", cf.minorVersion)
	fmt.Printf("major version: %d\n", cf.majorVersion)

	fmt.Printf("accessFlags: %d\n", cf.accessFlags)
	fmt.Printf("thisClass: #%d\n", cf.thisClass)
	fmt.Printf("superClass: #%d\n", cf.superClass)

	fmt.Println("**********************************************************")
	for i, length := 1, len(cf.constantPool); i < length; i++ {
		fmt.Printf(" #%2d = ", i)
		//fmt.Println(cf.constantPool[i])
		if cp, ok := cf.constantPool[i].(*ConstantClassInfo); ok {
			fmt.Printf("Class\t\t\t\t#%d\t\t\t// %s", cp.nameIndex, cp.String(cf.constantPool))
		} else if cp, ok := cf.constantPool[i].(*ConstantFieldrefInfo); ok {
			fmt.Printf("Fieldref\t\t#%d.#%d\t\t// %s", cp.classIndex, cp.nameAndTypeIndex, cp.String(cf.constantPool))
			//cf.constantPool[cp.classIndex]
		} else if cp, ok := cf.constantPool[i].(*ConstantUtf8Info); ok {
			fmt.Printf("Utf8\t\t\t\t\t%s", cp)
		} else if cp, ok := cf.constantPool[i].(*ConstantNameAndTypeInfo); ok {
			fmt.Printf("NameAndType\t\t\t#%d:#%d\t\t\t// %s", cp.nameIndex, cp.descriptorIndex, cp.String(cf.constantPool))
		}
		fmt.Println()
	}
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
	//fmt.Printf("cp_info count: %d\n", cpInfoCount)
	cf.constantPool = make([]ConstantPoolInfo, cpInfoCount)
	for i := uint16(1); i < cpInfoCount; i++ {
		//fmt.Printf(" #%2d = ", i)
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
	cf.fields = make([]FieldInfo, fieldsCount)
	for i := uint16(0); i < fieldsCount; i++ {
		fieldInfo := FieldInfo{classFile: cf}
		fieldInfo.ReadInfo(reader)
		cf.fields[i] = fieldInfo
	}
}

func (cf *ClassFile) readMethodInfo(reader *ClassReader) {
	methodsCount := reader.ReadUint16()
	cf.methods = make([]MethodInfo, methodsCount)
	for i := uint16(0); i < methodsCount; i++ {
		methodInfo := MethodInfo{classFile: cf}
		methodInfo.ReadInfo(reader)
		cf.methods[i] = methodInfo
	}
}

func (cf *ClassFile) readAttributes(reader *ClassReader) {
	attributesCount := reader.ReadUint16()
	cf.attributes = make([]AttributeInfo, attributesCount)
	for i := uint16(0); i < attributesCount; i++ {
		cf.attributes[i] = readAttributeInfo(reader, cf)
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
type FieldInfo struct {
	classFile       *ClassFile
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func (f *FieldInfo) ReadInfo(reader *ClassReader) {
	f.accessFlags = reader.ReadUint16()
	f.nameIndex = reader.ReadUint16()
	f.descriptorIndex = reader.ReadUint16()
	fmt.Printf("================= FieldInfo start ==========================\t\t#%d\n", f.nameIndex)
	var attributesCount = reader.ReadUint16()
	f.attributes = make([]AttributeInfo, attributesCount)
	for i := uint16(0); i < attributesCount; i++ {
		f.attributes[i] = readAttributeInfo(reader, f.classFile)
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
type MethodInfo struct {
	classFile       *ClassFile
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func (m *MethodInfo) Name() string {
	name := m.classFile.constantPool[m.nameIndex]
	if name, ok := name.(*ConstantUtf8Info); ok {
		return name.String()
	}
	return ""
}

func (m *MethodInfo) Descriptor() string {
	desc := m.classFile.constantPool[m.descriptorIndex]
	if desc, ok := desc.(*ConstantUtf8Info); ok {
		return desc.String()
	}
	return ""
}

func (m *MethodInfo) ConstantPool() []ConstantPoolInfo {
	return m.classFile.constantPool
}

func (m *MethodInfo) CodeAttribute() *CodeAttribute {
	for _, attr := range m.attributes {
		if code, ok := attr.(*CodeAttribute); ok {
			return code
		}
	}
	return nil
	//return CodeAttribute{}
}

func (m *MethodInfo) ReadInfo(reader *ClassReader) {
	m.accessFlags = reader.ReadUint16()
	m.nameIndex = reader.ReadUint16()
	m.descriptorIndex = reader.ReadUint16()
	var attributesCount = reader.ReadUint16()
	m.attributes = make([]AttributeInfo, attributesCount)
	//fmt.Printf("================= MethodInfo start ==========================\t\t#%d\n", m.nameIndex)
	for i := uint16(0); i < attributesCount; i++ {
		m.attributes[i] = readAttributeInfo(reader, m.classFile)
	}
	//fmt.Printf("================= MethodInfo end   ==========================\t\t#%d\n", m.nameIndex)
}

func readAttributeInfo(reader *ClassReader, cf *ClassFile) AttributeInfo {
	attributeNameIndex := reader.ReadUint16()
	attributeLength := reader.ReadUint32()
	//fmt.Printf("Code attributeLength\t\t%d\n", attributeLength)
	if cp, ok := cf.constantPool[attributeNameIndex].(*ConstantUtf8Info); ok {
		switch cp.String() {
		case "Code":
			code := &CodeAttribute{classFile: cf}
			code.ReadInfo(reader)
			return code
		case "LineNumberTable":
			attribute := &LineNumberTableAttribute{}
			attribute.ReadInfo(reader)
			return attribute
		case "SourceFile":
			//TODO add SourceFile attribute
			//fmt.Printf("SourceFile attribute\t\t#%d\n", attributeNameIndex)
		default:
			fmt.Printf("invalid attribute index: %d, length: %d\n", attributeNameIndex, attributeLength)
		}
	}
	return nil
}
