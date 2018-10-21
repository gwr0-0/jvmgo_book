package stores

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Store reference into local variable
Format
	astore
	index
Forms
	astore = 58 (0x3a)
Operand Stack
	..., objectref →
	...
Description
	The index is an unsigned byte that must be an index into the local variable array of the current frame (§2.6).
	The objectref on the top of the operand stack must be of type returnAddress or of type reference.
	It is popped from the operand stack, and the value of the local variable at index is set to objectref.
Notes
	The astore instruction is used with an objectref of type returnAddress when implementing the finally clause of the Java programming language (§3.13).
	The aload instruction (§aload) cannot be used to load a value of type returnAddress from a local variable onto the operand stack.
	This asymmetry with the astore instruction is intentional.
	The astore opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.

Operation
	Store reference into local variable
Format
	astore_<n>
Forms
	astore_0 = 75 (0x4b)
	astore_1 = 76 (0x4c)
	astore_2 = 77 (0x4d)
	astore_3 = 78 (0x4e)
Operand Stack
	..., objectref →
	...
Description
	The <n> must be an index into the local variable array of the current frame (§2.6).
	The objectref on the top of the operand stack must be of type returnAddress or of type reference.
	It is popped from the operand stack, and the value of the local variable at <n> is set to objectref.
Notes
	An astore_<n> instruction is used with an objectref of type returnAddress when implementing the finally clauses of the Java programming language (§3.13).
	An aload_<n> instruction (§aload_<n>) cannot be used to load a value of type returnAddress from a local variable onto the operand stack.
	This asymmetry with the corresponding astore_<n> instruction is intentional.
	Each of the astore_<n> instructions is the same as astore with an index of <n>, except that the operand <n> is implicit.
*/

type ASTORE struct {
	base.Index8Instruction
}
type ASTORE_0 struct {
	base.NoOperandsInstruction
}
type ASTORE_1 struct {
	base.NoOperandsInstruction
}
type ASTORE_2 struct {
	base.NoOperandsInstruction
}
type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func _astore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}

func (self *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, uint(self.Index))
}
func (self *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}
func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}
func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}
func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
