package utilswapper

import "runtime"

//获取调用堆栈
func GetStack(all bool) string {
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], all)])
}

// 获取正在运行的函数名
func CurrFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
