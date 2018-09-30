package classfile

/**
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
The tag item has the value CONSTANT_Utf8 (1).
The value of the length item gives the number of bytes in the bytes array (not the length of the resulting string).
The bytes array contains the bytes of the string. No byte may have the value (byte)0. No byte may lie in the range (byte)0xf0 to (byte)0xff.
*/

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
// TODO 这里是简化版，字符串中不能包含null字符（代码点U+0000）或补充字符（代码点大于U+FFFF的Unicode字符）
// 完整版见，https://github.com/gwr0-0/jvmgo-book/blob/master/v1/code/go/src/jvmgo/ch03/classfile/cp_utf8.go
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
