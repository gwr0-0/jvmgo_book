package classfile

/**
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
The tag item has the value CONSTANT_Class (7).
The value of the name_index item must be a valid index into the constant_pool table.
The constant_pool entry at that index must be a CONSTANT_Utf8_info structure representing a valid binary class or interface name encoded in internal form (§4.2.1).
https://docs.oracle.com/javase/specs/jvms/se10/html/jvms-4.html#jvms-4.2.1
*/

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
