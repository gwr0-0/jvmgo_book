package math

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
irem = 112 (0x70)
lrem = 113 (0x71)
frem = 114 (0x72)
drem = 115 (0x73)

Operand Stack
..., value1, value2 →
..., result
	value1 - (value1 / value2) * value2
*/

type IREM struct {
	base.NoOperandsInstruction
}
type LREM struct {
	base.NoOperandsInstruction
}
type FREM struct {
	base.NoOperandsInstruction
}
type DREM struct {
	base.NoOperandsInstruction
}

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}
