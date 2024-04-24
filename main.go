package main

import (
	"context"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"main/src"
	"os"
	"os/signal"
	"syscall"
)

// go get -u github.com/kbinani/screenshot 截图
func main() {
	rect := src.GetWindow("雷电模拟器")

	var nameInitialization = src.NmaeInitialization(rect)
	var vsInitialization = src.VsInitialization(rect)
	var imgName = src.FNameCycle(nameInitialization)

	var vs = make(chan bool)
	var lr = make(chan string)
	src.FVsCycle(vsInitialization, vs)
	src.FSelfCycle(src.SelfDownInitialization(rect), imgName, vs, lr)
	src.SelfLRCycle(rect, lr)
	// 初始化模拟器窗口数据

	var countDownInitialization = src.CountDownInitialization(rect)
	//寻找开始倒计时
	var ctx = make(chan context.Context)
	src.FCountDownCycle(countDownInitialization, ctx)
	//寻豆
	var beanleft = make(chan int)
	src.TimingCycle(beanleft)
	src.FBeansLCycle(beanleft, ctx)

	// 创建一个信号通道
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

}
