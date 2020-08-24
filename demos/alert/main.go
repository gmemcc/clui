package main

import ui "github.com/VladimirMarkelov/clui"

func main() {
	ui.InitLibrary()
	defer ui.DeinitLibrary()
	ui.CreateAlertDialog(" 提示信息 ", "can 中文字符串 be displayed correctly?", "确定")
	ui.MainLoop()
}
