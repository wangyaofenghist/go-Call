package main

import (
	"fmt"
	"localhostTest/go-Call/call"
	"localhostTest/go-Call/test"
)

func main() {
	call := call.CreateCall()
	call.AddCall("test1", test.Test1)
	call.AddCall("test4", test.Test4)
	call.ReplaceCall("test1", test.Test1)
	call.ReplaceCall("call", call.Call)

	s := "test1"
	result1, err := call.Call(s)
	result2, _ := call.Call("test4", "hello ", "world")
	result3, _ := call.Call("test1", "ni  ", "hao ")
	call.RemoveCall(s)
	fmt.Println(result1, err, result2, result3)

	result1, err = call.Call(s)
	fmt.Println(result1, err)
	result3, err = call.Call("call", "test4", "hello 2", "world 2")
	fmt.Println(result3, err, 2222)
	//会报错
	/*call.ReplaceCall("test4", "test.Test4")
	result2, err = call.Call("test4")
	fmt.Println(result2, err)*/
}
