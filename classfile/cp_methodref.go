package classfile

import "fmt"

type ConstantMethodrefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantMethodrefInfo) ReadInfo(reader *ClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()
	//fmt.Printf("Methodref\t#%d.#%d\n", this.classIndex, this.nameAndTypeIndex)
}

func (this *ConstantMethodrefInfo) String(constantPool ConstantPool) string {
	class, _ := constantPool[this.classIndex].(*ConstantClassInfo)
	nameAndType, _ := constantPool[this.nameAndTypeIndex].(*ConstantNameAndTypeInfo)
	return fmt.Sprintf("%s.%s", class.String(constantPool), nameAndType.String(constantPool))
}
