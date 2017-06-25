package classfile

import "fmt"

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	cp         ConstantPool
	MaxStack   uint16
	MaxLocals  uint16
	Code       []uint8          //u4 code_length
	Exceptions []exceptionTable //u2 exception_table_length
	Attributes []AttributeInfo  //u2 attributes_count
}

func (this *CodeAttribute) ReadInfo(reader *ClassReader) {
	this.MaxStack = reader.ReadUint16()
	this.MaxLocals = reader.ReadUint16()
	this.Code = reader.ReadBytes(int(reader.ReadUint32()))
	//parseCode(this.Code)
	exceptionTableLength := reader.ReadUint16()
	this.Exceptions = make([]exceptionTable, exceptionTableLength)
	for i := uint16(0); i < exceptionTableLength; i++ {
		exceptionTable := exceptionTable{}
		exceptionTable.ReadInfo(reader)
		this.Exceptions[i] = exceptionTable
	}
	this.Attributes = readAttributes(reader, this.cp)
}

func parseCode(code []uint8) {
	fmt.Printf("code: %v\n", code)
	for _, c := range code {
		switch c {
		case 0x2a:
			fmt.Printf("%v, ", "aload_0")
		case 0x2b:
			fmt.Printf("%v, ", "aload_1")
		case 0xb7:
			fmt.Printf("%v, ", "invokespecial")
		case 0x0:
			fmt.Printf("%v, ", "nop")
		case 0x1:
			fmt.Printf("%v, ", "aconst_null")
		case 0x2:
			fmt.Printf("%v, ", "iconst_m1")
		case 0xb1:
			fmt.Printf("%v, ", "return")
		case 176:
			fmt.Printf("%v, ", "areturn")
		case 180:
			fmt.Printf("%v, ", "getfield")
		case 181:
			fmt.Printf("%v, ", "putfield")
		default:
			fmt.Printf("%v, ", "invalid opcode")
		}
	}
	fmt.Println()
}

type exceptionTable struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (this *exceptionTable) ReadInfo(reader *ClassReader) {
	this.startPc = reader.ReadUint16()
	this.endPc = reader.ReadUint16()
	this.handlerPc = reader.ReadUint16()
	this.catchType = reader.ReadUint16()
}
