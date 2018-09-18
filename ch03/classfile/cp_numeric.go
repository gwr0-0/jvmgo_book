package classfile

import "math"

/**
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
The tag item of the CONSTANT_Integer_info structure has the value CONSTANT_Integer (3).
The bytes item of the CONSTANT_Integer_info structure represents the value of the int constant.
The bytes of the value are stored in big-endian (high byte first) order.
*/
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

/**
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
The tag item of the CONSTANT_Float_info structure has the value CONSTANT_Float (4).
The bytes item of the CONSTANT_Float_info structure represents the value of the float constant in IEEE 754 floating-point single format (§2.3.2).
The bytes of the single format representation are stored in big-endian (high byte first) order.
*/
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

/**
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
The tag item of the CONSTANT_Long_info structure has the value CONSTANT_Long (5).
The unsigned high_bytes and low_bytes items of the CONSTANT_Long_info structure together represent the value of the long constant
((long) high_bytes << 32) + low_bytes
*/
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

/**
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
The tag item of the CONSTANT_Double_info structure has the value CONSTANT_Double (6).
The value represented by the CONSTANT_Double_info structure is determined as follows.
The high_bytes and low_bytes items are converted into the long constant bits, which is equal to
((long) high_bytes << 32) + low_bytes
*/
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
