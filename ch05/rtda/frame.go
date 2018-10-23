package rtda

type Frame struct {
	lower        *Frame        // 用来实现链表数据结构
	localVars    LocalVars     // 局部变量表指针
	operandStack *OperandStack // 操作数栈指针
	thread       *Thread
	nextPC       int
}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),   // 保存局部变量表指针
		operandStack: newOperandStack(maxStack), // 保存操作数栈指针
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
