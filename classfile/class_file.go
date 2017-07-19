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
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []MemberInfo
	methods      []MemberInfo
	attributes   []AttributeInfo
}

func Parse(bytes []byte) *ClassFile {
	reader := NewClassReader(bytes)
	cf := &ClassFile{}
	cf.size = reader.Length()
	cf.magic = reader.ReadUint32()
	cf.minorVersion = reader.ReadUint16()
	cf.majorVersion = reader.ReadUint16()
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.ReadUint16()
	cf.thisClass = reader.ReadUint16()
	cf.superClass = reader.ReadUint16()
	cf.readInterfaces(reader)
	//TODO from now on, we can speed up by run concurrently
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)
	return cf
}

func (cf *ClassFile) readInterfaces(reader *ClassReader) {
	cf.interfaces = make([]uint16, reader.ReadUint16())
	for i := 0; i < len(cf.interfaces); i++ {
		cf.interfaces[i] = reader.ReadUint16()
		fmt.Printf("interface: #%d\n", cf.interfaces[i])
	}
}

func (cf *ClassFile) Methods() []MemberInfo {
	return cf.methods
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
	fmt.Println("Add support for travis")
}
