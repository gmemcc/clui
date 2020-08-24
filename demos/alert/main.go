package main

import ui "github.com/VladimirMarkelov/clui"

func main() {
	ui.InitLibrary()
	defer ui.DeinitLibrary()
	dlg := ui.CreateAlertDialog(" 提示信息 ", "can 中文字符串， be displayed correctly?", 30, "确定")
	dlg.View.SetTitleButtons(0)
	ui.MainLoop()
}
