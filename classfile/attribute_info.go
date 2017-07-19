package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	ReadInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributes := make([]AttributeInfo, reader.ReadUint16())
	for i := 0; i < len(attributes); i++ {
		attributes[i] = readAttributeInfo(reader, cp)
	}
	return attributes
}

func readAttributeInfo(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.ReadUint16()
	attrLength := reader.ReadUint32()
	if c, ok := cp.GetConstantInfo(attrNameIndex).(*ConstantUtf8Info); ok {
		var attrInfo AttributeInfo
		switch attrName := c.String(); attrName {
		case "Code":
			attrInfo = &CodeAttribute{cp: cp}
		case "LineNumberTable":
			attrInfo = &LineNumberTableAttribute{}
		default:
			//TODO not implemented yet, just discard the bytes read
			reader.ReadBytes(int(attrLength))
		}
		if attrInfo != nil {
			attrInfo.ReadInfo(reader)
			return attrInfo
		}
	}
	return nil
}
