package classfile

import "fmt"

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
