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

type ConstantPoolInfo interface {
	ReadInfo(reader *ClassReader)
}

func readConstantPool(reader *ClassReader) []ConstantPoolInfo {
	constantPool := make([]ConstantPoolInfo, reader.ReadUint16())
	for i := 1; i < len(constantPool); i++ {
		var cpInfo ConstantPoolInfo
		switch constType := reader.ReadUint8(); constType {
		case CONSTANT_Class:
			cpInfo = &ConstantClassInfo{}
		case CONSTANT_Fieldref:
			cpInfo = &ConstantFieldrefInfo{}
		case CONSTANT_Methodref:
			cpInfo = &ConstantMethodrefInfo{}
		case CONSTANT_InterfaceMethodref:
			cpInfo = &ConstantInterfaceMethodrefInfo{}
		case CONSTANT_String:
			cpInfo = &ConstantStringInfo{}
		case CONSTANT_Integer:
			cpInfo = &ConstantIntegerInfo{}
		case CONSTANT_Float:
			cpInfo = &ConstantFloatInfo{}
		case CONSTANT_Long:
			cpInfo = &ConstantLongInfo{}
		case CONSTANT_Double:
			cpInfo = &ConstantDoubleInfo{}
		case CONSTANT_NameAndType:
			cpInfo = &ConstantNameAndTypeInfo{}
		case CONSTANT_Utf8:
			cpInfo = &ConstantUtf8Info{}
		case CONSTANT_MethodHandle:
			cpInfo = &ConstantMethodHandleInfo{}
		case CONSTANT_MethodType:
			cpInfo = &ConstantMethodTypeInfo{}
		case CONSTANT_InvokeDynamic:
			cpInfo = &ConstantInvokeDynamicInfo{}
		default:
			break
		}
		if cpInfo != nil {
			cpInfo.ReadInfo(reader)
			constantPool[i] = cpInfo
		}
	}
	return constantPool
}
