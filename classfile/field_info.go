package classfile

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
	f.attributes = readAttributes(reader, f.classFile.constantPool)
}
