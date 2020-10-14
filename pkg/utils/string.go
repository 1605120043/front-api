package utils

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

//首字母转大写
func FirstLitterToUpper(str string) string {
	if len(str) == 0 {
		return ""
	}
	
	list := []rune(str)
	if list[0] >= 97 && list[0] <= 122 {
		list[0] -= 32
	}
	
	return string(list)
}

// 获取客户端ip
func GetClientIp() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// 生成短信验证码
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
