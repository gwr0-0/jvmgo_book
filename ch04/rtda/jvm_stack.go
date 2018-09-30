package rtda

type Stack struct {
	maxSize uint   // 栈的容量（最多可以容纳多少帧）
	size    uint   // 栈的当前大小
	_top    *Frame //栈顶指针
}

func newStack(maxSize uint) *Stack {

}
