package stringplus

type Testcase struct {
	name  string
	len   uint64
	times uint64
}

var testcase = []Testcase{
	{
		name:  "2个长度15字符串拼接",
		len:   15,
		times: 2,
	},
	{
		name:  "2个长度50字符串拼接",
		len:   10,
		times: 2,
	},
	{
		name:  "100个长度10字符串拼接",
		len:   10,
		times: 100,
	},
	{
		name:  "1000个长度10字符串拼接",
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
