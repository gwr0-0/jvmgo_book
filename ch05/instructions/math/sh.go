package math

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
ishl
Operation
	Shift left int
Format
	ishl
Forms
	ishl = 120 (0x78)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type int. The values are popped from the operand stack.
	An int result is calculated by shifting value1 left by s bit positions, where s is the value of the low 5 bits of value2.
	The result is pushed onto the operand stack.
Notes
	This is equivalent (even if overflow occurs) to multiplication by 2 to the power s.
	The shift distance actually used is always in the range 0 to 31, inclusive, as if value2 were subjected to a bitwise logical AND with the mask value 0x1f.

lshl
Operation
	Shift left long
Format
	lshl
Forms
	lshl = 121 (0x79)
Operand Stack
	..., value1, value2 →
	..., result
Description
	The value1 must be of type long, and value2 must be of type int. The values are popped from the operand stack.
	A long result is calculated by shifting value1 left by s bit positions, where s is the low 6 bits of value2. The result is pushed onto the operand stack.
Notes
	This is equivalent (even if overflow occurs) to multiplication by 2 to the power s.
	The shift distance actually used is therefore always in the range 0 to 63, inclusive, as if value2 were subjected to a bitwise logical AND with the mask value 0x3f.

ishr
Operation
	Arithmetic shift right int
Format
	ishr
Forms
	ishr = 122 (0x7a)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type int. The values are popped from the operand stack.
	An int result is calculated by shifting value1 right by s bit positions, with sign extension, where s is the value of the low 5 bits of value2.
	The result is pushed onto the operand stack.
Notes
	The resulting value is floor(value1 / 2s), where s is value2 & 0x1f. For non-negative value1, this is equivalent to truncating int division by 2 to the power s.
	The shift distance actually used is always in the range 0 to 31, inclusive, as if value2 were subjected to a bitwise logical AND with the mask value 0x1f.

lshr
Operation
	Arithmetic shift right long
Format
	lshr
Forms
	lshr = 123 (0x7b)
Operand Stack
	..., value1, value2 →
	..., result
Description
	The value1 must be of type long, and value2 must be of type int. The values are popped from the operand stack.
	A long result is calculated by shifting value1 right by s bit positions, with sign extension, where s is the value of the low 6 bits of value2.
	The result is pushed onto the operand stack.
Notes
	The resulting value is floor(value1 / 2s), where s is value2 & 0x3f. For non-negative value1, this is equivalent to truncating long division by 2 to the power s.
	The shift distance actually used is therefore always in the range 0 to 63, inclusive, as if value2 were subjected to a bitwise logical AND with the mask value 0x3f.

iushr
Operation
	Logical shift right int
Format
	iushr
Forms
	iushr = 124 (0x7c)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type int. The values are popped from the operand stack.
	An int result is calculated by shifting value1 right by s bit positions, with zero extension, where s is the value of the low 5 bits of value2.
	The result is pushed onto the operand stack.
Notes
	If value1 is positive and s is value2 & 0x1f, the result is the same as that of value1 >> s;
	if value1 is negative, the result is equal to the value of the expression (value1 >> s) + (2 << ~s).
	The addition of the (2 << ~s) term cancels out the propagated sign bit. The shift distance actually used is always in the range 0 to 31, inclusive.

lushr
Operation
	Logical shift right long
Format
	lushr
Forms
	lushr = 125 (0x7d)
Operand Stack
	..., value1, value2 →
	..., result
Description
	The value1 must be of type long, and value2 must be of type int. The values are popped from the operand stack.
	A long result is calculated by shifting value1 right logically by s bit positions, with zero extension, where s is the value of the low 6 bits of value2.
	The result is pushed onto the operand stack.
Notes
	If value1 is positive and s is value2 & 0x3f, the result is the same as that of value1 >> s;
	if value1 is negative, the result is equal to the value of the expression (value1 >> s) + (2L << ~s).
	The addition of the (2L << ~s) term cancels out the propagated sign bit. The shift distance actually used is always in the range 0 to 63, inclusive.

*/

type ISHL struct {
	base.NoOperandsInstruction
}
type LSHL struct {
	base.NoOperandsInstruction
}
type ISHR struct {
	base.NoOperandsInstruction
}
type LSHR struct {
	base.NoOperandsInstruction
}
type IUSHR struct {
	base.NoOperandsInstruction
}
type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}
func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}
func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}
func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}
func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}
func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
