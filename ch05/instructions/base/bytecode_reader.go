package base

type BytecodeReader struct {
	code []byte // 存放字节码
	pc   int    // 记录读取到了哪个字节
}
