package classpath

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry // 启动类路径
	extClasspath  Entry // 扩展类路径
	userClasspath Entry // 用户类路径
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}

	return cp
}

func (self *Classpath) readClass(className string) ([]byte, Entry, error) {

}

func (self *Classpath) String() string {

}

// 判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
