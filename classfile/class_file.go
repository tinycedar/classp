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
	count := reader.ReadUint16()
	cf.constantPool = make([]ConstantPoolInfo, count)
	for i := uint16(1); i < count; i++ {
		var cpInfo ConstantPoolInfo
		switch constType := reader.ReadUint8(); constType {
		case CONSTANT_Class:
			cpInfo = &ConstantClassInfo{}
		case CONSTANT_Fieldref:
			cpInfo = &ConstantFieldrefInfo{}
		case CONSTANT_Methodref:
			cpInfo = &ConstantMethodrefInfo{}
		case CONSTANT_InterfaceMethodref:
			cpInfo = &ConstantInterfaceMethodrefInfo{}
		case CONSTANT_String:
			cpInfo = &ConstantStringInfo{}
		case CONSTANT_Integer:
			cpInfo = &ConstantIntegerInfo{}
		case CONSTANT_Float:
			cpInfo = &ConstantFloatInfo{}
		case CONSTANT_Long:
			cpInfo = &ConstantLongInfo{}
		case CONSTANT_Double:
			cpInfo = &ConstantDoubleInfo{}
		case CONSTANT_NameAndType:
			cpInfo = &ConstantNameAndTypeInfo{}
		case CONSTANT_Utf8:
			cpInfo = &ConstantUtf8Info{}
		case CONSTANT_MethodHandle:
			cpInfo = &ConstantMethodHandleInfo{}
		case CONSTANT_MethodType:
			cpInfo = &ConstantMethodTypeInfo{}
		case CONSTANT_InvokeDynamic:
			cpInfo = &ConstantInvokeDynamicInfo{}
		default:
			break
		}
		if cpInfo != nil {
			cpInfo.ReadInfo(reader)
			cf.constantPool[i] = cpInfo
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
	cf.attributes = readAttributes(reader, cf.constantPool)
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
