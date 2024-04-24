package src

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/kbinani/screenshot"
)

func FVsCycle(vsData VSNData, vs chan bool) {
	imgvsfile, err := os.Open("./img/finvs.png")
	if err != nil {
		log.Println("FCountDownCycle err:", err)
		return
	}
	imgvs, _, err := image.Decode(imgvsfile)
	if err != nil {
		log.Println("FCountDownCycle err:", err)
		return
	}
	imgvsfile.Close()
	go func(vs chan bool) {
		for {
			imgVS, err := screenshot.Capture(vsData.X, vsData.Y, vsData.W, vsData.H)
			if err != nil {
				log.Println("findSelf cycle err:", err)
				return
			}
			if Comparison(ImgtoBinaryzation(imgvs), ImgtoBinaryzation(imgVS)) < 15 {
				log.Println("发现对局VS")
				vs <- true
			}
		}
	}(vs)
}

func FNameCycle(nameData VSNData) image.Image {
	imgName, err := screenshot.Capture(nameData.X, nameData.Y, nameData.W, nameData.H)
	if err != nil {
		log.Println("findSelf cycle err:", err)
		return nil
	}
	SaveImg(imgName, "name.png")
	return imgName
}

func FSelfCycle(SelfData SelfData, imgName image.Image, vs chan bool, selfLR chan string) {
	go func(vsch chan bool) {
		for {
			<-vsch
			findSelf(SelfData, imgName, selfLR)
		}
	}(vs)
}

func findSelf(selfData SelfData, image59 image.Image, selfLR chan string) {
	imgL, err := screenshot.Capture(selfData.XL, selfData.YL, selfData.W, selfData.H)
	if err != nil {
		log.Println("findSelf cycle err:", err)
		return
	}
	imgR, err := screenshot.Capture(selfData.XR, selfData.YR, selfData.W, selfData.H)
	if err != nil {
		log.Println("findSelf cycle err:", err)
		return
	}
	if Comparison(ImgtoBinaryzation(imgL), image59) < 15 || Comparison(ImgtoBinaryzation(imgR), image59) < 15 {
		if Comparison(ImgtoBinaryzation(imgL), image59) < Comparison(ImgtoBinaryzation(imgR), image59) {
			SaveImg(imgL, fmt.Sprintf("%d-左.png", Comparison(ImgtoBinaryzation(imgL), image59)))
			SaveImg(imgR, fmt.Sprintf("%d-右.png", Comparison(ImgtoBinaryzation(imgR), image59)))
			log.Println("玩家在左边")
			selfLR <- "R"
		} else {
			log.Println(Comparison(ImgtoBinaryzation(imgL), image59), Comparison(ImgtoBinaryzation(imgR), image59))
			SaveImg(imgL, fmt.Sprintf("%d-左.png", Comparison(ImgtoBinaryzation(imgL), image59)))
			SaveImg(imgR, fmt.Sprintf("%d-右.png", Comparison(ImgtoBinaryzation(imgR), image59)))
			log.Println("玩家在右边")
			selfLR <- "L"
		}
	}
}

type SelfXYI struct {
	XLMultiple, YLMultiple, XRMultiple, YRMultiple, WMultiple, HMultiple float64
}

// 通过不同的窗口尺寸分辨模拟器处于全屏，最大化，缩放状态
// 根据不同状态分配不同窗口尺寸计算比例
func (xyi *SelfXYI) WindowState(rect RECT) SelfXYI {
	var xfullScreen = int32(screenshot.GetDisplayBounds(0).Dx())
	var yfullScreen = int32(screenshot.GetDisplayBounds(0).Dy())
	if fullScreenError(xfullScreen, rect.Right) && fullScreenError(yfullScreen, rect.Bottom) {
		//最大化
		xyi.FullScreen()
		return *xyi
	} else if fullScreenError(xfullScreen, rect.Right) {
		//全屏
		xyi.MaxImize()
		return *xyi
	}
	//缩放状态
	xyi.Zoom()
	return *xyi
}

// 模拟器窗口倒计时数据初始化
func SelfDownInitialization(rect RECT) SelfData {
	var xyinter SelfXYI
	xyi := xyinter.WindowState(rect)
	var SelfData = SelfData{
		int(float64(rect.Left) + float64(rect.Right)/xyi.XRMultiple),
		int(float64(rect.Top) + float64(rect.Bottom)/xyi.YRMultiple),
		int(float64(rect.Left) + float64(rect.Right)/xyi.XLMultiple),
		int(float64(rect.Top) + float64(rect.Bottom)/xyi.YLMultiple),
		int(float64(rect.Right) / xyi.WMultiple),
		int(float64(rect.Bottom) / xyi.HMultiple),
	}
	return SelfData
}

