/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:57 2019-09-24
 */
package utils

import (
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Dow(url string) []byte {

	// 每一次休息一下 随机100~500  毫秒
	random := easyutils.Random(100, 500)
	duration := time.Duration(random)
	time.Sleep(duration * time.Millisecond)

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
		}else if e!= nil && i > 10 {
			return nil
		} else {
			if response == nil {
				clog.PrintWa(url)
				time.Sleep(time.Minute * 3)
				continue
			}
			bytes, e := ioutil.ReadAll(response.Body)
			if e != nil {
				clog.PrintWa(e)
				time.Sleep(time.Minute * 3)
				continue
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
	if len(submatch) >= 1 {
		if len(submatch[0]) >= 1 {
			s := submatch[0][1]
			i, e := strconv.Atoi(s)
			if e == nil {
				return i
			}
		}
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