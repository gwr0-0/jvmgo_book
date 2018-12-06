package math

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
iinc
Operation
	Increment local variable by constant
Format
	iinc
	index
	const
Forms
	iinc = 132 (0x84)
Operand Stack
	No change
Description
	The index is an unsigned byte that must be an index into the local variable array of the current frame (§2.6).
	The const is an immediate signed byte. The local variable at index must contain an int.
	The value const is first sign-extended to an int, and then the local variable at index is incremented by that amount.
Notes
	The iinc opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index and to increment it by a two-byte immediate signed value.

*/

type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
