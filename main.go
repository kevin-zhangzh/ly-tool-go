package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	// Load user32.dll for mouse events
	user32           = windows.NewLazySystemDLL("user32.dll")
	procSetCursorPos = user32.NewProc("SetCursorPos")
	procGetCursorPos = user32.NewProc("GetCursorPos")
)

// 定义 POINT 结构体，与 Windows API 的 POINT 对应
type POINT struct {
	X int32
	Y int32
}

// 获取当前鼠标位置
func getCursorPos() (POINT, error) {
	var p POINT
	ret, _, err := procGetCursorPos.Call(uintptr(unsafe.Pointer(&p)))
	if ret == 0 { // 如果返回值为 0，表示调用失败
		return p, err
	}
	return p, nil
}

// 打印时间的函数
func printTime() {
	currentTime := time.Now()
	timeString := fmt.Sprintf("你好打工人，现在是公元%d年%d月%d日%d时%d分",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute())

	// 逐字打印时间
	for _, ch := range timeString {
		fmt.Print(string(ch))
		time.Sleep(200 * time.Millisecond) // 每个字之间暂停
	}
	fmt.Println()
}

// 调用 Windows API 移动鼠标
func moveMouse() {
	for {
		// 获取当前鼠标位置
		currentPos, err := getCursorPos()
		if err != nil {
			fmt.Println("Error getting cursor position:", err)
			return
		}

		// 随机生成偏移量
		moveX := rand.Intn(3) - 1 // -1, 0, 1
		moveY := rand.Intn(3) - 1 // -1, 0, 1

		// 计算新位置
		newX := currentPos.X + int32(moveX)
		newY := currentPos.Y + int32(moveY)

		// 调用 Windows API 移动鼠标
		procSetCursorPos.Call(uintptr(newX), uintptr(newY))

		// 每 50 秒移动一次
		time.Sleep(10 * time.Second)
	}
}

func main() {
	// 打印当前时间
	printTime()

	// 启动一个 goroutine 来定期移动鼠标
	go moveMouse()

	// 程序保持运行
	select {}
}
