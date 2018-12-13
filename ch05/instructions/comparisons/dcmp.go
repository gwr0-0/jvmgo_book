package comparisons

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
dcmp<op>
Operation
	Compare double
Format
	dcmp<op>
Forms
	dcmpl = 151 (0x97)
	dcmpg = 152 (0x98)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type double. The values are popped from the operand stack and undergo value set conversion (§2.8.3), resulting in value1' and value2'. A floating-point comparison is performed:
		If value1' is greater than value2', the int value 1 is pushed onto the operand stack.
		Otherwise, if value1' is equal to value2', the int value 0 is pushed onto the operand stack.
		Otherwise, if value1' is less than value2', the int value -1 is pushed onto the operand stack.
		Otherwise, at least one of value1' or value2' is NaN. The dcmpg instruction pushes the int value 1 onto the operand stack and the dcmpl instruction pushes the int value -1 onto the operand stack.
	Floating-point comparison is performed in accordance with IEEE 754. All values other than NaN are ordered, with negative infinity less than all finite values and positive infinity greater than all finite values. Positive zero and negative zero are considered equal.
Notes
	The dcmpg and dcmpl instructions differ only in their treatment of a comparison involving NaN. NaN is unordered, so any double comparison fails if either or both of its operands are NaN. With both dcmpg and dcmpl available, any double comparison may be compiled to push the same result onto the operand stack whether the comparison fails on non-NaN values or fails because it encountered a NaN. For more information, see §3.5.
*/

type DCMPL struct {
	base.NoOperandsInstruction
}
type DCMPG struct {
	base.NoOperandsInstruction
}

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}
func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
