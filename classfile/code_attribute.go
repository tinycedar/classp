package classfile

import (
// "fmt"
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
	// attributeNameIndex uint16
	// attributeLength    uint32
	maxStack   uint16
	maxLocals  uint16
	code       []uint8          //u4 code_length
	exceptions []exceptionTable //u2 exception_table_length
	Attributes []AttributeInfo  //u2 attributes_count
}

func (this *CodeAttribute) ReadInfo(reader *ClassReader) {
	this.maxStack = reader.ReadUint16()
	this.maxLocals = reader.ReadUint16()
	this.code = reader.ReadBytes(int(reader.ReadUint32()))
	exceptionTableLength := reader.ReadUint16()
	this.exceptions = make([]exceptionTable, exceptionTableLength)
	for i := uint16(0); i < exceptionTableLength; i++ {
		exceptionTable := exceptionTable{}
		exceptionTable.ReadInfo(reader)
		this.exceptions[i] = exceptionTable
	}
	attributesCount := reader.ReadUint16()
	for i := uint16(0); i < attributesCount; i++ {
		// fmt.Printf("attributes in attributes.....%d\n", reader.ReadUint16())
		readAttributeInfo(reader)
	}
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

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
	lineNumberTables []lineNumberTable
}

type lineNumberTable struct {
	startPc    uint16
	lineNumber uint16
}

func (this *LineNumberTableAttribute) ReadInfo(reader *ClassReader) {
	lineNumberTableLength := reader.ReadUint16()
	this.lineNumberTables = make([]lineNumberTable, lineNumberTableLength)
	for i := uint16(0); i < lineNumberTableLength; i++ {
		this.lineNumberTables[i] = lineNumberTable{reader.ReadUint16(), reader.ReadUint16()}
	}
}
