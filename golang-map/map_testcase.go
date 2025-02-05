package gomap

type Testcase struct {
	msize uint64 // size of map
	name  string // testcase name
	read  uint64 // read goroutine
	write uint64 // write goroutine
	times uint64 // read or write times
}

var testcase = []Testcase{
	{
		name:  "10000s-30r0w",
		msize: 1000,
		read:  10,
		write: 0,
		times: 10000,
	},
	{
		name:  "10000s-0r10w",
		msize: 1000,
		read:  0,
		write: 10,
		times: 10000,
	},
}
