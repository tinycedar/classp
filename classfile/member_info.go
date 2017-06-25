package classfile

/*
field/method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []MemberInfo {
	members := make([]MemberInfo, reader.ReadUint16())
	for i := 0; i < len(members); i++ {
		members[i] = MemberInfo{
			cp,
			reader.ReadUint16(),
			reader.ReadUint16(),
			reader.ReadUint16(),
			readAttributes(reader, cp)}
	}
	return members
}

func (m *MemberInfo) Name() string {
	name := m.cp[m.nameIndex]
	if name, ok := name.(*ConstantUtf8Info); ok {
		return name.String()
	}
	return ""
}

func (m *MemberInfo) Descriptor() string {
	desc := m.cp[m.descriptorIndex]
	if desc, ok := desc.(*ConstantUtf8Info); ok {
		return desc.String()
	}
	return ""
}

func (m *MemberInfo) ConstantPool() ConstantPool {
	return m.cp
}

func (m *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attr := range m.attributes {
		if code, ok := attr.(*CodeAttribute); ok {
			return code
		}
	}
	return nil
}
