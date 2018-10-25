package classfile

/*
LocalVariableTypeTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_type_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 signature_index;
        u2 index;
    } local_variable_type_table[local_variable_type_table_length];
}

用于调试器，以获取在方法运行时泛型局部变量的信息
LocalVariableTable Attribute和LocalVariableTypeTable Attribute表达的信息是类似的，
他们的区别是对泛型类型的局部变量，需要用Signature的形式表达，而不能仅仅用Descriptor的形式表达，
因而对泛型类型的局部变量，需要在LocalVariableTable Attribute和LocalVariableTypeTable Attribute中同时存在一项；
而对非泛型类型的局部变量来说，只要在LocalVariableTable Attribute存在表项就可以了。
从这里我们也可以看出泛型是后期才被字节码所支持的痕迹
*/

type LocalVariableTypeTableAttribute struct {
	localVariableTypeTable []*LocalVariableTypeTableEntry
}

type LocalVariableTypeTableEntry struct {
	startPc        uint16
	length         uint16
	nameIndex      uint16
	signatureIndex uint16
	index          uint16
}

func (self *LocalVariableTypeTableAttribute) readInfo(reader *ClassReader) {
	localVariableTypeTableLength := reader.readUint16()
	self.localVariableTypeTable = make([]*LocalVariableTypeTableEntry, localVariableTypeTableLength)
	for i := range self.localVariableTypeTable {
		self.localVariableTypeTable[i] = &LocalVariableTypeTableEntry{
			startPc:        reader.readUint16(),
			length:         reader.readUint16(),
			nameIndex:      reader.readUint16(),
			signatureIndex: reader.readUint16(),
			index:          reader.readUint16(),
		}
	}
}
