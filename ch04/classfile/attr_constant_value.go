package classfile

/**
The ConstantValue attribute is a fixed-length attribute in the attributes table of a field_info structure.
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;  // 定长为2
    u2 constantvalue_index;
}
The value of the attribute_name_index item must be a valid index into the constant_pool table.
The constant_pool entry at that index must be a CONSTANT_Utf8_info structure (§4.4.7) representing the string "ConstantValue".
The value of the attribute_length item of a ConstantValue_attribute structure must be two.
The value of the constantvalue_index item must be a valid index into the constant_pool table
*/

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
