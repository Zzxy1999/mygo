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

func StrAppendTest() {
	var s []byte
	fmt.Printf("cap: %d\t", cap(s))
	printStrInfo(string(s))
	s = append(s, "1"...)
	fmt.Printf("cap: %d\t", cap(s))
	printStrInfo(string(s))
	s = append(s, "1234"...)
	fmt.Printf("cap: %d\t", cap(s))
	printStrInfo(string(s))
}

func SprintfDebug() {
	s1 := "0123456789" // strconv.Itoa(rand.Int())
	s2 := "0123456789" //strconv.Itoa(rand.Int())
	s := fmt.Sprintf("%s%s", s1, s2)
	fmt.Println(s)
}
