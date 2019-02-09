package control

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
goto
Operation
	Branch always
Format
	goto
	branchbyte1
	branchbyte2
Forms
	goto = 167 (0xa7)
Operand Stack
	No change
Description
	The unsigned bytes branchbyte1 and branchbyte2 are used to construct a signed 16-bit branchoffset, where branchoffset is (branchbyte1 << 8) | branchbyte2.
	Execution proceeds at that offset from the address of the opcode of this goto instruction.
	The target address must be that of an opcode of an instruction within the method that contains this goto instruction.
*/

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
