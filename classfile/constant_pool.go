package classfile

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantPool []ConstantPoolInfo

type ConstantPoolInfo interface {
	ReadInfo(reader *ClassReader)
}

func (p ConstantPool) GetConstantInfo(index uint16) ConstantPoolInfo {
	return p[index]
}

func readConstantPool(reader *ClassReader) ConstantPool {
	constantPool := make([]ConstantPoolInfo, reader.ReadUint16())
	for i := 1; i < len(constantPool); i++ {
		cpInfo := newConstantPoolInfo(reader.ReadUint8())
		if cpInfo != nil {
			cpInfo.ReadInfo(reader)
			constantPool[i] = cpInfo
		}
	}
	return constantPool
}

func newConstantPoolInfo(constType uint8) ConstantPoolInfo {
	switch constType {
	case CONSTANT_Class:
		return &ConstantClassInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{}
	case CONSTANT_String:
		return &ConstantStringInfo{}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("Invalid const type: ")
	}
}
