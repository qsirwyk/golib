package util

import (
	"bytes"
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"runtime"
	"strconv"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	log.SetPrefix("[Util.Log] ")
}

// BlockMain 优雅的堵塞Main函数 直至手动退出 , 如果需要子协程 完成任务后自动退出 有两种方法
// a.使用可终止的上下文 ctx, cancel := context.WithCancel(context.Background())
// b.使用等待组 wg := &sync.WaitGroup{}
func BlockMain() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
}

// CheckErr 如果存在错误打印并退出
func CheckErr(err error) {
	if err != nil {
		//red := color.FgRed.Render
		//log.Printf(red("%v"), err)
		Trace(err.Error(), 2)
		os.Exit(1)
	}
}

// CheckErrf  如果存在错误打印并退出
func CheckErrf(err error, str string, args ...interface{}) {
	if err != nil {
		red := color.FgRed.Render
		green := color.FgGreen.Render
		log.Printf(green(str)+"\n"+red(err.Error()), args...)
		os.Exit(1)
	}
}

// ExecCommand 执行外部命令 注意参数得一个一个写且不要有空格
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

// String2Int 将字符串切片转换成整型切片 []string => []int
func String2Int(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}
	return res
}

// FileExist 检测文件是否存在
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// =================================================================================================

// Trace 使用自带的Log可以达到同样的效果
// log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
// log.SetPrefix("[GoLib] ")
func Trace(str string, skip int) {
	funcName, file, line, _ := runtime.Caller(skip)
	red := color.FgRed.Render
	green := color.FgGreen.Render
	//yellow := color.FgYellow.Render
	fmt.Printf(green("[%s | %s:%d | %s] ")+red("%v\n"), time.Now().Format("2006-01-02 15:04:05"), path.Base(file), line, runtime.FuncForPC(funcName).Name(), str)
}

func ClearScreen() {
	clear := make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	cls, ok := clear[runtime.GOOS]
	if ok {
		cls()
	} else {
		Trace("此终端不支持清屏", 1)
	}
}
