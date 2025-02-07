package stringplus

type Testcase struct {
	name  string
	len   uint64
	times uint64
}

var testcase = []Testcase{
	{
		name:  "len15 * 2",
		len:   15,
		times: 2,
	},
	{
		name:  "len50 * 2",
		len:   10,
		times: 2,
	},
	{
		name:  "len10 * 100",
		len:   10,
		times: 100,
	},
	{
		name:  "len10 * 1000",
		len:   10,
		times: 1000,
	},
}

func genStr(len uint64) string {
	str := ""
	for i := 0; i < int(len); i++ {
		str += "#"
	}
	return str
}
