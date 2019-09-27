/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 18:45 2019-09-24
 */
package utils

import (
	"fmt"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	Ch0 selenium.WebDriver
	Ch1 selenium.WebDriver
	Ch2 selenium.WebDriver
	Ch3 selenium.WebDriver
	Ch4 selenium.WebDriver
	Ch5 selenium.WebDriver
	Ch6 selenium.WebDriver
	Ch7 selenium.WebDriver
	Ch8 selenium.WebDriver
	Ch9 selenium.WebDriver

	num int

	loc  sync.Mutex
	loc0 sync.Mutex
	loc1 sync.Mutex
	loc2 sync.Mutex
	loc3 sync.Mutex
	loc4 sync.Mutex
	loc5 sync.Mutex
	loc6 sync.Mutex
	loc7 sync.Mutex
	loc8 sync.Mutex
	loc9 sync.Mutex
	err  error
)

func init() {
	Ch0 = GenerateChrome()
	Ch1 = GenerateChrome()
	Ch2 = GenerateChrome()
	Ch3 = GenerateChrome()
	Ch4 = GenerateChrome()
	Ch5 = GenerateChrome()
	Ch6 = GenerateChrome()
	Ch7 = GenerateChrome()
	Ch8 = GenerateChrome()
	Ch9 = GenerateChrome()

	num = 0
}

func GenerateChrome() selenium.WebDriver {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images":     2,
		"profile.managed_default_content_settings.javascript": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		//Args: []string{
		//	"--headless", // 设置Chrome无头模式
		//	"--no-sandbox",
		//	"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		//},
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	service, err := selenium.NewChromeDriverService("/usr/local/bin/chromedriver", 9515, opts...)
	if service != nil {
		//defer func() {
		//	service.Stop()
		//}()
	}
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
	}
	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	//if webDriver != nil {
	//	defer func() {
	//		webDriver.Close()
	//	}()
	//}
	if err != nil {
		panic(err)
	}

	return webDriver
}

// StartChrome 启动谷歌浏览器headless模式
func StartChrome(url string) string {
	//time.Sleep(10 * time.Second)
	//return ""
	// 每一次休息一下 随机100~500  毫秒
	random := easyutils.Random(500, 1000)
	duration := time.Duration(random)
	time.Sleep(duration * time.Millisecond)

	loc.Lock()
	cc := 0
	if num <= 10 {
		cc = num
		num += 1
	} else {
		num = 0
	}
	loc.Unlock()

	switch cc {
	case 0:
		loc0.Lock()
		defer loc0.Unlock()
		for i := 0; i < 100; i++ {

			err := Ch0.Get(url)

			if err == nil {
				// 重试
				s, err := Ch0.PageSource()
				s = strings.TrimSpace(s)
				if s == "" {
					time.Sleep(time.Minute * 3)
					continue
				}

				if err == nil {
					return s
				} else {
					clog.Println(err)
					time.Sleep(time.Second * 10)
					continue
				}
			} else {
				clog.Println(err)
				time.Sleep(time.Second * 10)
				continue
			}
		}
	case 1:
		loc1.Lock()
		defer loc1.Unlock()
		for i := 0; i < 100; i++ {

			err := Ch1.Get(url)

			if err == nil {
				// 重试
				s, err := Ch1.PageSource()
				if s == "" {
					time.Sleep(time.Minute * 3)
					continue
				}
				if err == nil {
					return s
				} else {
					clog.Println(err)
					time.Sleep(time.Second * 10)
					continue
				}
			} else {
				clog.Println(err)
				time.Sleep(time.Second * 10)
				continue
			}
		}
	case 2:
		loc2.Lock()
		defer loc2.Unlock()
		for i := 0; i < 100; i++ {

			err := Ch2.Get(url)

			if err == nil {
				// 重试
				s, err := Ch2.PageSource()
				if s == "" {
					clog.PrintWa("this is none")
					time.Sleep(time.Minute * 3)
					continue
				}
				if err == nil {
					return s
				} else {
					clog.Println(err)
					time.Sleep(time.Second * 10)
					continue
				}
			} else {
				clog.Println(err)
				time.Sleep(time.Second * 10)
				continue
			}
		}
	case 3:
		loc3.Lock()
		defer loc3.Unlock()
		for i := 0; i < 100; i++ {

			err := Ch3.Get(url)

			if err == nil {
				// 重试
				s, err := Ch3.PageSource()
				if s == "" {
					clog.PrintWa("this is none")
					time.Sleep(time.Minute * 3)
					continue
				}
				if err == nil {
					return s
				} else {
					clog.Println(err)
					time.Sleep(time.Second * 10)
					continue
				}
			} else {
				clog.Println(err)
				time.Sleep(time.Second * 10)
				continue
			}
		}
	case 4:
		loc4.Lock()
		defer loc4.Unlock()
		for i := 0; i < 100; i++ {

			err := Ch4.Get(url)

			if err == nil {
				// 重试
				s, err := Ch4.PageSource()
				if s == "" {
					clog.PrintWa("this is none")
					time.Sleep(time.Minute * 3)
					continue
				}
				if err == nil {
					return s
				} else {
					clog.Println(err)
					time.Sleep(time.Second * 10)
					continue
				}
			} else {
				clog.Println(err)
				time.Sleep(time.Second * 10)
				continue
			}
		}
	case 5:
		loc5.Lock()
		defer loc5.Unlock()
		for i := 0; i < 100; i++ {

			err := Ch5.Get(url)

			if err == nil {
				// 重试
				s, err := Ch5.PageSource()
				if s == "" {
					clog.PrintWa("this is none")
					time.Sleep(time.Minute * 3)
					continue
				}
				if err == nil {
					return s
				} else {
					clog.Println(err)
					time.Sleep(time.Second * 10)
					continue
				}
			} else {
				clog.Println(err)
				time.Sleep(time.Second * 10)
				continue
			}
		}
	case 6:
		loc6.Lock()
		defer loc6.Unlock()
		for i := 0; i < 100; i++ {

			err := Ch6.Get(url)

			if err == nil {
				// 重试
				s, err := Ch6.PageSource()
				if s == "" {
					clog.PrintWa("this is none")
					time.Sleep(time.Minute * 3)
					continue
				}
				if err == nil {
					return s
				} else {
					clog.Println(err)
					time.Sleep(time.Second * 10)
					continue
				}
			} else {
				clog.Println(err)
				time.Sleep(time.Second * 10)
				continue
			}
		}
	case 7:
		loc7.Lock()
		defer loc7.Unlock()
		for i := 0; i < 100; i++ {

			err := Ch7.Get(url)

			if err == nil {
				// 重试
				s, err := Ch7.PageSource()
				if s == "" {
					clog.PrintWa("this is none")
					time.Sleep(time.Minute * 3)
					continue
				}
				if err == nil {
					return s
				} else {
					clog.Println(err)
					time.Sleep(time.Second * 10)
					continue
				}
			} else {
				clog.Println(err)
				time.Sleep(time.Second * 10)
				continue
			}
		}
	case 8:
		loc8.Lock()
		defer loc8.Unlock()
		for i := 0; i < 100; i++ {

			err := Ch8.Get(url)

			if err == nil {
				// 重试
				s, err := Ch8.PageSource()
				if s == "" {
					clog.PrintWa("this is none")
					time.Sleep(time.Minute * 3)
					continue
				}
				if err == nil {
					return s
				} else {
					clog.Println(err)
					time.Sleep(time.Second * 10)
					continue
				}
			} else {
				clog.Println(err)
				time.Sleep(time.Second * 10)
				continue
			}
		}
	case 9:
		loc9.Lock()
		defer loc9.Unlock()
		for i := 0; i < 100; i++ {

			err := Ch9.Get(url)

			if err == nil {
				// 重试
				s, err := Ch9.PageSource()

				if s == "" {
					clog.PrintWa("超量")
					time.Sleep(time.Minute * 3)
					continue
				}
				if err == nil {
					return s
				} else {
					clog.Println(err)
					time.Sleep(time.Second * 10)
					continue
				}
			} else {
				clog.Println(err)
				time.Sleep(time.Second * 10)
				continue
			}
		}
	}

	return ""
}

