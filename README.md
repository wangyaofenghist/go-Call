# go-Call

利用golang 反射机制，实现一个通用回调函数，方便开发过程中使用

###实例

```
call := call.CreateCall()        		//得到一个回调句柄
call.AddCall("test1", test.Test1)		//添加一个回调
call.AddCall("test4", test.Test4)
call.ReplaceCall("test1", test.Test2)	//重置一个回调
call.ReplaceCall("call", call.Call)

s := "test1"
result1, err := call.Call(s)			
result2, _ := call.Call("test4", "hello ", "world") //调用回调并传输参数
```

### 基于go-worker-base实现demo 

可以在协程数可控的情况下实现，方便的实现业务并发执行

go-worker-base 地址：https://github.com/wangyaofenghist/go-worker-base

```
package main

import (
   "fmt"
   "github.com/wangyaofenghist/go-Call/call"
   "github.com/wangyaofenghist/go-Call/test"
   "github.com/wangyaofenghist/go-worker-base/worker"
   "time"
)

//声明一号池子
var poolOne worker.WorkPool

//声明二号池子 待用
var poolTwo worker.WorkPool

//声明回调变量
var funcs call.CallMap

//以结构体方式调用
type runWorker struct{}

//初始化协程池 和回调参数
func init() {
   poolOne = worker.GetPool("one")
   poolOne.Start(50)
   funcs = call.CreateCall()

}

//通用回调
func (f *runWorker) Run(param []interface{}) {
   name := param[0].(string)
   //调用回调并拿回结果
   funcs.Call(name, param[1:]...)
}

//主函数
func main() {
   var runFunc runWorker = runWorker{}
   funcs.AddCall("test4", test.Test4)
   var startTime = time.Now().UnixNano()
   for i := 0; i < 10000; i++ {
      poolOne.Run(runFunc.Run, "test4", " aa ", " BB")
      poolOne.Run(runFunc.Run, "test4", " cc ", " dd")
      poolOne.Run(runFunc.Run, "test4", " ee ", " ff")
   }
   var modTime = time.Now().UnixNano()

   for k := 0; k < 10000; k++ {
      test.Test4(" aa ", "BB")
      test.Test4(" cc ", " dd")
      test.Test4(" ee ", " ff")
   }
   var endTime = time.Now().UnixNano()
   for j := 0; j < 10000; j++ {
      funcs.Call("test4", " aa ", "BB")
      funcs.Call("test4", " cc ", " dd")
      funcs.Call("test4", " ee ", " ff")
   }
   var lastTime = time.Now().UnixNano()
   fmt.Println(modTime - startTime)
   fmt.Println(endTime - modTime)
   fmt.Println(lastTime - endTime)

   fmt.Println(startTime, modTime, endTime)
   time.Sleep(time.Millisecond * 1000)
   poolOne.Stop()
   
}
```