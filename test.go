package main

import (
	"errors"
	"fmt"
	"github.com/gookit/color"
	"github.com/qsirwyk/golib/util"
	"log"
	"time"
)

func init() {
	//log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	//log.SetPrefix("[GoLib] ")
}

func main() {
	//testColor()
	//testTrace()
	//util.BlockMain()
	testErr()
	//testClear()
}

func testRc4() {

}

func rc4() {

}

func testClear() {
	fmt.Println("I will clean the screen in 2 seconds!")
	time.Sleep(2 * time.Second)
	util.ClearScreen()
	fmt.Println("I'm alone...")
}

func testErr() {
	util.CheckErrf(errors.New("test"), "%d|%d", 123, 456)
	util.CheckExitErr(errors.New("err"))
	util.CheckErrf(errors.New("test"), "%d|%d", 123, 456)
	util.CheckErr(errors.New("err"))
}

func testTrace() {
	util.Trace("util.trace", 1)
	log.Printf("%s test2", "func")
}

func testColor() {
	// quick use package func
	color.Redp("Simple to use color")
	color.Redln("Simple to use color")
	color.Greenp("Simple to use color\n")
	color.Cyanln("Simple to use color")
	color.Yellowln("Simple to use color")

	// quick use like fmt.Print*
	color.Red.Println("Simple to use color")
	color.Green.Print("Simple to use color\n")
	color.Cyan.Printf("Simple to use %s\n", "color")
	color.Yellow.Printf("Simple to use %s\n", "color")

	// use like func
	red := color.FgRed.Render
	green := color.FgGreen.Render
	fmt.Printf("%s line %s library\n", red("Command"), green("color"))

	// custom color
	color.New(color.FgWhite, color.BgBlack).Println("custom color style")

	// can also:
	color.Style{color.FgCyan, color.OpBold}.Println("custom color style")

	// internal theme/style:
	color.Info.Tips("message")
	color.Info.Prompt("message")
	color.Info.Println("message")
	color.Warn.Println("message")
	color.Error.Println("message")

	// use style tag
	color.Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")
	// Custom label attr: Supports the use of 16 color names, 256 color values, rgb color values and hex color values
	color.Println("<fg=11aa23>he</><bg=120,35,156>llo</>, <fg=167;bg=232>wel</><fg=red>come</>")

	// apply a style tag
	color.Tag("info").Println("info style text")

	// prompt message
	color.Info.Prompt("prompt style message")
	color.Warn.Prompt("prompt style message")

	// tips message
	color.Info.Tips("tips style message")
	color.Warn.Tips("tips style message")
}
