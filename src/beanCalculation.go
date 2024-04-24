package src

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/kbinani/screenshot"
)

var beansAfter int

func TimingCycle(beanleft chan int) {
	go func(beanleft chan int) {
		for {
			beans := <-beanleft
			if beans < beansAfter && beansAfter-beans == 1 {
				SubCuntDown()
				jietu(fmt.Sprintf("替身前为%d豆,目前是%d豆", beansAfter, beans))
			}
			beansAfter = beans
		}
	}(beanleft)
}

var beansss int

func jietu(str string) {
	beansss++
	rect := GetWindow("雷电模拟器")
	img, err := screenshot.Capture(
		int(rect.Left),
		int(rect.Top),
		int(rect.Right),
		int(rect.Bottom))
	if err != nil {
		log.Println("Screenshot Capture:", err)
	}
	SaveImg(img, str+strconv.Itoa(beansss)+".png")
}

var Subnum int

// 替身倒计时
func SubCuntDown() {
	// for {
	// 	if Subnum == 0 {
	// 		log.Println("等待对局")
	// 	} else {
	// 		log.Printf("替身倒计时：%d 秒\n", Subnum)
	// 	}
	// }

	duration := 15 // 倒计时时长，单位为秒
	ticker := time.NewTicker(1 * time.Second)
	for i := duration; i > 0; i-- {
		fmt.Printf("替身：%d 秒\n", i)
		<-ticker.C
	}
	ticker.Stop()
}
