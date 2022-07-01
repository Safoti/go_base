package main
//接口嵌套
//模拟一个
type Buffer struct {

}
//文件接口
type ReadWrite  interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}
//加锁接口
type Lock interface {
	Lock()
	Unlock()
}
//文件接口
type File  interface {
	ReadWrite  //使用读写
	Lock  //使用锁接口
	Close() //自己还有一个关闭接口
}
func main() {

}