// StartChrome 启动谷歌浏览器headless模式
//func StartChrome(url string) string {
//	// 每一次休息一下 随机100~500  毫秒
//	random := easyutils.Random(100, 500)
//	duration := time.Duration(random)
//	time.Sleep(duration * time.Millisecond)
//	cc:
//	opts := []selenium.ServiceOption{}
//	caps := selenium.Capabilities{
//		"browserName":                      "chrome",
//	}
//
//	// 禁止加载图片，加快渲染速度
//	imagCaps := map[string]interface{}{
//		"profile.managed_default_content_settings.images": 2,
//		"profile.managed_default_content_settings.javascript": 2,
//	}
//
//	chromeCaps := chrome.Capabilities{
//		Prefs: imagCaps,
//		Path:  "",
//		//Args: []string{
//		//	"--headless", // 设置Chrome无头模式
//		//	"--no-sandbox",
//		//	"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
//		//},
//	}
//	caps.AddChrome(chromeCaps)
//	// 启动chromedriver，端口号可自定义
//	service, err := selenium.NewChromeDriverService("/usr/local/bin/chromedriver", 9515, opts...)
//	if service != nil {
//		defer func() {
//			service.Stop()
//		}()
//	}
//	if err != nil {
//		log.Printf("Error starting the ChromeDriver server: %v", err)
//	}
//	// 调起chrome浏览器
//	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
//	if webDriver != nil {
//		defer func() {
//			webDriver.Close()
//		}()
//	}
//	if err != nil {
//		panic(err)
//	}
//	// 这是目标网站留下的坑，不加这个在linux系统中会显示手机网页，每个网站的策略不一样，需要区别处理。
//	webDriver.AddCookie(&selenium.Cookie{
//		Name:  "defaultJumpDomain",
//		Value: "www",
//	})
//	// 导航到目标网站
//	err = webDriver.Get(url)
//	if err != nil {
//		clog.PrintWa(err)
//		clog.PrintWa(url)
//		//panic(fmt.Sprintf("Failed to load page: %s\n", err))
//		// 重试
//		time.Sleep(time.Second * 10)
//		goto cc
//	}
//	s, err := webDriver.PageSource()
//	if err != nil {
//		clog.Println(err)
//		clog.PrintWa(url)
//		time.Sleep(time.Second * 10)
//		goto cc
//	}
//
//
//	return s
//}
