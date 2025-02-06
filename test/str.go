package test

import (
	"fmt"
	"reflect"
	"unsafe"
)

func printStrInfo(s string) {
	fmt.Printf("&s: %x\n", &s)
	strHeader := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("s.data: %x, s.len: %d\n\n", strHeader.Data, strHeader.Len)
}

func strCopy(s string) {
	printStrInfo(s)
	s += "123"
	printStrInfo(s)
}

func StrCopyTest() {
	s := "test"
	printStrInfo(s)
	strCopy(s)
	printStrInfo(s)
}

func SprintfDebug() {
	s1 := "123" // strconv.Itoa(rand.Int())
	s2 := "456" //strconv.Itoa(rand.Int())
	s := fmt.Sprintf("%s%s", s1, s2)
	fmt.Println(s)
}
