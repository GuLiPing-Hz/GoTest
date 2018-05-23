package codeerror

import "fmt"

type CodeError struct {
	//小写是私有变量 private
	code int
	msg  string
	//大写是公开变量 public
	Reserve string
}

//实现error接口,
func (e *CodeError) Error() string { //函数前面 声明 结构的实现 interface error
	return getErrorString(e)
}

//大写是公开方法
func New(code int, msg string) error {
	//构造一个对象,没有指定参数名，则必须所有值都必须初始化
	//CodeError{code, msg, 0}
	//构造一个对象,指定参数名，可以省略一些参数初始化,其他采用默认值,int赋值0,string赋值""
	return &CodeError{code: code, msg: msg}
}

func getErrorString(e *CodeError) string {
	return fmt.Sprintf("%d|%s", e.code, e.msg)
}
