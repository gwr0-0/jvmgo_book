package classfile

/**
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
The value of the attribute_name_index item must be a valid index into the constant_pool table.
The constant_pool entry at that index must be a CONSTANT_Utf8_info structure (§4.4.7) representing the string "Deprecated".
The value of the attribute_length item must be zero.

Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
The value of the attribute_name_index item must be a valid index into the constant_pool table.
The constant_pool entry at that index must be a CONSTANT_Utf8_info structure (§4.4.7) representing the string "Synthetic".
The value of the attribute_length item must be zero.
*/

type MarkerAttribute struct{}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
