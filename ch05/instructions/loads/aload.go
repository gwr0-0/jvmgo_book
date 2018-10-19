package loads

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Load reference from local variable
Format
	aload
	index
Forms
	aload = 25 (0x19)
Operand Stack
	... →
	..., objectref
Description
	The index is an unsigned byte that must be an index into the local variable array of the current frame (§2.6).
The local variable at index must contain a reference. The objectref in the local variable at index is pushed onto the operand stack.
Notes
	The aload instruction cannot be used to load a value of type returnAddress from a local variable onto the operand stack.
	This asymmetry with the astore instruction (§astore) is intentional.
	The aload opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.


Operation
	Load reference from local variable
Format
	aload_<n>
Forms
	aload_0 = 42 (0x2a)
	aload_1 = 43 (0x2b)
	aload_2 = 44 (0x2c)
	aload_3 = 45 (0x2d)
Operand Stack
	... →
	..., objectref
Description
	The <n> must be an index into the local variable array of the current frame (§2.6). The local variable at <n> must contain a reference.
	The objectref in the local variable at <n> is pushed onto the operand stack.
Notes
	An aload_<n> instruction cannot be used to load a value of type returnAddress from a local variable onto the operand stack.
	This asymmetry with the corresponding astore_<n> instruction (§astore_<n>) is intentional.
	Each of the aload_<n> instructions is the same as aload with an index of <n>, except that the operand <n> is implicit.
*/

type ALOAD struct {
	base.Index8Instruction
}
type ALOAD_0 struct {
	base.NoOperandsInstruction
}
type ALOAD_1 struct {
	base.NoOperandsInstruction
}
type ALOAD_2 struct {
	base.NoOperandsInstruction
}
type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func _aload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, self.Index)
}
func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}
func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}
func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}
func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}
