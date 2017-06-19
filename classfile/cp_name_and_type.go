package classfile

import "fmt"

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (this *ConstantNameAndTypeInfo) ReadInfo(reader *ClassReader) {
	this.nameIndex = reader.ReadUint16()
	this.descriptorIndex = reader.ReadUint16()
	//fmt.Printf("NameAndType\t#%d:#%d\n", this.nameIndex, this.descriptorIndex)
}

func (this ConstantNameAndTypeInfo) String(constantPool []ConstantPoolInfo) string {
	return fmt.Sprintf("%s:%s", constantPool[this.nameIndex], constantPool[this.descriptorIndex])
}
