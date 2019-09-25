/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:57 2019-09-24
 */
package utils

import (
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Dow(url string) []byte {
	// 下载网页
	for i := 0; i < 100; i++ {
		response, e := httplib.ProxyDow(url, "127.0.0.1:8001")
		if response != nil {
			defer response.Body.Close()
		}

		// 如果本次请求没有异常
		if e != nil && i < 3 {
			clog.PrintEr("第" + strconv.Itoa(i) + " 下载超时 Url:" + url)
			time.Sleep(3 * time.Second)
			continue
		} else if e != nil && i > 3 {
			clog.PrintEr("第" + strconv.Itoa(i) + " 下载超时 Url:" + url)
			time.Sleep(20 * time.Second)
			continue
		} else {
			bytes, e := ioutil.ReadAll(response.Body)
			if e != nil {
				return nil
			}
			return bytes
		}
	}
	return nil
}

func GetNum(str string) int {
	reg := `(\d+)`
	compile := regexp.MustCompile(reg)
	submatch := compile.FindAllStringSubmatch(str, -1)
	i, e := strconv.Atoi(submatch[0][1])
	if e == nil {
		return i
	}
	return 0
}

func GetPage(str string) int {
	index := strings.LastIndex(str, "/") + 1
	s := str[index:]
	i, e := strconv.Atoi(s)
	if e != nil {
		clog.Println(i)
		panic(e)
	}else {
		return i
	}
}

func DowChrom(url string) []byte  {
	// 下载网页
	for i := 0; i < 100; i++ {
		s, e := AnalysisHtml(url)
		// 如果本次请求没有异常
		if e != nil && i < 3 {
			clog.PrintEr("第" + strconv.Itoa(i) + " 下载超时 Url:" + url)
			time.Sleep(3 * time.Second)
			continue
		} else if e != nil && i > 3 {
			clog.PrintEr("第" + strconv.Itoa(i) + " 下载超时 Url:" + url)
			time.Sleep(20 * time.Second)
			continue
		} else {
			return []byte(s)
		}
	}
	return nil
}