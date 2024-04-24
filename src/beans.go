package src

import (
	"context"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"sync"

	"github.com/kbinani/screenshot"
)

var beansNow int
var mutex sync.Mutex
var BeansDataLR BeansData

func SelfLRCycle(rect RECT, selfLR chan string) {
	go func() {
		for {
			lr := <-selfLR
			BeansDataLR = BeansInitialization(rect, lr)
		}
	}()
}

func FBeansLCycle(beanleft chan int, ctx chan context.Context) {
	go func() {
		for {
			start := <-ctx
			go func() {
				for {
					select {
					case <-start.Done():
						log.Println("对局结束")
						return
					default:
						FindBeansleft(beanleft)
					}
				}
			}()
		}
	}()
}

func FindBeansleft(beanleft chan int) {
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func(i int, wg *sync.WaitGroup) {
			var garyCount int
			img, err := screenshot.Capture(BeansDataLR.X[i], BeansDataLR.Y, 8, 8)
			if err != nil {
				log.Println("Four-bean cycle err:", err)
				return
			}
			for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
				for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
					beans := img.At(x, y)
					beansIsExits := color.GrayModel.Convert(beans).(color.Gray)
					if beansIsExits.Y >= 128 {
						garyCount++
					}
				}
			}
			if garyCount > 60 {
				increment()
			}
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
	beanleft <- beansNow
	beansNow = 0
}

func increment() {
	// 锁定互斥锁
	mutex.Lock()
	beansNow++
	// 解锁互斥锁
	mutex.Unlock()
}
