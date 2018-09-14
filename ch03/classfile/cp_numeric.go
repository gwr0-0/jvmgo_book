package classfile

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

/**
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
The tag item of the CONSTANT_Double_info structure has the value CONSTANT_Double (6).
*/
type ConstantDoubleInfo struct {
	val float64
}
