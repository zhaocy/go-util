package go_util

import (
	"runtime"
	"os"
	"os/signal"
	"syscall"
)

// 产生panic时的调用栈打印
func PrintPanicStack(extras ...interface{}) {
	if x := recover(); x != nil {
		Errorf("%v", x)
		i := 0
		funcName, file, line, ok := runtime.Caller(i)
		for ok {
			Errorf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
			i++
			funcName, file, line, ok = runtime.Caller(i)
		}

		for k := range extras {
			Errorf("EXRAS#%v DATA:%v\n", k, extras[k])
		}
	}
}

func WaitSignal() os.Signal {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	return <-ch
}