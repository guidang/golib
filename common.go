package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

/**
键盘输入
*/
func Input(r *bufio.Reader) string {
	v, _, e := r.ReadLine()

	s := string(v)

	if e != nil {
		fmt.Println("输入错误，退出程序！")
		os.Exit(0)
	}

	if strings.EqualFold(s, "Q") {
		fmt.Println("您选择了 “退出程序”")
		os.Exit(0)
	}

	return s
}

/**
是否为自动更新
*/
func AutoUpdate(in bool, tb string) bool {
	if in {
		return true
	}

	fmt.Println("是否更新数据表 " + tb + "(Y/N):")
	var auto bool
	buf := bufio.NewReader(os.Stdin)
	for {
		s := Input(buf)
		if strings.EqualFold(s, "Y") {
			auto = true
			break
		} else if strings.EqualFold(s, "N") {
			auto = false
			break
		}
		fmt.Println("只能输入 Y 或 N,请重输")
	}

	return auto
}

/**
进度展示到窗口
*/
func UpdateProcess(str string, t int) {
	if t > 0 {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("\r%s", str)
	os.Stdout.Sync()
}

/**
时间戳　转　时间字符串
*/
func StrToDate(old string) string {
	st, err := strconv.ParseInt(old, 10, 64)
	if err != nil {
		return ""
	}

	return time.Unix(st, 0).Format("2006-01-02 15:04:05")
}
