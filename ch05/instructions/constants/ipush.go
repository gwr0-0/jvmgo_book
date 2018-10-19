package constants

import "github.com/gwr0-0/jvmgo/ch05/instructions/base"

/**
Operation
	Push byte
Format
	bipush
	byte
Forms
	bipush = 16 (0x10)
Operand Stack
	... →
	..., value
Description
	The immediate byte is sign-extended to an int value. That value is pushed onto the operand stack.
*/
type BIPUSH struct {
	val int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

/**
Operation
	Push short
Format
	sipush
	byte1
	byte2
Forms
	sipush = 17 (0x11)
Operand Stack
	... →
	..., value
Description
	The immediate unsigned byte1 and byte2 values are assembled into an intermediate short, where the value of the short is (byte1 << 8) | byte2.
	The intermediate value is then sign-extended to an int value. That value is pushed onto the operand stack.
*/
type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
