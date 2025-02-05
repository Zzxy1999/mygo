package gomap

type Testcase struct {
	name   string // testcase name
	msize  uint64 // size of map
	isread bool   // is read?
	count  uint64 // goroutines
}

var testcase = []Testcase{
	{
		msize:  10000,
		name:   "read 10000",
		isread: true,
		count:  10000,
	},
}
