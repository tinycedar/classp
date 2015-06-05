package classfile

import (
	"encoding/binary"
)

var bigEndian = binary.BigEndian

type ClassReader struct {
	bytecode []byte
}

func NewClassReader(bytecode []byte) *ClassReader {
	return &ClassReader{bytecode: bytecode}
}

func (this *ClassReader) ReadUint32() uint32 {
	value := bigEndian.Uint32(this.bytecode[:4])
	this.bytecode = this.bytecode[4:]
	return value
}

func (this *ClassReader) ReadUint16() uint16 {
	value := bigEndian.Uint16(this.bytecode[:2])
	this.bytecode = this.bytecode[2:]
	return value
}

func (this *ClassReader) ReadUint8() uint8 {
	return uint8(this.ReadBytes(1)[0])
}

func (this *ClassReader) ReadBytes(len int) []byte {
	bytes := this.bytecode[:len]
	this.bytecode = this.bytecode[len:]
	return bytes
}

func (this *ClassReader) Length() int {
	return len(this.bytecode)
}
