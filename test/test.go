package test

import "time"

func Test1() string {
	return "test1"
}
func Test2() string {
	return "test2"
}
func Test3() string {
	return "test3"
}
func Test4(s1 string, s2 string) string {
	time.Sleep(time.Millisecond * 1000)
	return s1 + s2 + " this is test"
}
func Hi5(s1 int, s2 int) (result int) {
	result = s1 + s2
	return
}
