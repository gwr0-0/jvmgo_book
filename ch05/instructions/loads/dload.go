package loads

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Load double from local variable
Format
	dload
	index
Forms
	dload = 24 (0x18)
Operand Stack
	... →
	..., value
Description
	The index is an unsigned byte. Both index and index+1 must be indices into the local variable array of the current frame (§2.6).
	The local variable at index must contain a double. The value of the local variable at index is pushed onto the operand stack.
Notes
	The dload opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.

Operation
	Load double from local variable
Format
	dload_<n>
Forms
	dload_0 = 38 (0x26)
	dload_1 = 39 (0x27)
	dload_2 = 40 (0x28)
	dload_3 = 41 (0x29)
Operand Stack
	... →
	..., value
Description
	Both <n> and <n>+1 must be indices into the local variable array of the current frame (§2.6). The local variable at <n> must contain a double.
	The value of the local variable at <n> is pushed onto the operand stack.
Notes
	Each of the dload_<n> instructions is the same as dload with an index of <n>, except that the operand <n> is implicit.
*/

type DLOAD struct {
	base.Index8Instruction
}
type DLOAD_0 struct {
	base.NoOperandsInstruction
}
type DLOAD_1 struct {
	base.NoOperandsInstruction
}
type DLOAD_2 struct {
	base.NoOperandsInstruction
}
type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, self.Index)
}
func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}
func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}
func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}
func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
