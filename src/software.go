package src

import (
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// 寻找模拟器
// check what is your screen scale, because if you have a 125% of scale, you must to multiply the rectangle for this percent, 1.25

func init() {
	var libuser32 = windows.NewLazySystemDLL("user32.dll")
	windowRect = libuser32.NewProc("GetWindowRect")
	findWindow = libuser32.NewProc("FindWindowW")
}

var (
	windowRect *windows.LazyProc
	findWindow *windows.LazyProc
)

// x,y 宽 高
type RECT struct {
	Left, Top, Right, Bottom int32
}

// 入口
func GetWindow(str string) RECT {
	simulatorExits, simulatorC := GetDnfWindows(str)
	if !simulatorExits {
		return RECT{}
	}
	//自定义截图
	// Left+Top左上角, Right宽, Bottom高 int32
	simulatorC.ExampleModifyDPI()
	return simulatorC
}

func (r *RECT) ExampleModifyDPI() {
	//获取动态的系统DPI_X 大小  防止分辨率改变后拿不到实际改变之后的dpi
	dpihwd, err := FindWindow("Program Manager")
	if err != nil {
		log.Println("Progman FindWindow err:", err)
	}
	conversionDpi := float64(GetSystemMetrics(dpihwd)) / 96
	r.Left = int32(float64(r.Left) * conversionDpi)
	r.Top = int32(float64(r.Top) * conversionDpi)
	r.Right = int32(float64(r.Right)*conversionDpi) - r.Left
	r.Bottom = int32(float64(r.Bottom)*conversionDpi) - r.Top
}

func GetSystemMetrics(nIndex uintptr) int {
	ret, _, _ := syscall.NewLazyDLL(`User32.dll`).NewProc(`GetDpiForWindow`).Call(uintptr(nIndex))
	return int(ret)
}

func GetDnfWindows(str string) (bool, RECT) {
	winN, err := FindWindow(str)
	if err != nil {
		return false, RECT{}
	}
	var window = &RECT{}
	return GetWindowRect(winN, window), *window
}

// 寻找模拟器窗口
func FindWindow(str string) (uintptr, error) {
	lpWindowName, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		return 0, err
	}
	hwnd := FindWindowNmae(nil, lpWindowName)
	return hwnd, nil
}

func FindWindowNmae(lpClassName, lpWindowName *uint16) uintptr {
	ret, _, _ := syscall.Syscall(findWindow.Addr(), 2,
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		0)

	return uintptr(ret)
}

// 获取窗口数据
// rect里面的left，top则表示左上角，right宽，bottom高
func GetWindowRect(hWnd uintptr, rect *RECT) bool {
	ret, _, _ := syscall.Syscall(windowRect.Addr(), 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rect)), 0)
	return ret != 0
}
