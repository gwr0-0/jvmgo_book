package classfile

/**
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
The tag item has the value CONSTANT_NameAndType (12).
The value of the name_index item must be a valid index into the constant_pool table.
The value of the descriptor_index item must be a valid index into the constant_pool table.

方法重载（overload），不同方法可以有相同的名字，只要参数列表不同即可。
同理，理论上字段也可实现重载，这里是java语法的限制
*/

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
