package classfile

import (
//"fmt"
)

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

func readAttributes(reader *ClassReader, cp []ConstantPoolInfo) []AttributeInfo {
	attributes := make([]AttributeInfo, reader.ReadUint16())
	for i := 0; i < len(attributes); i++ {
		attributes[i] = readAttributeInfo(reader, cp)
	}
	return attributes
}

func readAttributeInfo(reader *ClassReader, cp []ConstantPoolInfo) AttributeInfo {
	attributeNameIndex := reader.ReadUint16()
	_ = reader.ReadUint32() // attributeLength
	//fmt.Printf("Code attributeLength\t\t%d\n", attributeLength)
	if c, ok := cp[attributeNameIndex].(*ConstantUtf8Info); ok {
		attrInfo := newAttributeInfo(c.String(), cp)
		if attrInfo != nil {
			attrInfo.ReadInfo(reader)
			return attrInfo
		}
	}
	return nil
}

func newAttributeInfo(attrName string, cp []ConstantPoolInfo) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "SourceFile":
		//TODO add SourceFile attribute
		//fmt.Printf("SourceFile attribute\t\t#%d\n", attributeNameIndex)
	default:
		//fmt.Printf("invalid attribute index: %d, length: %d\n", attributeNameIndex, attributeLength)
	}
	return nil
}
