package main

import (
	"fmt"
	"github.com/gwr0-0/jvmgo/ch05/classfile"
	"github.com/gwr0-0/jvmgo/ch05/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

/**
1. 读取并解析class文件
2. 查找类的main方法
3. 解释执行main方法
*/
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethof(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {

}

func getMainMethof(cf *classfile.ClassFile) *classfile.MemberInfo {

}
