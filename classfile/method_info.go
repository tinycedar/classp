package classfile

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

func (m *MethodInfo) ReadInfo(reader *ClassReader) {
	m.accessFlags = reader.ReadUint16()
	m.nameIndex = reader.ReadUint16()
	m.descriptorIndex = reader.ReadUint16()
	m.attributes = readAttributes(reader, m.classFile.constantPool)
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
}
