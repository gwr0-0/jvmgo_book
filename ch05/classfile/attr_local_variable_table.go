package classfile

/**
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}

用于调试器，以获取在方法运行时局部变量的信息
*/

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16 // 记录该局部变量的有值范围，字节码数组中的索引范围[start_pc, start_pc+length)
	length          uint16 // 记录该局部变量的有值范围，字节码数组中的索引范围[start_pc, start_pc+length)
	nameIndex       uint16 // constant_pool中索引，CONSTANT_Utf8_info类型，记录该项代表的局部变量名。
	descriptorIndex uint16 // constant_pool中索引，CONSTANT_Utf8_info类型，记录该项代表的局部变量的字段描述符
	index           uint16 // 记录该项代表的局部变量在方法的局部变量数组中的索引
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}
