package classfile

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (this *ConstantMethodHandleInfo) ReadInfo(reader *ClassReader) {
	this.referenceKind = reader.ReadBytes(1)[0]
	this.referenceIndex = reader.ReadUint16()
	//fmt.Printf("MethodHandle\t\t%s%s\n", this.referenceKind, this.referenceIndex)
}

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (this *ConstantMethodTypeInfo) ReadInfo(reader *ClassReader) {
	this.descriptorIndex = reader.ReadUint16()
	//fmt.Printf("MethodType\t%s\n", this.descriptorIndex)
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (this *ConstantInvokeDynamicInfo) ReadInfo(reader *ClassReader) {
	this.bootstrapMethodAttrIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()
	//fmt.Printf("InvokeDynamic\t\t%s%s\n", this.bootstrapMethodAttrIndex, this.nameAndTypeIndex)
}
