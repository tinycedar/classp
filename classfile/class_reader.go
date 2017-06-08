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

func (cr *ClassReader) ReadUint32() uint32 {
	value := bigEndian.Uint32(cr.bytecode[:4])
	cr.bytecode = cr.bytecode[4:]
	return value
}

func (cr *ClassReader) ReadUint16() uint16 {
	value := bigEndian.Uint16(cr.bytecode[:2])
	cr.bytecode = cr.bytecode[2:]
	return value
}

func (cr *ClassReader) ReadUint8() uint8 {
	return uint8(cr.ReadBytes(1)[0])
}

func (cr *ClassReader) ReadBytes(len int) []byte {
	bytes := cr.bytecode[:len]
	cr.bytecode = cr.bytecode[len:]
	return bytes
}

func (cr *ClassReader) Length() int {
	return len(cr.bytecode)
}
