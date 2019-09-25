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
	"github.com/dollarkillerx/easyutils/concurrent"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"sync"
	"time"
)

var (
	poll *concurrent.ObjPoll
	loc  sync.Mutex
	err  error
)

func init() {
	for i := 0; i < 10; i++ {
		poll = concurrent.NewObjPoll(PoolItemGenerate, 10)
	}
}

func PoolItemGenerate() interface{} {
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
kb:
	// 获取一个webdriver
	drive, e := poll.GetObj(time.Minute * 10)
	if e != nil {
		goto kb
	}
	defer func() {
		poll.Release(drive)
	}()
	driver := drive.(selenium.WebDriver)

	// 每一次休息一下 随机100~500  毫秒
	random := easyutils.Random(100, 500)
	duration := time.Duration(random)
	time.Sleep(duration * time.Millisecond)
ko:
	// 这是目标网站留下的坑，不加这个在linux系统中会显示手机网页，每个网站的策略不一样，需要区别处理。
	driver.AddCookie(&selenium.Cookie{
		Name:  "defaultJumpDomain",
		Value: "www",
	})
	// 导航到目标网站
	err = driver.Get(url)
	if err != nil {
		clog.PrintWa(fmt.Sprintf("Failed to load page: %s\n", err))
		clog.PrintWa(url)
		time.Sleep(10 * time.Second)
		goto ko
	}
	s, err := driver.PageSource()
	if err != nil {
		clog.PrintWa(err)
		clog.PrintWa(url)
		time.Sleep(10 * time.Second)
		goto ko
	}
	return s
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
