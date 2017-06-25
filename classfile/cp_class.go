package classfile

import "fmt"

type ConstantClassInfo struct {
	nameIndex uint16
}

func (this *ConstantClassInfo) ReadInfo(reader *ClassReader) {
	this.nameIndex = reader.ReadUint16()
	//fmt.Printf("Class\t\t#%d\n", this.nameIndex)
}

func (this ConstantClassInfo) String(constantPool ConstantPool) string {
	return fmt.Sprint(constantPool[this.nameIndex])
}
