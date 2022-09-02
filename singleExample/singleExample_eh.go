package singleExample

// 唯一数据
type SingleExa_eh struct {
}

// 初始化单例
var singleExa_eh *SingleExa_eh

func init() {
	singleExa_eh = &SingleExa_eh{}
}

// 获取单例
func GetInstance_eh() *SingleExa_eh {
	return singleExa_eh
}
