package classfile

/**
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
The tag item has the value CONSTANT_String (8).
The value of the string_index item must be a valid index into the constant_pool table.
The constant_pool entry at that index must be a CONSTANT_Utf8_info structure representing the sequence of Unicode code points to which the String object is to be initialized.
*/

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
