package math

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Add int
Format
	iadd
Forms
	iadd = 96 (0x60)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type int. The values are popped from the operand stack. The int result is value1 + value2. The result is pushed onto the operand stack.
	The result is the 32 low-order bits of the true mathematical result in a sufficiently wide two's-complement format, represented as a value of type int.
	If overflow occurs, then the sign of the result may not be the same as the sign of the mathematical sum of the two values.
	Despite the fact that overflow may occur, execution of an iadd instruction never throws a run-time exception.

Operation
	Add long
Format
	ladd
Forms
	ladd = 97 (0x61)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type long. The values are popped from the operand stack. The long result is value1 + value2.
	The result is pushed onto the operand stack.
	The result is the 64 low-order bits of the true mathematical result in a sufficiently wide two's-complement format, represented as a value of type long.
	If overflow occurs, the sign of the result may not be the same as the sign of the mathematical sum of the two values.
	Despite the fact that overflow may occur, execution of an ladd instruction never throws a run-time exception.

Operation
	Add float
Format
	fadd
Forms
	fadd = 98 (0x62)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type float. The values are popped from the operand stack and undergo value set conversion (§2.8.3), resulting in value1' and value2'.
	The float result is value1' + value2'. The result is pushed onto the operand stack.
	The result of an fadd instruction is governed by the rules of IEEE arithmetic:
		If either value1' or value2' is NaN, the result is NaN.
		The sum of two infinities of opposite sign is NaN.
		The sum of two infinities of the same sign is the infinity of that sign.
		The sum of an infinity and any finite value is equal to the infinity.
		The sum of two zeroes of opposite sign is positive zero.
		The sum of two zeroes of the same sign is the zero of that sign.
		The sum of a zero and a nonzero finite value is equal to the nonzero value.
		The sum of two nonzero finite values of the same magnitude and opposite sign is positive zero.
		In the remaining cases, where neither operand is an infinity, a zero, or NaN and the values have the same sign or have different magnitudes,
		the sum is computed and rounded to the nearest representable value using IEEE 754 round to nearest mode.
		If the magnitude is too large to represent as a float, we say the operation overflows; the result is then an infinity of appropriate sign.
		If the magnitude is too small to represent as a float, we say the operation underflows; the result is then a zero of appropriate sign.
	The Java Virtual Machine requires support of gradual underflow as defined by IEEE 754.
	Despite the fact that overflow, underflow, or loss of precision may occur, execution of an fadd instruction never throws a run-time exception.

Operation
	Add double
Format
	dadd
Forms
	dadd = 99 (0x63)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type double. The values are popped from the operand stack and undergo value set conversion (§2.8.3),resulting in value1' and value2'.
	The double result is value1' + value2'. The result is pushed onto the operand stack.
	The result of a dadd instruction is governed by the rules of IEEE arithmetic:
		If either value1' or value2' is NaN, the result is NaN.
		The sum of two infinities of opposite sign is NaN.
		The sum of two infinities of the same sign is the infinity of that sign.
		The sum of an infinity and any finite value is equal to the infinity.
		The sum of two zeroes of opposite sign is positive zero.
		The sum of two zeroes of the same sign is the zero of that sign.
		The sum of a zero and a nonzero finite value is equal to the nonzero value.
		The sum of two nonzero finite values of the same magnitude and opposite sign is positive zero.
		In the remaining cases, where neither operand is an infinity, a zero, or NaN and the values have the same sign or have different magnitudes,
		the sum is computed and rounded to the nearest representable value using IEEE 754 round to nearest mode.
		If the magnitude is too large to represent as a double, we say the operation overflows; the result is then an infinity of appropriate sign.
		If the magnitude is too small to represent as a double, we say the operation underflows; the result is then a zero of appropriate sign.
	The Java Virtual Machine requires support of gradual underflow as defined by IEEE 754.
	Despite the fact that overflow, underflow, or loss of precision may occur, execution of a dadd instruction never throws a run-time exception.

*/

type IADD struct {
	base.NoOperandsInstruction
}
type LADD struct {
	base.NoOperandsInstruction
}
type FADD struct {
	base.NoOperandsInstruction
}
type DADD struct {
	base.NoOperandsInstruction
}

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}
func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}
func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}
func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}
