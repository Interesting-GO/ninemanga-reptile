/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 18:45 2019-09-24
 */
package utils

import (
	"fmt"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
)

var (
	webDriver selenium.WebDriver
	err error
)

func init() {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName":                      "chrome",
	}

	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
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
	webDriver, err = selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	//if webDriver != nil {
	//	defer func() {
	//		webDriver.Close()
	//	}()
	//}
	if err != nil {
		panic(err)
	}
}

// StartChrome 启动谷歌浏览器headless模式
func StartChrome(url string) string {
	// 这是目标网站留下的坑，不加这个在linux系统中会显示手机网页，每个网站的策略不一样，需要区别处理。
	webDriver.AddCookie(&selenium.Cookie{
		Name:  "defaultJumpDomain",
		Value: "www",
	})
	// 导航到目标网站
	err = webDriver.Get(url)
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	s, err := webDriver.PageSource()
	if err != nil {
		clog.Println(err)
		return ""
	}
	return s
}
