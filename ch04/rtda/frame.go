package rtda

type Frame struct {
	lower        *Frame        // 用来实现链表数据结构
	localVars    LocalVars     // 局部变量表指针
	operandStack *OperandStack // 操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame {

}
