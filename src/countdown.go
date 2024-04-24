package src

import (
	"context"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func FCountDownCycle(countDownData CountDwonData, ctx chan context.Context) {
	img59file, err := os.Open("./img/fin59.png")
	if err != nil {
		log.Println("FCountDownCycle err:", err)
		return
	}
	img59, _, err := image.Decode(img59file)
	if err != nil {
		log.Println("FCountDownCycle err:", err)
		return
	}
	img59file.Close()

	go func(ctx chan context.Context) {
		for {
			time.Sleep(time.Second)
			findCountDown(countDownData, img59, ctx)
		}
	}(ctx)
}

func findCountDown(countDownData CountDwonData, image59 image.Image, ctx chan context.Context) {
	img, err := screenshot.Capture(countDownData.X, countDownData.Y, countDownData.W, countDownData.H)
	if err != nil {
		log.Println("Four-bean cycle err:", err)
		return
	}
	if Comparison(ImgtoBinaryzation(img), image59) < 15 {
		SaveImg(img, fmt.Sprintf("%s-%d.png", time.Now().Format("20060102-150304"), Comparison(ImgtoBinaryzation(img), image59)))
		start, _ := context.WithDeadline(context.Background(), time.Now().Add(59*time.Second))
		// defer cancel()
		ctx <- start
		log.Println("对局开始")
	}
}
