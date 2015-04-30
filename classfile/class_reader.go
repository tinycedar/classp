package classfile

import (
	"encoding/binary"
)

var bigEndian = binary.BigEndian

type ClassReader struct {
	bytecode []byte
}

func NewClassReader(bytes []byte) *ClassReader {
	return &ClassReader{bytes}
}

func (self *ClassReader) Length() int {
	return len(self.bytecode)
}

func (self *ClassReader) ReadUint8() uint8 {
	val := self.bytecode[0]
	self.bytecode = self.bytecode[1:]
	return val
}

func (self *ClassReader) ReadUint16() uint16 {
	val := bigEndian.Uint16(self.bytecode)
	self.bytecode = self.bytecode[2:]
	return val
}

func (self *ClassReader) ReadUint32() uint32 {
	val := bigEndian.Uint32(self.bytecode)
	self.bytecode = self.bytecode[4:]
	return val
}
