package rtda

type Stack struct {
	maxSize uint   // 栈的容量（最多可以容纳多少帧）
	size    uint   // 栈的当前大小
	_top    *Frame // 栈顶指针
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// 把帧推入栈顶
func (self *Stack) push(frame *Frame) {
	// 按java虚拟机规范，若栈已经满了，应抛出StackOverflowError异常
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size++

	return top
}

// 返回栈顶帧
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}
