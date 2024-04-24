package src

import (
	"github.com/kbinani/screenshot"
)

// 模拟器窗口豆子数据初始化
func BeansInitialization(rect RECT, selfLR string) BeansData {
	var xyinter XYInter
	xyi := xyinter.WindowState(rect)
	var beansData BeansData
	var XMultiple float64
	beansData.X = make(map[int]int, 4)
	if selfLR == "L" {
		XMultiple = xyi.XLMultiple
	} else {
		XMultiple = xyi.XRMultiple
	}
	for i := 0; i < 4; i++ {
		if i == 0 {
			beansData.X[i] = int(float64(rect.Left) + float64(rect.Right)/XMultiple)
		}
		beansData.X[i] = int(float64(rect.Left) + float64(rect.Right)/XMultiple + (float64(rect.Right)/xyi.Inter)*float64(i))
	}
	beansData.Y = int(float64(rect.Top) + float64(rect.Bottom)/xyi.YMultiple)
	return beansData
}

// 模拟器窗口倒计时数据初始化
func CountDownInitialization(rect RECT) CountDwonData {
	var xyinter XYInter
	xyi := xyinter.WindowState(rect)
	var countDownData = CountDwonData{
		int(float64(rect.Left) + float64(rect.Right)/xyi.XCMultiple),
		int(float64(rect.Top) + float64(rect.Bottom)/xyi.YCMultiple),
		int(float64(rect.Right) / xyi.WCMultiple),
		int(float64(rect.Bottom) / xyi.HCMultiple),
	}
	return countDownData
}

type CountDwonData struct {
	X, Y, W, H int
}

type BeansData struct {
	X map[int]int
	Y int
}

func fullScreenError(sof, win int32) bool {
	if sof >= win-1 && sof <= win+1 {
		return true
	}
	return false
}

type XYInter struct {
	XLMultiple, XRMultiple, YMultiple, Inter, XCMultiple, YCMultiple, WCMultiple, HCMultiple float64
}

// 通过不同的窗口尺寸分辨模拟器处于全屏，最大化，缩放状态
// 根据不同状态分配不同窗口尺寸计算比例
func (xyi *XYInter) WindowState(rect RECT) XYInter {
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

// 缩放
func (xyi *XYInter) Zoom() XYInter {
	xyi.Inter = 63.1
	xyi.YMultiple = 8.0
	// xyi.XMultiple = 11.1
	return *xyi
}

// 最大化
func (xyi *XYInter) MaxImize() XYInter {
	xyi.Inter = 66.3
	xyi.XLMultiple = 8.7
	xyi.XRMultiple = 1.23
	xyi.YMultiple = 8.2
	xyi.XCMultiple = 2.25
	xyi.YCMultiple = 10.2
	xyi.WCMultiple = 12.65
	xyi.HCMultiple = 9.2
	return *xyi
}

// F11 全屏
func (xyi *XYInter) FullScreen() XYInter {
	xyi.Inter = 67
	xyi.YMultiple = 10.8
	// xyi.XMultiple = 11
	return *xyi
}
