package singleExample

import "sync"

type SingleExa_lh struct {
}

var singleExa_lh *SingleExa_lh

//sync.Once go 只执行一次的函数  常在单例模式中使用  用于初始化函数
var once = &sync.Once{}

func init_lh() {
	singleExa_lh = &SingleExa_lh{}
}

func GetInstance_lh() *SingleExa_lh {

	if singleExa_lh == nil {
		once.Do(init_lh)
	}
	return singleExa_lh
}
