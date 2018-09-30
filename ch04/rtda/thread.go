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

}

// setter
func (self *Thread) SetPC(pc int) {

}

func (self *Thread) PushFrame(frame *Frame) {

}

func (self *Thread) PopFrame() *Frame {

}

// 返回当前帧
func (self *Thread) CurrentFrame() *Frame {

}
