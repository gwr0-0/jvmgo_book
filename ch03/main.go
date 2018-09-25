package main

import (
	"fmt"
	"github.com/gwr0-0/jvmgo/ch02/classpath"
	"github.com/gwr0-0/jvmgo/ch03/classfile"
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

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)

	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {

	fmt.Println("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
}
