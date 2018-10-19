package loads

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Load long from local variable
Format
	lload
	index
Forms
	lload = 22 (0x16)
Operand Stack
	... →
	..., value
Description
	The index is an unsigned byte. Both index and index+1 must be indices into the local variable array of the current frame (§2.6).
	The local variable at index must contain a long. The value of the local variable at index is pushed onto the operand stack.
Notes
	The lload opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.

Operation
	Load long from local variable
Format
	lload_<n>
Forms
	lload_0 = 30 (0x1e)
	lload_1 = 31 (0x1f)
	lload_2 = 32 (0x20)
	lload_3 = 33 (0x21)
Operand Stack
	... →
	..., value
Description
	Both <n> and <n>+1 must be indices into the local variable array of the current frame (§2.6). The local variable at <n> must contain a long.
	The value of the local variable at <n> is pushed onto the operand stack.
Notes
	Each of the lload_<n> instructions is the same as lload with an index of <n>, except that the operand <n> is implicit.

*/

type LLOAD struct {
	base.Index8Instruction
}
type LLOAD_0 struct {
	base.NoOperandsInstruction
}
type LLOAD_1 struct {
	base.NoOperandsInstruction
}
type LLOAD_2 struct {
	base.NoOperandsInstruction
}
type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

func (self *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, self.Index)
}
func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}
func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}
func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}
func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
