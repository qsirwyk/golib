package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gookit/color"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
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

// CheckErr 如果存在错误打印
func CheckErr(err error) {
	if err != nil {
		//red := color.FgRed.Render
		//log.Printf(red("%v"), err)
		Trace(err.Error(), 2)
	}
}

// CheckErr 如果存在错误打印并退出
func CheckExitErr(err error) {
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
		log.SetPrefix("[CheckErrf] ")
		log.Printf(green(str)+"\n"+red(err.Error()), args...)
		log.SetPrefix("[Util.Log] ")
	}
}

// CheckErrf  如果存在错误打印并退出
func CheckExitErrf(err error, str string, args ...interface{}) {
	if err != nil {
		red := color.FgRed.Render
		green := color.FgGreen.Render
		log.SetPrefix("[CheckExitErrf] ")
		log.Printf(green(str)+"\n"+red(err.Error()), args...)
		log.SetPrefix("[Util.Log] ")
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
	clear["darwin"] = func() {
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
		Trace("此终端不支持清屏:"+runtime.GOOS, 2)
	}
}

// 生成s,e之间的随机数
func RndInt(s, e int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(e-s+1) + s
	return num
}

func Rc4(key, data string) string {
	src := data
	S := make([]int, 256)
	lenKey := len(key)
	for i := 0; i < 256; i++ {
		S[i] = i
	}
	j := 0
	for i := 0; i < 256; i++ {
		j = (j + S[i] + int(key[i%lenKey])) % 256
		S[i], S[j] = S[j], S[i]
	}
	lenData := len(src)
	var output []byte

	for a, j, i := 0, 0, 0; i < lenData; i++ {
		a = (a + 1) % 256
		j = (j + S[a]) % 256
		S[a], S[j] = S[j], S[a]
		k := S[((S[a] + S[j]) % 256)]
		output = append(output, byte(int(src[i])^k))
	}
	return string(Bin2Hex(output))
}

func UnRc4(key, data string) string {
	src := Hex2Bin([]byte(data))
	S := make([]int, 256)
	lenKey := len(key)
	for i := 0; i < 256; i++ {
		S[i] = i
	}
	j := 0
	for i := 0; i < 256; i++ {
		j = (j + S[i] + int(key[i%lenKey])) % 256
		S[i], S[j] = S[j], S[i]
	}
	lenData := len(src)
	var output []byte

	for a, j, i := 0, 0, 0; i < lenData; i++ {
		a = (a + 1) % 256
		j = (j + S[a]) % 256
		S[a], S[j] = S[j], S[a]
		k := S[((S[a] + S[j]) % 256)]
		output = append(output, byte(int(src[i])^k))
	}
	return string(output)
}

func Hex2Bin(src []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(dst, src)
	return dst
}

func Bin2Hex(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

func HexToBin(str string) []byte {
	slen := len(str)
	bHex := make([]byte, len(str)/2)
	ii := 0
	for i := 0; i < len(str); i = i + 2 {
		if slen != 1 {
			ss := string(str[i]) + string(str[i+1])
			bt, _ := strconv.ParseInt(ss, 16, 32)
			bHex[ii] = byte(bt)
			ii = ii + 1
			slen = slen - 2
		}
	}
	return bHex
}

// CreateDir 创建多级目录 调用os.MkdirAll递归创建文件夹
func CreateDir(dirPath string) error {
	if !FileExist(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("文件夹("+dirPath+")创建失败,error info:", err)
			return err
		}
		return err
	}
	return nil
}

// CreateFile 创建文件
func CreateFile(filePath string) {
	if !FileExist(filePath) {
		//创建文件
		f, err := os.Create(filePath)
		//判断是否出错
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
	}
}

// Md5 返回一个32位md5加密后的字符串
func Md5(str string, upper bool) string {
	h := md5.New()
	h.Write([]byte(str))
	if upper {
		return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	}
	return hex.EncodeToString(h.Sum(nil))
}

// GetMac 获取物理网卡MAC
func GetMac() string {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("获取物理网卡失败:" + err.Error())
	}
	inter := interfaces[0]
	mac := inter.HardwareAddr.String() //获取本机MAC地址
	return mac
}

// GetCpuId 获取CPUID
func GetCpuId() string {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	str := string(out)
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")
	return str[11:]
}
