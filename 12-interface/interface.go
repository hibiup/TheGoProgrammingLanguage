package interfaces

import (
	"os"
)

/**
  给 int 类型指针加上 Write 接口方法. 因为不能直接给 non-local 类型加方法，因此要将 int 先定义一个本地的 ByteCounter 类型，内嵌int
*/
type ByteCounter int

// 为 Bytecounter 实现 Write 方法，该方法在接收一个 []byte 数据的时候返回它的长度。
func (c *ByteCounter) Write(p []byte) (int, error) {
	length := len(p)
	os.Stdout.Close()
	*c += ByteCounter(length) // convert int to ByteCounter
	return length, nil
}

// Recipter.Write 获得输入（name）后加上问候语
type Recipter string

func (r *Recipter) Write(name []byte) (int, error) {
	buffer := []byte(*r)
	buffer = append(buffer, name...)
	*r = Recipter(buffer)
	return 0, nil
}
