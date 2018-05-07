//一个通用的回调函数
package call

import (
	"github.com/pkg/errors"
	"reflect"
)

type Params interface{}
type CallMap map[string]interface{}

//创建一个map
func CreateCall() CallMap {
	return make(CallMap)
}

//添加一个回调 重复绑定 将会替换
func (call CallMap) AddCall(name string, param Params) (err error) {
	if _, ok := call[name]; !ok {
		call[name] = param
	}
	return
}

//移除一个方法
func (call CallMap) RemoveCall(name string) {
	if _, ok := call[name]; ok {
		delete(call, name)
	}
	return
}

//替换
func (call CallMap) ReplaceCall(name string, param Params) {
	call[name] = param
	return
}

//调用回调
func (call CallMap) Call(name string, param ...Params) (result []reflect.Value, err error) {
	if _, ok := call[name]; !ok {
		err = errors.New("name:" + name + " is not find.")
		return
	}
	callBack := reflect.ValueOf(call[name])
	if len(param) != callBack.Type().NumIn() {
		err = errors.New("The number of Params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(param))
	for k, param := range param {
		in[k] = reflect.ValueOf(param)
	}
	result = callBack.Call(in)
	defer call.checkError()
	return
}

//通用错误日志
func (call CallMap) checkError() {
	if e := recover(); e != nil {
		panic(e)
	}
}
