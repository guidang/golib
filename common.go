package lib

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
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

/**
IP 转整型

*/
func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

/**
整型转 ip
*/
func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

/**
截取字符串
*/
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	//取末尾
	endLen := start + end

	//开始和结束大于总长度
	if start > length || endLen > length {
		return ""
	}

	//负数时，从后面起取值
	if start < 0 {
		startAbs := length + start
		if startAbs >= 0 {
			start = startAbs
			endLen = length
		}
	}

	if start < 0 || endLen > length {
		return ""
	}

	return string(rs[start:endLen])
}
