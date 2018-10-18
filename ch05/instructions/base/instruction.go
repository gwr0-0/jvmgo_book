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

// 跳转指令
type BranchInstruction struct {
	Offset int // 跳转偏移量
}

// 从字节码中读取一个int16整数，转成int，赋值给Offset
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}
