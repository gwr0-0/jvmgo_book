package loads

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Load float from local variable
Format
	fload
	index
Forms
	fload = 23 (0x17)
Operand Stack
	... →
	..., value
Description
	The index is an unsigned byte that must be an index into the local variable array of the current frame (§2.6).
The local variable at index must contain a float. The value of the local variable at index is pushed onto the operand stack.
Notes
	The fload opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.

Operation
	Load float from local variable
Format
	fload_<n>
Forms
	fload_0 = 34 (0x22)
	fload_1 = 35 (0x23)
	fload_2 = 36 (0x24)
	fload_3 = 37 (0x25)
Operand Stack
	... →
	..., value
Description
	The <n> must be an index into the local variable array of the current frame (§2.6). The local variable at <n> must contain a float.
	The value of the local variable at <n> is pushed onto the operand stack.
Notes
	Each of the fload_<n> instructions is the same as fload with an index of <n>, except that the operand <n> is implicit.
*/
type FLOAD struct {
	base.Index8Instruction
}
type FLOAD_0 struct {
	base.NoOperandsInstruction
}
type FLOAD_1 struct {
	base.NoOperandsInstruction
}
type FLOAD_2 struct {
	base.NoOperandsInstruction
}
type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

func (self *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, self.Index)
}
func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}
func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}
func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}
func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
