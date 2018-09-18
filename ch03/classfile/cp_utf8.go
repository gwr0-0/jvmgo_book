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

func (self *ConstantUtf8Info) readerInfo(reader *ClassReader) {

}

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(bytearr []byte) string {

}
