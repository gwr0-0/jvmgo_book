package main

import (
	"fmt"
	"github.com/gwr0-0/jvmgo/ch05/classfile"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

/**
loop函数循环执行 计算PC、解码指令、执行指令 这三个步骤，直到遇到错误
*/
func loop(thread *rtda.Thread, bytecode []byte) {

}
