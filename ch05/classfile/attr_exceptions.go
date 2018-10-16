package classfile

/**
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
The value of the attribute_name_index item must be a valid index into the constant_pool table.
The constant_pool entry at that index must be the CONSTANT_Utf8_info structure (§4.4.7) representing the string "Exceptions".
The value of the attribute_length item indicates the attribute length, excluding the initial six bytes.
The value of the number_of_exceptions item indicates the number of entries in the exception_index_table.
Each value in the exception_index_table array must be a valid index into the constant_pool table
The constant_pool entry at that index must be a CONSTANT_Class_info structure (§4.4.1) representing a class type that this method is declared to throw.

变长属性，记录方法抛出的异常表
*/

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
