package rtda

type Thread struct {
	pc    int    // program counter，程序计数器
	stack *Stack // Java Virtual Machine stack，Java虚拟机栈
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

// getter
func (self *Thread) PC() int {
	return self.pc
}

// setter
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

// 返回当前帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
