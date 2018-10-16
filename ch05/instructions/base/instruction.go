package base

import "github.com/gwr0-0/jvmgo/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader) // 从字节码中提取操作数
	Execute(frame *rtda.Frame)            // 执行指令逻辑
}

/**
nop
Operation
	Do nothing
Format
	nop
Forms
	nop = 0 (0x0)
Operand Stack
	No change
Description
	Do nothing.
*/
// 表示没有操作数的指令，所有没有定义任何字段
type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}
