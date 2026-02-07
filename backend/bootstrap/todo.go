package bootstrap

import "fast-image/global"

var TodoFunc func()

func InitTodo() {
	if TodoFunc != nil {
		TodoFunc()
	}
	global.Logger.Info("初始化TODO完成")
}
