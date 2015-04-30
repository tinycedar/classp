package classfile

/*
Class file format defined in JVM 7 specification

ClassFile {
    u4              magic;
    u2              minor_version;
    u2              major_version;
    u2              constant_pool_count;
    cp_info         constant_pool[constant_pool_count-1];
    u2              access_flags;
    u2              this_class;
    u2              super_class;
    u2              interfaces_count;
    u2              interfaces[interfaces_count];
    u2              fields_count;
    field_info      fields[fields_count];
    u2              methods_count;
    method_info     methods[methods_count];
    u2              attributes_count;
    attribute_info  attributes[attributes_count];
}
*/

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
}

func (self *ClassFile) ReadMagic(reader *ClassReader) {

}
