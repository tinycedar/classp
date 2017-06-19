package classfile

type ConstantStringInfo struct {
	stringIndex uint16
}

func (this *ConstantStringInfo) ReadInfo(reader *ClassReader) {
	this.stringIndex = reader.ReadUint16()
	//fmt.Printf("String\t\t#%d\n", this.stringIndex)
}

func (this *ConstantStringInfo) String(constantPool []ConstantPoolInfo) string {
	if cp, ok := constantPool[this.stringIndex].(*ConstantUtf8Info); ok {
		return cp.String()
	}
	return ""
}
