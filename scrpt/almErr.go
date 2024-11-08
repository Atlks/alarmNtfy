package main

import (
	"fmt"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/kbinani/screenshot"
	"gocv.io/x/gocv"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	scrname     = "second_screen.png"
	threshold   = 0.5
	iconMeeting = "C:/db/ggl3.jpg"
	soundFile   = "C:/db/laor.mp3"
)

/*
*cant use...
可以运行go mod tidy来自动清理并更新缺少任何条目的go.mod和文件：go.sum

狂欢

复制代码
go mod tidy
go get -u gocv.io/x/gocv
*/
func main() {
	for {
		fmt.Println("这是一个无限循环，每 5 秒检测一次。")
		foreachx()
		time.Sleep(5 * time.Second)
	}
}

// mkdir2024 创建目录及其所有级联目录
func mkdir2024(pngFilepath string) error {
	// 提取文件路径中的目录部分
	directory := filepath.Dir(pngFilepath)
	// 创建目录，包括所有缺失的级联目录
	return os.MkdirAll(directory, os.ModePerm)
}

// generateFilenameTimepart 根据当前时间生成文件名的一部分
func generateFilenameTimepart() string {
	// 获取当前日期和时间
	now := time.Now()
	// 格式化为 "YYYY-MM-DD_HH-MM-SS"
	return now.Format("2006-01-02_15-04-05")
}

// getSecScr 获取第二个显示器的截图
func getSecScr() (string, error) {
	// 获取显示器信息
	numMonitors := screenshot.NumActiveDisplays()
	if numMonitors < 2 {
		return "", fmt.Errorf("没有找到第二个显示器")
	}

	// 打印所有显示器的信息
	for i := 0; i < numMonitors; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		fmt.Printf("Monitor %d: %dx%d at %d,%d\n", i+1, bounds.Dx(), bounds.Dy(), bounds.Min.X, bounds.Min.Y)
	}

	// 假设第二个屏幕是第二个显示器（索引为1）
	bounds := screenshot.GetDisplayBounds(1) // 在 Go 中，显示器索引从0开始，所以第二个显示器的索引是1
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return "", err
	}

	// 转换为 PNG 文件并保存
	fileName := "/zscrpic/second_screen" + generateFilenameTimepart() + ".png"
	err = mkdir2024(fileName)
	if err != nil {
		return "", err
	}

	// 打开文件用于写入
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 保存图片为 PNG 格式
	err = png.Encode(file, img)
	if err != nil {
		return "", err
	}

	// 打印保存的路径
	fmt.Println("第二个屏幕的截图已保存为 " + fileName)
	return fileName, nil
}

// 查找图标
func foreachx() {
	err, scrnm := getSecScr()
	if err != nil {
		fmt.Println("截屏失败：", err)
		return
	}

	// 读取截图和目标图标
	screenshot := gocv.IMRead(scrnm, gocv.IMReadColor)
	defer screenshot.Close()
	icon := gocv.IMRead(iconMeeting, gocv.IMReadColor)
	defer icon.Close()

	// 模板匹配
	result := gocv.NewMat()
	defer result.Close()
	gocv.MatchTemplate(screenshot, icon, &result, gocv.TmCCoeffNormed, gocv.NewMat())

	// 查找匹配的位置
	_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
	if maxVal >= threshold {
		fmt.Printf("图标找到，位置： %v\n", maxLoc)
		playSound(soundFile) // 播放声音
	} else {
		fmt.Println("图标未找到。")
	}
}

// 播放音乐文件
func playSound(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	time.Sleep(time.Second * 50) // 播放 5 秒
}
