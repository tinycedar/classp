package classfile

type ConstantInterfaceMethodrefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantInterfaceMethodrefInfo) ReadInfo(reader *ClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()
	//fmt.Printf("InterfaceMethodref\t\t#%d.#%d\n", this.classIndex, this.nameAndTypeIndex)
}
