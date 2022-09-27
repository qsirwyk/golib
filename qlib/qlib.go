package qlib

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"syscall"
)

func StackTrace() string {
	funcName, file, line, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s:%d:%s\n", path.Base(file), line, runtime.FuncForPC(funcName).Name())
	}
	return "StackTrace"
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func CheckErrf(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// 判断文件是否存在
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// []string => []int
func String2Int(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}
	return res
}

func ExecCommand(name string, args ...string) {
	cmd := exec.Command(name, args...) // 拼接参数与命令

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var err error

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		log.Println(err)
	}

	//fmt.Print(stdout.String())
	fmt.Print(stderr.String())
}

type Color struct {
	black        int // 黑色 0
	blue         int // 蓝色 1
	green        int // 绿色 2
	cyan         int // 青色 3
	red          int // 红色 4
	purple       int // 紫色 5
	yellow       int // 黄色 6
	light_gray   int // 淡灰色（系统默认值）7
	gray         int // 灰色 8
	light_blue   int // 亮蓝色 9
	light_green  int // 亮绿色 10
	light_cyan   int // 亮青色 11
	light_red    int // 亮红色 12
	light_purple int // 亮紫色 13
	light_yellow int // 亮黄色 14
	white        int // 白色 15
}

func ColorPrint(s string, i int) { //设置终端字体颜色
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	fmt.Print(s)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}

var ClearScreen = func() {
	//执行clear指令清除控制台
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
