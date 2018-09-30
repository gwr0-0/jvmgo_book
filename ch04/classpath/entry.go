package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator) // OS-specific path list separator

type Entry interface {
	readClass(className string) ([]byte, Entry, error) // 负责寻找和加载class文件，根据参数创建不同类型的Entry实例
	String() string                                    // 用于返回变量的字符串表示（相当于toString()）
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