type SelfData struct {
	XR, YR, XL, YL, W, H int
}

type VSNData struct {
	X, Y, W, H int
}

type VSXYI struct {
	XVSMultiple, YVSMultiple, WVSMultiple, HVSMultiple, XNMultiple, YNMultiple, WNMultiple, HNMultiple float64
}

// 通过不同的窗口尺寸分辨模拟器处于全屏，最大化，缩放状态
// 根据不同状态分配不同窗口尺寸计算比例
func (xyi *VSXYI) WindowState(rect RECT) VSXYI {
	var xfullScreen = int32(screenshot.GetDisplayBounds(0).Dx())
	var yfullScreen = int32(screenshot.GetDisplayBounds(0).Dy())
	if fullScreenError(xfullScreen, rect.Right) && fullScreenError(yfullScreen, rect.Bottom) {
		//最大化
		xyi.FullScreen()
		return *xyi
	} else if fullScreenError(xfullScreen, rect.Right) {
		//全屏
		xyi.MaxImize()
		return *xyi
	}
	//缩放状态
	xyi.Zoom()
	return *xyi
}

// 模拟器窗口VS检测数据初始化
func VsInitialization(rect RECT) VSNData {
	var xyinter VSXYI
	xyi := xyinter.WindowState(rect)
	var vsData = VSNData{
		int(float64(rect.Left) + float64(rect.Right)/xyi.XVSMultiple),
		int(float64(rect.Top) + float64(rect.Bottom)/xyi.YVSMultiple),
		int(float64(rect.Right) / xyi.WVSMultiple),
		int(float64(rect.Bottom) / xyi.HVSMultiple),
	}
	return vsData
}

// 模拟器窗口VS检测数据初始化
func NmaeInitialization(rect RECT) VSNData {
	var xyinter VSXYI
	xyi := xyinter.WindowState(rect)
	var nameData = VSNData{
		int(float64(rect.Left) + float64(rect.Right)/xyi.XNMultiple),
		int(float64(rect.Top) + float64(rect.Bottom)/xyi.YNMultiple),
		int(float64(rect.Right) / xyi.WNMultiple),
		int(float64(rect.Bottom) / xyi.HNMultiple),
	}
	return nameData
}

// 缩放
// 868 748 180 240
// 207 91 195 38
func (xyi *VSXYI) Zoom() VSXYI {
	xyi.XVSMultiple = 2.21
	xyi.YVSMultiple = 1.36
	xyi.WVSMultiple = 10.66
	xyi.HVSMultiple = 4.25
	return *xyi
}

// 最大化
func (xyi *VSXYI) MaxImize() VSXYI {
	xyi.XVSMultiple = 2.21
	xyi.YVSMultiple = 1.36
	xyi.WVSMultiple = 10.66
	xyi.HVSMultiple = 4.25
	xyi.XNMultiple = 9.27
	xyi.YNMultiple = 11.2
	xyi.WNMultiple = 9.84
	xyi.HNMultiple = 26.84
	return *xyi
}

// F11 全屏
func (xyi *VSXYI) FullScreen() VSXYI {
	xyi.XVSMultiple = 2.21
	xyi.YVSMultiple = 1.36
	xyi.WVSMultiple = 10.66
	xyi.HVSMultiple = 4.25
	return *xyi
}

// 缩放
// 204 936 160 40
// 1522 931 204 40
func (xyi *SelfXYI) Zoom() SelfXYI {
	xyi.XRMultiple = 1.26
	xyi.YRMultiple = 1.09
	xyi.XLMultiple = 9.41
	xyi.YLMultiple = 1.08
	xyi.WMultiple = 12
	xyi.HMultiple = 25.5
	return *xyi
}

// 最大化
func (xyi *SelfXYI) MaxImize() SelfXYI {
	xyi.XRMultiple = 1.26
	xyi.YRMultiple = 1.10
	xyi.XLMultiple = 9.41
	xyi.YLMultiple = 1.09
	xyi.WMultiple = 9.6
	xyi.HMultiple = 25.5
	return *xyi
}

// F11 全屏
func (xyi *SelfXYI) FullScreen() SelfXYI {
	xyi.XRMultiple = 1.26
	xyi.YRMultiple = 1.09
	xyi.XLMultiple = 9.41
	xyi.YLMultiple = 1.08
	xyi.WMultiple = 9.6
	xyi.HMultiple = 25.5
	return *xyi
}
