package base

import "github.com/gwr0-0/jvmgo/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader) // 从字节码中提取操作数
	Execute(frame *rtda.Frame)            // 执行指令逻辑
}

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
	Index uint // 局部变量表索引
}

// 存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint // 常量池索引
}

// 一些指令需要访问运行时常量池，常量池索引由两个字节操作数给出
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
